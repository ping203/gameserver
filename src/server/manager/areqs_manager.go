package manager

import (
	"errors"
	"sync"
	"time"
)

var (
	// ErrTimeOut 超时错误
	ErrTimeOut = errors.New("ErrTimeOut")
	// ErrCancel 取消请求.
	ErrCancel = errors.New("ErrCancel")
	// ErrUnknown ...
	ErrUnknown = errors.New("ErrUnknown")
)

// Requester ...
type Requester interface {
	// Req 发送请求，等待回应
	Req(req func(int64) error, cbk func(interface{}, error))
	// ReqTimeOut 发送请求，等待回应, 超时返回 ErrTimeOut
	ReqTimeOut(req func(int64) error, cbk func(interface{}, error), timeOut time.Duration)
}

type respCallBack func(interface{}, error)

type waitResp struct {
	cbk   respCallBack
	timer *time.Timer
}

type iWorker interface {
	Post(func())
}

// RespDispatcher 派发响应消息
type RespDispatcher struct {
	waits  map[int64]waitResp
	seq    int64
	worker iWorker

	waitsMtx sync.Mutex
}

// NewRespDispatcher ...
func NewRespDispatcher(worker iWorker) *RespDispatcher {
	return &RespDispatcher{
		waits:  make(map[int64]waitResp),
		worker: worker,
	}
}

func (r *RespDispatcher) newSeqID() int64 {
	r.seq++
	if r.seq == 0 { // 0 用于验错.
		r.seq++
	}
	return r.seq
}

func (r *RespDispatcher) addWaitResp(cbk respCallBack) (seq int64) {
	r.waitsMtx.Lock()
	seq = r.newSeqID()
	r.waits[seq] = waitResp{cbk: cbk}
	r.waitsMtx.Unlock()
	return
}

func (r *RespDispatcher) addWaitRespTimeout(cbk respCallBack, timeout time.Duration) (seq int64) {
	r.waitsMtx.Lock()
	seq = r.newSeqID()
	timer := time.AfterFunc(timeout, func() {
		if r.worker != nil {
			r.worker.Post(func() {
				r.OnErr(seq, ErrTimeOut)
			})
		} else {
			r.OnErr(seq, ErrTimeOut)
		}
	})
	r.waits[seq] = waitResp{cbk: cbk, timer: timer}
	r.waitsMtx.Unlock()
	return
}

func (r *RespDispatcher) removeWaitResp(seq int64) (ret waitResp, ok bool) {
	r.waitsMtx.Lock()
	ret, ok = r.waits[seq]
	if ok {
		delete(r.waits, seq)
	}
	r.waitsMtx.Unlock()
	return
}

func (r *RespDispatcher) doResp(wait waitResp, resp interface{}, err error) {
	if wait.timer != nil {
		wait.timer.Stop()
	}
	wait.cbk(resp, err)
}

// OnResp 接收响应消息
func (r *RespDispatcher) OnResp(seq int64, resp interface{}) bool {
	if wait, ok := r.removeWaitResp(seq); ok {
		r.doResp(wait, resp, nil)
		return true
	}
	return false
}

// OnErr 接收错误消息
func (r *RespDispatcher) OnErr(seq int64, err error) bool {
	if err == nil {
		err = ErrUnknown
	}

	if wait, ok := r.removeWaitResp(seq); ok {
		r.doResp(wait, nil, err)
		return true
	}
	return false
}

// OnErrAll ...
func (r *RespDispatcher) OnErrAll(err error) {
	if err == nil {
		err = ErrUnknown
	}

	r.waitsMtx.Lock()
	waits := r.waits
	r.waits = make(map[int64]waitResp)
	r.waitsMtx.Unlock()

	for _, wait := range waits {
		r.doResp(wait, nil, err)
	}
}

type requesterS RespDispatcher

// NewRequester ...
func NewRequester(r *RespDispatcher) Requester {
	return (*requesterS)(r)
}

func (r *requesterS) getRespr() *RespDispatcher {
	return (*RespDispatcher)(r)
}

func (r *requesterS) Req(req func(int64) error, cbk func(interface{}, error)) {
	if cbk == nil {
		panic("cbk is nil")
	}
	respr := r.getRespr()
	seq := respr.addWaitResp(cbk)
	// 执行真正的请求.
	err := req(seq)
	// 错误处理.
	if err != nil {
		if respr.worker != nil {
			// 必须通过另一协程Post到工作协程, 如果在同一协程Post是可能死锁的.
			go func() {
				respr.worker.Post(func() {
					respr.OnErr(seq, err)
				})
			}()
		} else {
			respr.OnErr(seq, err)
		}
	}
}

func (r *requesterS) ReqTimeOut(req func(int64) error, cbk func(interface{}, error), timeOut time.Duration) {
	if cbk == nil {
		panic("cbk is nil")
	}
	respr := r.getRespr()
	seq := respr.addWaitRespTimeout(cbk, timeOut)
	// 执行真正的请求.
	err := req(seq)
	// 错误处理.
	if err != nil {
		if respr.worker != nil {
			// 必须通过另一协程Post到工作协程, 如果在同一协程Post是可能死锁的.
			go func() {
				respr.worker.Post(func() {
					respr.OnErr(seq, err)
				})
			}()
		} else {
			respr.OnErr(seq, err)
		}
	}
}

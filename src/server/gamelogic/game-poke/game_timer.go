package poke

import "time"

type gameTimer struct {
	startTime   time.Time
	duraion     time.Duration
	timerCancel func()
	post        func(func())
}

func NewGameTimer(post func(func())) *gameTimer {
	if post == nil {
		panic("card game util new gametimer err!")
	}

	result := &gameTimer{post: post}
	return result
}

func (g *gameTimer) timeAfter(d time.Duration, f func()) (cancel func()) {
	//if TestOption.EnableFastSpeedTest {
	//	if d > time.Millisecond {
	//		d = time.Millisecond
	//	}
	//}

	var stop bool
	timer := time.AfterFunc(d, func() {
		g.post(func() {
			if !stop {
				f()
			}
		})
	})
	return func() { stop = true; timer.Stop() }
}

func (t *gameTimer) start(d time.Duration, f func()) {
	t.stop()

	t.startTime = time.Now()
	t.duraion = d

	t.timerCancel = t.timeAfter(d, func() {
		t.timerCancel = nil
		t.startTime = time.Time{}
		t.duraion = 0
		f()
	})
}

func (t *gameTimer) stop() {
	if t.isStart() {
		t.timerCancel()
		t.timerCancel = nil
		t.startTime = time.Time{}
		t.duraion = 0
	}
}

func (t *gameTimer) isStart() bool {
	return t.timerCancel != nil
}

func (t *gameTimer) elapsed() time.Duration {
	return time.Now().Sub(t.startTime)
}

func (t *gameTimer) left() time.Duration {
	d := t.duraion - t.elapsed()
	if d < 0 {
		d = 0
	}
	return d
}

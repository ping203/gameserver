package internal

import (
	"encoding/binary"
	"errors"
	"io"
	"net"
	"reflect"

	"server/base"
	"server/msg"
	"server/protobuf"

	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/module"
	"github.com/wenxiu2199/gameserver/src/server/gameproto/gamedef"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type writemsg struct {
	ext   interface{}
	msgid uint16
	msg   interface{}
}

type Client struct {
	*module.Skeleton
	address  string
	closeSig []chan bool

	writeChan    chan writemsg
	conn         net.Conn
	lenMsgLen    uint32
	minMsgLen    uint32
	littleEndian bool
	maxMsgLen    uint32
	processor    *protobuf.Processor

	userID  uint64
	general *gamedef.General
}

func (p *Client) Init() {
	p.Skeleton = skeleton
	p.closeSig = make([]chan bool, 0, 3)
	p.writeChan = make(chan writemsg, 10000)
	p.lenMsgLen = 2
	p.minMsgLen = 1
	p.maxMsgLen = 4096
	p.littleEndian = false
	p.processor = msg.Processor
	p.init()
}

func (p *Client) Run(closeSig chan bool) {
	p.Close(closeSig)
	conn, err := net.Dial("tcp", "127.0.0.1:3563")
	if err != nil {
		panic(err)
	}
	p.conn = conn
	p.process()

	sigSkeleton := p.getCloseSig()
	go func() {
		p.req()
	}()
	p.Skeleton.Run(sigSkeleton)

}

func (p *Client) getCloseSig() chan bool {
	sig := make(chan bool)
	p.closeSig = append(p.closeSig, sig)
	return sig
}
func (p *Client) process() {
	go func() {
		p.readLoop()
	}()
	go func() {
		p.writeLoop()
	}()

}

func (p *Client) read() ([]byte, error) {
	var b [4]byte
	bufMsgLen := b[:p.lenMsgLen]

	// read len
	_, err := io.ReadFull(p.conn, bufMsgLen)
	if err != nil {
		return nil, err
	}

	// parse len
	var msgLen uint32
	switch p.lenMsgLen {
	case 1:
		msgLen = uint32(bufMsgLen[0])
	case 2:
		if p.littleEndian {
			msgLen = uint32(binary.LittleEndian.Uint16(bufMsgLen))
		} else {
			msgLen = uint32(binary.BigEndian.Uint16(bufMsgLen))
		}
	case 4:
		if p.littleEndian {
			msgLen = binary.LittleEndian.Uint32(bufMsgLen)
		} else {
			msgLen = binary.BigEndian.Uint32(bufMsgLen)
		}
	}

	// check len
	if msgLen > p.maxMsgLen {
		return nil, errors.New("message too long")
	} else if msgLen < p.minMsgLen {
		return nil, errors.New("message too short")
	}

	// data
	msgData := make([]byte, msgLen)
	if _, err := io.ReadFull(p.conn, msgData); err != nil {
		return nil, err
	}

	return msgData, nil
}

func (p *Client) readLoop() {
	for {
		data, err := p.read()
		if err == io.EOF {
			return
		}
		if err != nil {
			panic("read err")
		}
		msgRaw, err := p.processor.Unmarshal(data)
		if err != nil {
			continue
		}
		err = p.processor.Route(msgRaw, nil)
		if err != nil {
			panic(err.Error())
		}
	}
}

func (p *Client) writeLoop() {
	for msg := range p.writeChan {
		// len + data
		data, err := proto.Marshal(msg.msg.(proto.Message))
		if err != nil {
			panic("msg err")
		}

		if msg.msgid == 0 {
			msg.msgid = p.stringHash(reflect.TypeOf(msg.msg).String())
		}
		m := make([]byte, 2+2+len(data))
		// 默认使用大端序
		binary.BigEndian.PutUint16(m, uint16(len(data))+2)
		binary.BigEndian.PutUint16(m[2:], msg.msgid)
		copy(m[4:], data)
		// 发送消息
		p.conn.Write(m)
	}
}

func (p *Client) WriteMsg(message proto.Message) {
	p.writeChan <- writemsg{
		msg: message,
	}
}

func (p *Client) Close(closeSig chan bool) {
	go func() {
		select {
		case <-closeSig:
			for _, v := range p.closeSig {
				v <- true
			}
			break
		}
	}()
}

func (p *Client) stringHash(s string) (hash uint16) {
	for _, c := range s {
		ch := uint16(c)
		hash = hash + ((hash) << 5) + ch + (ch << 7)
	}
	return
}

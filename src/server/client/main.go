package main

import (
	"encoding/binary"
	"net"
	"reflect"
	"time"

	"server/gameproto/cmsg"

	"github.com/golang/protobuf/proto"
)

func StringHash(s string) (hash uint16) {
	for _, c := range s {
		ch := uint16(c)
		hash = hash + ((hash) << 5) + ch + (ch << 7)
	}
	return
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3563")
	if err != nil {
		panic(err)
	}

	// Hello 消息（JSON 格式）
	// 对应游戏服务器 Hello 消息结构体
	msg := &cmsg.CReqAuth{
		Account: "1",
	}
	data, _ := proto.Marshal(msg)

	// len + data
	m := make([]byte, 2+2+len(data))

	// 默认使用大端序
	binary.BigEndian.PutUint16(m, uint16(len(data))+2)
	binary.BigEndian.PutUint16(m[2:], StringHash(reflect.TypeOf(msg).String()))

	copy(m[4:], data)

	// 发送消息
	conn.Write(m)

	time.Sleep(10 * time.Second)
}

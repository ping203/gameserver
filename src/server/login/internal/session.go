package internal

import (
	"sync"

	"github.com/name5566/leaf/gate"
)

type loginState int

const (
	stateUnAuth   loginState = 0
	stateAuthing  loginState = 1
	stateAuthed   loginState = 2
	stateLogining loginState = 3
	stateLogined  loginState = 4
	stateClosed   loginState = -1
)

type session struct {
	agent   *gate.Agent
	userID  uint64
	account string
	state   loginState

	// 等待登录和关闭
	wgLogin  sync.WaitGroup
	wgLogout sync.WaitGroup
}

package internal

import (
	"sync"
	"time"

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
	agent   gate.Agent
	userID  uint64
	account string
	state   loginState

	waitAuth *time.Timer
	// 等待登录和关闭
	wgLogin  sync.WaitGroup
	wgLogout sync.WaitGroup
}

func (p *session) isAuthed() bool {
	if p.state >= stateAuthed {
		return true
	}
	return false
}

func (p *session) isLogined() bool {
	if p.state >= stateLogined {
		return true
	}
	return false
}

func (p *session) isClosed() bool {
	return p.state == stateClosed
}

func (p *session) setClosed() {
	p.state = stateClosed
	if p.waitAuth != nil {
		p.waitAuth.Stop()
	}
}

func (p *session) addLogin() {
	p.wgLogin.Add(1)
}

func (p *session) doneLogin() {
	p.wgLogin.Done()
}

func (p *session) addLogout() {
	p.wgLogout.Add(1)
}

func (p *session) doneLogout() {
	p.wgLogout.Done()
}

package internal

import (
	"fmt"
	"math/rand"
	"sync"

	"server/util"

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

type sign struct {
	token  string
	expire int64
}

type session struct {
	agent   gate.Agent
	userID  uint64
	account string
	state   loginState
	sign    *sign

	// 等待登录和关闭
	wgLogin  sync.WaitGroup
	wgLogout sync.WaitGroup
}

func (p *session) isAuthing() bool {
	return p.state == stateAuthing
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

// 鉴权
func (p *session) auth(userID uint64) string {
	p.sign = &sign{}
	p.sign.token = p.createSign(userID)
	p.sign.expire = util.GetCurrentTimestamp() + 300
	return p.sign.token
}

func (p *session) createSign(userID uint64) string {
	num := rand.Int31n(100000)
	s := util.MD5(fmt.Sprintf("%d%d", userID, num))
	return s
}

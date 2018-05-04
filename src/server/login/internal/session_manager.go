package internal

import "github.com/name5566/leaf/gate"

type sessionManager struct {
	agent2Session   map[gate.Agent]*session
	userID2Session  map[uint64]*session
	account2Session map[string]*session
}

func (p *sessionManager) init() {
	p.agent2Session = make(map[gate.Agent]*session)
	p.userID2Session = make(map[uint64]*session)
	p.account2Session = make(map[string]*session)
}

func (p *sessionManager) addSession(agent gate.Agent) {
	s := &session{}
	s.agent = agent
	p.agent2Session[agent] = s
}

func (p *sessionManager) removeSession(agent gate.Agent) {

}

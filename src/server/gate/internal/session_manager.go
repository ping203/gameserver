package internal

import (
	"github.com/name5566/leaf/gate"
)

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
	if s, ok := p.agent2Session[agent]; ok {
		ss, exist := p.userID2Session[s.userID]
		if exist && ss == s {
			delete(p.userID2Session, s.userID)
		}

		ss, exist = p.account2Session[s.account]
		if exist && ss == s {
			delete(p.userID2Session, s.userID)
		}
		s.setClosed()
		agent.Destroy()
		agent.Close()

		delete(p.agent2Session, agent)
	}
}

func (p *sessionManager) getSessionByUserID(userID uint64) (*session, bool) {
	if s, ok := p.userID2Session[userID]; ok {
		return s, true
	}
	return nil, false
}

func (p *sessionManager) addUserOnAuth(agent gate.Agent) {
	if s, ok := p.agent2Session[agent]; ok {
		s.state = stateAuthing
	}
}

func (p *sessionManager) addUserOnAuthSuccess(agent gate.Agent, userID uint64, account string) {
	if s, ok := p.agent2Session[agent]; ok {
		s.userID = userID
		s.state = stateAuthed
		s.account = account
		if s.waitAuth != nil {
			s.waitAuth.Stop()
		}
	}
}

func (p *sessionManager) addUserOnLogin(agent gate.Agent) {
	if s, ok := p.agent2Session[agent]; ok {
		s.state = stateLogining
	}
}

func (p *sessionManager) addUserOnLoginSuccess(agent gate.Agent) {
	if s, ok := p.agent2Session[agent]; ok {
		s.state = stateLogined
		p.userID2Session[s.userID] = s
		p.account2Session[s.account] = s
	}
}

func (p *sessionManager) execByEverySession(f func(*session)) {
	for _, s := range p.userID2Session {
		f(s)
	}
}

func (p *sessionManager) getSessionByAccount(account string) (*session, bool) {
	if s, ok := p.account2Session[account]; ok {
		return s, true
	}
	return nil, false
}

package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/wenxiu2199/gameserver/src/server/gameproto/gamedef"
)

type userManager struct {
	users map[uint64]*user
}

func (p *userManager) init() {
	p.users = make(map[uint64]*user)
}

func (p *userManager) close() {

}

func (p *userManager) onUserEnter(userID uint64, account string, extra *gamedef.ExtraAccountInfo, agent gate.Agent, callBack func(*user, error)) {
	u, exist := p.findUser(userID)
	agent.SetUserData(userID)
	if !exist {
		p.addUser(userID, account, extra, agent, callBack)
	} else {
		// 替换原来的登录
		if u.Agent != agent {
			// todo 通知账号被登录
			//u.send2Msg(&)
			u.Agent = agent
			u.Agent.SetUserData(uint64(0))
		}
		if callBack != nil {
			callBack(u, nil)
		}
	}
}

func (p *userManager) addUser(userID uint64, account string, extra *gamedef.ExtraAccountInfo, agent gate.Agent, callBack func(*user, error)) {
	// 从db拉取数据
	u := &user{
		Agent:   agent,
		account: account,
	}

	// 保存玩家信息
	p.users[userID] = u
	p.loadGameUserData(userID, u, callBack)
}

func (p *userManager) removeUser(userID uint64) {
	delete(p.users, userID)
}

func (p *userManager) findUser(userID uint64) (*user, bool) {
	u, exist := p.users[userID]
	return u, exist
}

// 加载用户数据,如果不存在则尝试拉取
func (p *userManager) loadUser(userID uint64, callBack func(*user, error)) {
	if user, exist := p.findUser(userID); exist {
		callBack(user, nil)
	}

	p.addUser(userID, "", nil, nil, callBack)
}

func (p *userManager) loadGameUserData(userID uint64, u *user, callBack func(*user, error)) {
	f := func(data *gamedef.UserData, err error) {
		if err != nil {
			if callBack != nil {
				callBack(u, err)
			}
			return
		}

		u.info = data
		u.general.init(u, data.Generals)

		if callBack != nil {
			callBack(u, nil)
		}
	}
	dbMgr.LoadUserAsync(userID, f)
}

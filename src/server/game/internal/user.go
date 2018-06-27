package internal

import (
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/gate"
	"github.com/wenxiu2199/gameserver/src/server/gameproto/cmsg"
	"github.com/wenxiu2199/gameserver/src/server/gameproto/emsg"
	"github.com/wenxiu2199/gameserver/src/server/gameproto/gamedef"
)

type user struct {
	gate.Agent
	// 用户数据
	account string
	info    *gamedef.UserData

	general

	connected bool

	// 游戏数据
	gameID uint32

	saveCancel   func()
	updateNotify *cmsg.CNotifyDataChange
}

func (p *user) login() {
	p.connected = true
}

func (p *user) logout() {
	p.connected = false
}

func (p *user) isLogin() bool {
	return p.connected
}

func (p *user) setGameID(gameID uint32) {
	p.gameID = gameID
}

func (p *user) clearGameID() {
	p.gameID = 0
}

// GetAccount 获取账号.
func (p *user) GetAccount() string {
	return p.account
}

// SendMsg向玩家发送消息
func (p *user) SendMsg(message proto.Message) {
	p.SendUpdate()
	serverMgr.SendMsg("Send2Client", message, p.Agent)
}

func (p *user) Send2Gate(message proto.Message) {
	serverMgr.Send2Gate(message, p.Agent)
}

func (p *user) SendUpdate() {
	if p.updateNotify != nil {
		serverMgr.SendMsg("Send2Client", p.updateNotify, p.Agent)
		p.updateNotify = nil
	}
}

// ID 获取Uid
func (p *user) ID() uint64 {
	return p.info.User.UserID
}

func (p *user) notifyUpdate(typ cmsg.CNotifyDataChangeType, data interface{}) {
	if p.updateNotify == nil {
		p.updateNotify = &cmsg.CNotifyDataChange{}
	}
	switch typ {
	case cmsg.CNotifyDataChange_DCTUser:
		p.updateNotify.User = data.(*gamedef.User)
	case cmsg.CNotifyDataChange_DCTGeneral:
		p.updateNotify.Generals = append(p.updateNotify.Generals, data.([]*gamedef.General)...)
	default:
		return
	}

	p.updateNotify.Changes = append(p.updateNotify.Changes, typ)
}

func (p *user) UpdateData(typ cmsg.CNotifyDataChangeType, data interface{}) {
	p.notifyUpdate(typ, data)

	switch typ {
	case cmsg.CNotifyDataChange_DCTGeneral:
		p.info.Generals = p.general.toSlice()
	default:
		return
	}

	p.SaveUserDataDelay(time.Minute * 5)
}

func (p *user) SaveUserDataDelay(t time.Duration) {
	if t == 0 {
		if p.saveCancel != nil {
			p.saveCancel()
		}
		dbMgr.FlushUserAsync(p.info)
		return
	} else {
		if p.saveCancel != nil {
			return
		}
		p.saveCancel = AfterPost(t, func() {
			dbMgr.FlushUserAsync(p.info)
		})
	}
}

func (p *user) IsRobot() bool {
	return false
}

func (p *user) GetData() *gamedef.User {
	return p.info.User
}

func (p *user) UseItem(uint32) bool {
	return true
}

func (p *user) GetGeneral() *gamedef.General {
	g, exist := p.getFightGeneral()
	// 获取不到上阵武将说明进入游戏前逻辑有错误
	if !exist {
		panic("no user general data")
	}
	return g
}

func (p *user) onReqUserInit(req *cmsg.CReqUserInit) {
	resp := &cmsg.CRespUserInit{}

	if req.NickName == "" || len(req.NickName) > 20 {
		resp.ErrCode = uint32(emsg.BizErr_BE_NickNameInvalid)
		resp.ErrMsg = emsg.BizErr_BE_NickNameInvalid.String()
		p.SendMsg(resp)
		return
	}

	if p.info.User.Nickname != "" {
		resp.ErrCode = uint32(emsg.BizErr_BE_UserInitAlready)
		resp.ErrMsg = emsg.BizErr_BE_UserInitAlready.String()
		p.SendMsg(resp)
		return
	}

	// 检查初始
	flag := false
	for _, v := range cfgMgr.GetConfig().GetGlobalConfig().UserInitGeneral {
		if req.FirstGeneral == v {
			flag = true
		}
	}

	if !flag {
		resp.ErrCode = uint32(emsg.BizErr_BE_FirstGeneralInvalid)
		resp.ErrMsg = emsg.BizErr_BE_FirstGeneralInvalid.String()
		p.SendMsg(resp)
		return
	}

	// 可重名
	p.info.User.Nickname = req.NickName

	g, err := p.general.chooseGeneral(req.FirstGeneral)
	if err != nil {
		resp.ErrCode = uint32(emsg.BizErr_BE_FirstGeneralInvalid)
		resp.ErrMsg = emsg.BizErr_BE_FirstGeneralInvalid.String()
		p.SendMsg(resp)
		return
	}

	p.UpdateData(cmsg.CNotifyDataChange_DCTUser, p.info.User)
	p.UpdateData(cmsg.CNotifyDataChange_DCTGeneral, []*gamedef.General{g})

	p.SaveUserDataDelay(0)
	p.SendMsg(resp)
}

func (p *user) onReqNotifyUserData(req *cmsg.CReqNotifyUserData) {
	resp := &cmsg.CRespNotifyUserData{}

	resp.Generals = p.general.toSlice()
	p.SendMsg(resp)
}

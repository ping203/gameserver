package internal

import (
	"server/logs"

	"github.com/name5566/leaf/log"
	"github.com/wenxiu2199/gameserver/src/server/gameproto/cmsg"
	"github.com/wenxiu2199/gameserver/src/server/gameproto/gamedef"
)

func (p *Client) handler() {
	skeleton.RegisterHandlerClient(p.onRespAuth)
	skeleton.RegisterHandlerClient(p.onRespLogin)
	skeleton.RegisterHandlerClient(p.onRespUserInit)
	skeleton.RegisterHandlerClient(p.onRespStageFight)
	skeleton.RegisterHandlerClient(p.onRespNotifyUserData)
	skeleton.RegisterHandlerClient(p.onNotifyGameResult)
	skeleton.RegisterHandlerClient(p.onNotifyGameStage)
	skeleton.RegisterHandlerClient(p.onNotifyGameStart)
	skeleton.RegisterHandlerClient(p.onRespUseSkill)
}

func (p *Client) req() {
	p.reqAuth()
}

func (p *Client) reqAuth() {
	logs.Debug("=========reqAuth=========")
	p.WriteMsg(&cmsg.CReqAuth{
		Account:  "1",
		Password: "xxx",
	})
}

func (p *Client) onRespAuth(msg *cmsg.CRespAuth) {
	log.Debug("%v", msg)
	if msg.ErrCode == 0 {
		p.reqLogin(msg.Sign, msg.UserID)
	}
}

func (p *Client) reqLogin(sign string, userID uint64) {
	logs.Debug("=========reqLogin=========")
	p.WriteMsg(&cmsg.CReqLogin{
		Sign:   sign,
		UserID: userID,
	})
}

func (p *Client) onRespLogin(msg *cmsg.CRespLogin) {
	log.Debug("onRespLogin:%v", msg)
	p.userID = msg.User.UserID
	//p.reqUserInit()
	//p.reqNotifyUserData()
	p.reqStageFight()
}

func (p *Client) reqUserInit() {
	logs.Debug("=========reqUserInit=========")
	p.WriteMsg(&cmsg.CReqUserInit{
		NickName:     "asd",
		FirstGeneral: 1,
	})
}

func (p *Client) onRespUserInit(msg *cmsg.CRespUserInit) {
	log.Debug("onRespUserInit:%v", msg)
}

func (p *Client) reqNotifyUserData() {
	logs.Debug("=========reqNotifyUserData=========")
	p.WriteMsg(&cmsg.CReqNotifyUserData{})
}

func (p *Client) onRespNotifyUserData(msg *cmsg.CRespNotifyUserData) {
	log.Debug("onRespNotifyUserData:%v", msg)
}

func (p *Client) reqStageFight() {
	logs.Debug("=========reqStageFight=========")
	p.WriteMsg(&cmsg.CReqStageFight{})
}

func (p *Client) onRespStageFight(msg *cmsg.CRespStageFight) {
	log.Debug("onRespStageFight:%v", msg)
}

func (p *Client) onNotifyGameResult(msg *cmsg.CNotifyGameResult) {
	log.Debug("onNotifyGameResult:%v", msg)
	p.reqStageFight()
}

func (p *Client) onNotifyGameStage(msg *cmsg.CNotifyGameStage) {
	switch msg.Stage {
	case gamedef.GameStageTyp_GSTChoose:
		p.reqUseSkill()
	}
}

func (p *Client) onNotifyGameStart(msg *cmsg.CNotifyGameStart) {
	for _, v := range msg.Generals {
		if v.UserID == p.userID {
			p.general = &gamedef.General{
				Skills: v.Skills,
			}
			break
		}
	}
	log.Debug("%v", msg)
}

func (p *Client) reqUseSkill() {
	logs.Debug("=========reqUseSkill=========")
	p.WriteMsg(&cmsg.CReqUseSkill{
		SkillID: p.general.Skills[0],
	})
}

func (p *Client) onRespUseSkill(msg *cmsg.CRespUseSkill) {
	log.Debug("%v", msg)
}

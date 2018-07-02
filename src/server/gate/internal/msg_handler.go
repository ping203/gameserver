package internal

import (
	"server/logs"

	"github.com/name5566/leaf/gate"
	"github.com/wenxiu2199/gameserver/src/server/gameproto/cmsg"
	"github.com/wenxiu2199/gameserver/src/server/gameproto/emsg"
	"github.com/wenxiu2199/gameserver/src/server/gameproto/smsg"
)

func init() {
	skeleton.RegisterHandler(onReqAuth)
	skeleton.RegisterHandler(onReqLogin)
}

func onReqAuth(req *cmsg.CReqAuth, agent gate.Agent) {
	resp := &cmsg.CRespAuth{}
	session, exist := sessionMgr.getSessionByAgent(agent)
	if !exist {
		resp.ErrCode = uint32(emsg.LoginErr_LE_UnAuthenticated)
		resp.ErrMsg = emsg.LoginErr_LE_UnAuthenticated.String()
		writeMsg(agent, resp)
		return
	}

	if session.isAuthing() {
		resp.ErrCode = uint32(emsg.LoginErr_LE_Authenticating)
		resp.ErrMsg = emsg.LoginErr_LE_Authenticating.String()
		writeMsg(agent, resp)
		return
	}

	sessionMgr.addUserOnAuth(agent)
	serverMgr.Send2Login(&smsg.GtLsReqAuth{
		Account:  req.Account,
		Password: req.Password,
	}, agent)

}

func onReqLogin(req *cmsg.CReqLogin, agent gate.Agent) {
	resp := &cmsg.CRespLogin{}
	ses, exist := sessionMgr.getSessionByAgent(agent)
	if !exist {
		resp.ErrCode = 1
		writeMsg(agent, resp)
		return
	}

	if ses.userID != req.UserID {
		resp.ErrCode = 2
		writeMsg(agent, resp)
		return
	}

	ok := ses.checkSign(req.Sign)
	if !ok {
		resp.ErrCode = 3
		writeMsg(agent, resp)
		return
	}

	if ses.isLoging() {
		resp.ErrCode = 3
		writeMsg(agent, resp)
		return
	}

	sessionMgr.addUserOnLogin(agent)

	// todo 超时回调.
	requester.ReqTimeOut(func(seqID int64) error {
		return serverMgr.Send2Game(&smsg.GtGsReqLogin{
			SeqID:   seqID,
			UserID:  ses.userID,
			Account: ses.account,
		}, agent)
	}, func(data interface{}, err error) {
		if err != nil {
			resp.ErrCode = uint32(emsg.SystemErr_SE_Service)
			resp.ErrMsg = emsg.SystemErr_SE_Service.String()
			writeMsg(agent, resp)
			return
		}
		msg := data.(*smsg.GtGsRespLogin)
		if msg.ErrCode != 0 {
			resp.ErrMsg = msg.ErrMsg
			resp.ErrCode = msg.ErrCode
			writeMsg(agent, resp)
			return
		}
		_, exist := sessionMgr.getSessionByAgent(agent)
		if !exist {
			logs.Debug("返回时, 用户已经断开连接")
			return
		}

		sessionMgr.addUserOnLoginSuccess(agent)
		resp.UserID = msg.UserID
		resp.User = msg.User

		writeMsg(agent, resp)
	}, requesterTimeOut)
}

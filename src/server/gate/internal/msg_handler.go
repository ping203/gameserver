package internal

import (
	"server/gameproto/cmsg"
	"server/gameproto/emsg"
	"server/gameproto/smsg"

	"github.com/name5566/leaf/gate"
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
		agent.WriteMsg(resp)
		return
	}

	if session.isAuthing() {
		resp.ErrCode = uint32(emsg.LoginErr_LE_Authenticating)
		resp.ErrMsg = emsg.LoginErr_LE_Authenticating.String()
		agent.WriteMsg(resp)
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
		agent.WriteMsg(resp)
		return
	}

	if ses.userID != req.UserID {
		resp.ErrCode = 1
		agent.WriteMsg(resp)
		return
	}

	ok := ses.checkSign(req.Sign)
	if !ok {
		resp.ErrCode = 3
		agent.WriteMsg(resp)
		return
	}

	if ses.isLoging() {
		resp.ErrCode = 3
		agent.WriteMsg(resp)
		return
	}

	sessionMgr.addUserOnLogin(agent)

	// todo 超时回调.
	requester.ReqTimeOut(func(seqID int64) error {
		return serverMgr.Send2Game(&smsg.GtGsReqLogin{
			SeqID:   seqID,
			UserID:  req.UserID,
			Account: ses.account,
		}, agent)
	}, func(data interface{}, err error) {
		if err != nil {
			resp.ErrCode = uint32(emsg.SystemErr_SE_Service)
			resp.ErrMsg = emsg.SystemErr_SE_Service.String()
			agent.WriteMsg(resp)
			return

		}
		msg := data.(*smsg.GtGsRespLogin)
		if msg.ErrCode != 0 {
			resp.ErrMsg = msg.ErrMsg
			resp.ErrCode = msg.ErrCode
			agent.WriteMsg(resp)
			return
		}

		resp.UserID = msg.UserID
		resp.User = msg.User
	}, requesterTimeOut)
}

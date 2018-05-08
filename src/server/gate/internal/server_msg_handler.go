package internal

import (
	"server/gameproto/cmsg"
	"server/gameproto/smsg"
	"server/logs"

	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/gate"
)

func init() {
	// 消息分发
	skeleton.RegisterChanRPC("Send2Client", onSend2Client)
	skeleton.RegisterChanRPC("Send2Clients", onSend2Clients)
	// 具体逻辑
	skeleton.RegisterHandler(onGtLsRespAuth)
}

// 参数1 消息, 参数2 目标用户群
func onSend2Clients(args []interface{}) {
	msg := args[0].(proto.Message)
	users := args[1].([]uint64)

	for _, v := range users {
		session, exist := sessionMgr.getSessionByUserID(v)
		if !exist {
			continue
		}
		session.agent.WriteMsg(msg)
	}
}

// 参数1 消息, 参数2 用户
func onSend2Client(args []interface{}) {
	msg := args[0].(proto.Message)
	userID := args[1].(uint64)
	session, exist := sessionMgr.getSessionByUserID(userID)
	if !exist {
		return
	}
	session.agent.WriteMsg(msg)
}

func onGtLsRespAuth(req *smsg.GtLsRespAuth, agent gate.Agent) {
	resp := &cmsg.CRespAuth{}
	if req.ErrCode != 0 {
		resp.ErrCode = req.ErrCode
		resp.ErrMsg = req.ErrMsg
		agent.WriteMsg(resp)
		return
	}

	_, exist := sessionMgr.getSessionByAgent(agent)
	if !exist {
		logs.Debug("返回时, 用户已经断开连接")
		return
	}

	sign := sessionMgr.addUserOnAuthSuccess(agent, req.UserID, req.Account)
	resp.Sign = sign
	resp.UserID = req.UserID
	resp.Account = req.Account

	agent.WriteMsg(resp)

	_, exist = sessionMgr.getSessionByUserID(req.UserID)
	if exist {
		// 通知已经登录
	}

}

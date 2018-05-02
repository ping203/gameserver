package gate

import (
	"github.com/wenxiu2199/gameserver/src/server/msg"
	"github.com/wenxiu2199/gameserver/src/server/gameproto/cmsg"
)

// 初始化路由
func init() {
	bindClientMessage()
}

func bindClientMessage() {
	msg.Router(&cmsg.CReqAuth{},)
}
package internal

var sessionMgr *sessionManager

func init() {
	sessionMgr = &sessionManager{}
	sessionMgr.init()
}

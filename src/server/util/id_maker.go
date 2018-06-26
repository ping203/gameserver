package util

import (
	"sanguosha.com/games/sgs/framework/util"
)

var projectStartTime = uint64(util.GetCurrentMicroTimestamp()) - 1514736000*1e3 // 2018/1/1 00:00:00
const maxSeqID = 1000

var lastSeqID uint64 = 1

func getSeqID() uint64 {
	lastSeqID += 1
	return lastSeqID
}

func GeneratePKID() uint64 {
	id := projectStartTime + getSeqID()
	return id
}

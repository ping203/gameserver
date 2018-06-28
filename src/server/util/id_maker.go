package util

import (
	"github.com/wenxiu2199/gameserver/src/server/gameproto/gamedef"
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

func RandIndividual() *gamedef.Individual {
	rands := util.GetRandomN(32, 6)
	return &gamedef.Individual{
		Hp:        int32(rands[0]),
		Attack:    int32(rands[1]),
		Defense:   int32(rands[2]),
		SpAttack:  int32(rands[3]),
		SpDefense: int32(rands[4]),
		Speed:     int32(rands[5]),
	}
}

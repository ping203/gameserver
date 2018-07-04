package internal

import (
	"fmt"
	"math"

	"server/util"

	"github.com/sirupsen/logrus"
	"github.com/wenxiu2199/gameserver/src/server/gameproto/cmsg"
	"github.com/wenxiu2199/gameserver/src/server/gameproto/gamedef"
)

type general struct {
	*user

	pkID2General      map[uint64]*gamedef.General
	generalID2General map[uint32]*gamedef.General
}

func (p *general) init(user *user, generals []*gamedef.General) {
	p.user = user
	p.pkID2General = make(map[uint64]*gamedef.General)
	p.generalID2General = make(map[uint32]*gamedef.General)
	for _, v := range generals {
		p.setByPkID(v.PkID, v)
	}
	logrus.WithFields(logrus.Fields{
		"len": len(p.pkID2General),
	}).Debug("init general")
}

func (p *general) getByPkID(pkID uint64) (*gamedef.General, bool) {
	general, ok := p.pkID2General[pkID]
	return general, ok
}

func (p *general) setByPkID(pkID uint64, general *gamedef.General) {
	p.pkID2General[pkID] = general
	p.generalID2General[general.GeneralID] = general
}

func (p *general) delByPkID(pkID uint64) {
	if _, ok := p.pkID2General[pkID]; ok {
		delete(p.pkID2General, pkID)
	}
}

func (p *general) addGeneral(cp *gamedef.General) *gamedef.General {
	general := &gamedef.General{
		PkID:       util.GeneratePKID(),
		GeneralID:  cp.GeneralID,
		Individual: cp.Individual,
		Effort:     cp.Effort,
		Skills:     cp.Skills,
		Level:      cp.Level,
	}
	p.setByPkID(general.PkID, general)
	return general
}

func (p *general) getFightGeneral() (*gamedef.General, bool) {
	g, exist := p.pkID2General[p.info.User.FightGeneralID]
	return g, exist
}

func (p *general) chooseGeneral(generalID uint32) (*gamedef.General, error) {
	conf, exist := cfgMgr.GetConfig().GetGeneralConfByGeneralID(generalID)
	if !exist {
		return nil, fmt.Errorf("general.chooseGeneral GetGeneralConfByGeneralID %v", generalID)
	}
	general := &gamedef.General{
		GeneralID:  generalID,
		Individual: util.RandIndividual(),
		Effort:     &gamedef.Individual{},
		Skills:     conf.BaseSkills,
		Level:      5,
	}
	g := p.addGeneral(general)

	p.user.info.User.FightGeneralID = g.PkID
	return g, nil
}

func (p *general) toSlice() []*gamedef.General {
	generals := make([]*gamedef.General, 0, len(p.generalID2General))
	for _, v := range p.generalID2General {
		generals = append(generals, v)
	}

	return generals
}

func (p *general) addExp(pkID uint64, exp int32) {
	general, exist := p.getByPkID(pkID)
	if !exist {
		return
	}

	general.Exp += exp
	general.Level = uint32(math.Sqrt(float64(general.Exp / 100)))
	p.user.UpdateData(cmsg.CNotifyDataChange_DCTGeneral, []*gamedef.General{general})
}

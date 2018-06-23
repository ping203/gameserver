package internal

import (
	"server/gameproto/gamedef"
	"server/util"

	"github.com/sirupsen/logrus"
)

type general struct {
	*user

	pkID2General map[uint64]*gamedef.General
}

func (p *general) init(user *user, generals []*gamedef.General) {
	p.user = user
	p.pkID2General = make(map[uint64]*gamedef.General)
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
}

func (p *general) delByPkID(pkID uint64) {
	if _, ok := p.pkID2General[pkID]; ok {
		delete(p.pkID2General, pkID)
	}
}

func (p *general) addGeneral(pkID uint64, cp *gamedef.General) {
	general := &gamedef.General{
		PkID:       util.GeneratePKID(),
		GeneralID:  cp.GeneralID,
		Individual: cp.Individual,
		Effort:     cp.Effort,
		Skills:     cp.Skills,
	}
	p.setByPkID(general.PkID, general)
}

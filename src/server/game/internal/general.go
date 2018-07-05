package internal

import (
	"fmt"
	"math"

	"server/gameproto/emsg"
	"server/util"

	"github.com/sirupsen/logrus"
	"github.com/wenxiu2199/gameserver/src/server/gameproto/cmsg"
	"github.com/wenxiu2199/gameserver/src/server/gameproto/gamedef"
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
	p.UpdateGeneral(general)
	return general
}

func (p *general) getFightGeneral() (*gamedef.General, bool) {
	g, exist := p.pkID2General[p.info.User.FightGeneralID]
	return g, exist
}

func (p *general) toSlice() []*gamedef.General {
	generals := make([]*gamedef.General, 0, len(p.pkID2General))
	for _, v := range p.pkID2General {
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
	p.user.UpdateGeneral(general)
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

func (p *general) onReqLearnSkill(req *cmsg.CReqLearnSkill) {
	resp := &cmsg.CRespLearnSkill{}
	general, exist := p.getByPkID(req.GeneralPkID)
	if !exist {
		resp.ErrCode = uint32(emsg.BizErr_BE_GeneralPkID)
		resp.ErrMsg = emsg.BizErr_BE_GeneralPkID.String()
		p.SendMsg(resp)
		return
	}

	conf, exist := cfgMgr.GetConfig().GetGeneralConfByGeneralID(general.GeneralID)
	if !exist {
		resp.ErrCode = uint32(emsg.BizErr_BE_GeneralConf)
		resp.ErrMsg = emsg.BizErr_BE_GeneralConf.String()
		p.SendMsg(resp)
		return
	}

	if req.Position > 3 {
		resp.ErrCode = uint32(emsg.BizErr_BE_CanNotLearn)
		resp.ErrMsg = emsg.BizErr_BE_CanNotLearn.String()
		p.SendMsg(resp)
		return
	}

	// 检查是否已经学习
	for _, v := range general.Skills {
		if v == req.SkillID {
			resp.ErrCode = uint32(emsg.BizErr_BE_AlreadyLearn)
			resp.ErrMsg = emsg.BizErr_BE_AlreadyLearn.String()
			p.SendMsg(resp)
			return
		}
	}

	learn := false
	for _, v := range conf.LearnSkills {
		if v.SkillID == req.SkillID {
			if general.Level >= v.SkillID {
				if len(general.Skills) >= 3 {
					general.Skills[req.Position] = req.SkillID
				} else {
					general.Skills = append(general.Skills, req.SkillID)
				}
				learn = true
				break
			}
		}
	}

	if !learn {
		resp.ErrCode = uint32(emsg.BizErr_BE_CanNotLearn)
		resp.ErrMsg = emsg.BizErr_BE_CanNotLearn.String()
		p.SendMsg(resp)
		return
	}

	p.UpdateGeneral(general)
	p.SendMsg(resp)
}

func (p *general) onReqSwitchGeneral(req *cmsg.CReqSwitchGeneral) {
	resp := &cmsg.CRespSwitchGeneral{}
	_, exist := p.getByPkID(req.GeneralPKID)
	if !exist {
		resp.ErrCode = uint32(emsg.BizErr_BE_GeneralPkID)
		resp.ErrMsg = emsg.BizErr_BE_GeneralPkID.String()
		p.SendMsg(resp)
		return
	}

	p.user.info.User.FightGeneralID = req.GeneralPKID
	p.UpdateUser(p.user.info.User)
	p.SendMsg(resp)
}

package sixsweep

import "server/gameproto/gamedef"

type chess struct {
	*gamedef.Chess
	adjacentChess [side]*chess
}

func newChess() *chess {
	c := &chess{
		Chess: &gamedef.Chess{
			ChessType: gamedef.ChessTyp_CTBlank,
		},
		adjacentChess: [side]*chess{},
	}
	return c
}

// IsMine 是否是雷
func (p *chess) IsMine() bool {
	return p.IsMine()
}

// GetCount 获取周围雷数量
func (p *chess) GetCount() uint32 {
	var count uint32 = 0
	for _, v := range p.adjacentChess {
		if v == nil {
			continue
		}
		if v.ChessType == gamedef.ChessTyp_CTMine {
			count++
		}
	}
	return count
}

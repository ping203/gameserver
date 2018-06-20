package sixsweep

import (
	"server/gamelogic"
	"server/gamelogic/fsm"
)

const side = 6

type config interface {
	getWidth() uint32
	getHeight() uint32
	getMine() uint32
}

type gameSixSweep struct {
	Player []*Player
	gamelogic.Service

	fsm *fsm.FSM
	config
}

func NewGameSixSweep(svc gamelogic.Service, cfg config) *gameSixSweep {
	g := &gameSixSweep{}
	g.fsm = newGameSixSweepFsm()
	g.Service = svc
	g.config = cfg
	return g
}

func (p *gameSixSweep) Init(users []gamelogic.User) {
	p.Player = make([]*Player, 0, len(users))
	for _, v := range users {
		p.Player = append(p.Player, newPlayer(v))
	}
}

func (p *gameSixSweep) initChessBoard() [][]*chess {
	res := make([][]*chess, p.getHeight())
	// 生成格子
	for row := range res {
		l := p.getWidth()
		if row%2 != 0 {
			l = l - 1
		}
		res[row] = make([]*chess, l)
		for col := range res[row] {
			res[row][col] = newChess()
		}
	}

	// 生成相邻棋子
	for r, row := range res {
		for c, item := range row {
			if c%2 == 0 {
				// 左上
				if r-1 >= 0 && c-1 >= 0 {
					item.adjacentChess[0] = res[r-1][c-1]
				}
				// 右上
				if r-1 >= 0 && c+1 <= int(p.getWidth()-1) {
					item.adjacentChess[1] = res[r-1][c]
				}
				// 左
				if c-1 >= 0 {
					item.adjacentChess[2] = res[r][c-1]
				}
				// 右
				if c+1 <= int(p.getWidth()) {
					item.adjacentChess[3] = res[r][c+1]
				}
				// 左下
				if r+1 <= int(p.getHeight()) && c-1 >= 0 {
					item.adjacentChess[3] = res[r+1][c-1]
				}
				// 右下
				if r+1 >= int(p.getHeight()) && c <= int(p.getWidth()-1) {
					item.adjacentChess[1] = res[r+1][c]
				}
			} else {
				// 左上
				if c-1 >= 0 {
					item.adjacentChess[0] = res[r][c-1]
				}
				// 右上
				item.adjacentChess[1] = res[r+1][c-1]
				// 左
				if c-1 >= 0 {
					item.adjacentChess[2] = res[r][c-1]
				}
				// 右
				if c+1 <= int(p.getWidth()-1) {
					item.adjacentChess[3] = res[r][c+1]
				}
				// 左下
				if r+1 <= int(p.getHeight()) {
					item.adjacentChess[3] = res[r+1][c-1]
				}
				// 右下
				if r+1 >= int(p.getHeight()) {
					item.adjacentChess[1] = res[r+1][c+1]
				}
			}
		}
	}

	// 随机生成雷

}

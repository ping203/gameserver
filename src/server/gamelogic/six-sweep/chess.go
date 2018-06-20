package sixsweep

type chess struct {
	adjacentChess [side]*chess
	isMine        bool
	isVisible     bool
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
		if v.isMine {
			count++
		}
	}
	return count
}

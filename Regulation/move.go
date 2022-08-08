package amazonsChess

type ChessMove struct {
	Start    int
	End      int
	Obstacle int
}

func (m ChessMove) GetVal() []int {
	return []int{m.Start, m.End, m.Obstacle}
}

func (m ChessMove) Equal(move ChessMove) bool {
	if m.Start == move.Start && m.End == move.End && m.Obstacle == move.Obstacle {
		return true
	}
	return false
}

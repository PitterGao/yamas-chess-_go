package Mcts

import (
	amazonsChess "github.com/PitterGao/Regulation"
)

type Node struct {
	Parent      *Node
	Children    []*Node
	Move        amazonsChess.ChessMove
	Moves       []amazonsChess.ChessMove
	State       *amazonsChess.State
	Visits      int
	Expand      bool
	FullyExpand bool
	Eva         float64
}

package Mcts

import (
	"errors"
	amazonsChess "github.com/PitterGao/Regulation"
	"log"
	"math"
	"math/rand"
	"sort"
	"time"
)

type Mcts struct {
	tree   *Node
	deepth int
}

var turns = 0

func NewMcts(tree *Node) (*Mcts, error) {
	if tree == nil {
		return nil, errors.New("error")
	}
	return &Mcts{tree: tree}, nil
}

func isTrue(s *amazonsChess.State, move amazonsChess.ChessMove) bool {
	//判断是否可行
	end, err := s.GetActionSpace(move.Start)
	if err != nil {
		log.Fatal(err)
	}
	for _, ends := range end {
		if move.End == ends {
			board := make([]int, 100)
			_ = copy(board, s.Board)
			board[move.Start] = 0
			board[move.End] = s.CurrentPlayer
			ss := amazonsChess.State{
				Board:         board,
				CurrentPlayer: s.CurrentPlayer,
			}
			ob, err := ss.GetActionSpace(move.End)
			if err != nil {
				log.Fatal(err)
			}
			for _, obs := range ob {
				if move.Obstacle == obs {
					return true
				}
			}
		}
	}
	return false
}

func AI(s *amazonsChess.State) amazonsChess.ChessMove {
	//固定棋谱
	//先手
	var loop int
	if s.CurrentPlayer == 1 && turns == 0 {
		turns += 1
		return amazonsChess.ChessMove{
			Start:    93,
			End:      23,
			Obstacle: 41,
		}
	}
	if s.CurrentPlayer == 1 && turns == 1 && s.Board[30] == -1 {
		move := amazonsChess.ChessMove{
			Start:    69,
			End:      36,
			Obstacle: 31,
		}
		if isTrue(s, move) {
			turns++
			return move
		}
	}
	//后手
	if s.CurrentPlayer == -1 && turns == 0 && s.Board[96] == 1 {
		move := amazonsChess.ChessMove{
			Start:    30,
			End:      85,
			Obstacle: 86,
		}
		if isTrue(s, move) {
			turns++
			return move
		}
	}
	if s.CurrentPlayer == -1 && turns == 0 && s.Board[93] == 1 {
		move := amazonsChess.ChessMove{
			Start:    39,
			End:      84,
			Obstacle: 83,
		}
		if isTrue(s, move) {
			turns++
			return move
		}
	}
	if s.CurrentPlayer == -1 && turns == 0 && s.Board[69] == 1 {
		move := amazonsChess.ChessMove{
			Start:    6,
			End:      76,
			Obstacle: 58,
		}
		if isTrue(s, move) {
			turns++
			return move
		}
	}
	if s.CurrentPlayer == -1 && turns == 0 && s.Board[60] == 1 {
		move := amazonsChess.ChessMove{
			Start:    3,
			End:      73,
			Obstacle: 51,
		}
		if isTrue(s, move) {
			turns++
			return move
		}
	}
	if s.CurrentPlayer == -1 && turns == 0 && s.Board[96] == 1 {
		move := amazonsChess.ChessMove{
			Start:    6,
			End:      66,
			Obstacle: 63,
		}
		if isTrue(s, move) {
			turns++
			return move
		}
	}
	if s.CurrentPlayer == -1 && turns == 0 && s.Board[93] == 1 {
		move := amazonsChess.ChessMove{
			Start:    3,
			End:      63,
			Obstacle: 66,
		}
		if isTrue(s, move) {
			turns++
			return move
		}
	}
	root := &Node{
		Children:    make([]*Node, 0),
		State:       s,
		FullyExpand: false,
	}
	m, err := NewMcts(root)
	if err != nil {
		log.Fatal(err)
	}
	if turns <= 15 {
		loop = 650000
	} else {
		loop = 800000
	}
	bestChild := m.getBestChild(s, loop)
	turns++
	return bestChild.Move
}

func (m *Mcts) getBestChild(s *amazonsChess.State, loop int) *Node {
	for i := 0; i < loop; i++ {
		m.UCTsearch(s)
	}
	var idx int
	for i := 0; i < len(m.tree.Children); i++ {
		if m.tree.Children[i].Visits > m.tree.Children[idx].Visits {
			idx = i
		} else if m.tree.Children[i].Visits == m.tree.Children[idx].Visits && m.tree.Children[i].Eva > m.tree.Children[idx].Eva {
			idx = i
		}
	}
	return m.tree.Children[idx]
}

func (m *Mcts) UCTsearch(s *amazonsChess.State) {
	var n *Node
	var newNode *Node
	n = m.tree
	length := len(n.Children)
	if length >= len(n.State.GetValid()) {
		n.FullyExpand = true
	}
	for n.FullyExpand && n.State.GameOver() == 0 {
		//进行正向剪枝
		if length > 100 && n.Parent == nil {
			sort.Slice(n.Children, func(i, j int) bool {
				return n.Children[i].Eva > n.Children[j].Eva
			})
			n.Children = n.Children[:100]
		}
		if length > 200 && n.Parent != nil && n.State.CurrentPlayer == s.CurrentPlayer {
			sort.Slice(n.Children, func(i, j int) bool {
				return n.Children[i].Eva > n.Children[j].Eva
			})
			n.Children = n.Children[:200]
		}
		if length > 200 && n.Parent != nil && n.State.CurrentPlayer == -s.CurrentPlayer {
			sort.Slice(n.Children, func(i, j int) bool {
				return n.Children[i].Eva < n.Children[j].Eva
			})
			n.Children = n.Children[:200]
		}
		n = m.Select(n, s)
	}
	newNode = m.expand(n, s)
	m.backupdate(newNode)

}

func (m *Mcts) Select(n *Node, s *amazonsChess.State) *Node {
	//选择
	var c *Node
	var ret *Node
	var curScore float64
	var bestScore float64 = -786554453
	for i := 0; i < len(n.Children); i++ {
		c = n.Children[i]
		if s.CurrentPlayer == n.State.CurrentPlayer {
			curScore = c.Eva/float64(c.Visits) + 1.4*math.Sqrt(math.Log(float64(n.Visits))/float64(c.Visits))
		} else {
			curScore = -c.Eva/float64(c.Visits) + 1.4*math.Sqrt(math.Log(float64(n.Visits))/float64(c.Visits))
		}

		if curScore > bestScore {
			ret = c
			bestScore = curScore
		}
	}
	return ret
}

func (m *Mcts) expand(n *Node, s *amazonsChess.State) *Node {
	//扩展
	var num = 0
	var sol amazonsChess.ChessMove
	if n == nil {
		return n
	}
	if !n.Expand {
		get := n.State.GetValid()
		if len(get) == 0 {
			return n
		}
		n.Moves = get
		n.Expand = true
	}
	k := len(n.Moves)
	if k == 0 {
		return n
	}
	rand.Seed(time.Now().UnixNano())
	num = rand.Intn(k)
	sol = n.Moves[num]
	n.Moves = append(n.Moves[:num], n.Moves[num+1:]...)

	State, _ := n.State.StateMove(sol)
	newNode := &Node{
		Parent:      n,
		State:       State,
		Move:        sol,
		FullyExpand: false,
	}
	newNode.Eva = m.rollout(newNode, s)
	n.Children = append(n.Children, newNode)
	return newNode
}

func (m *Mcts) backupdate(n *Node) *Node {
	//回溯
	if n == nil {
		return nil
	}
	n.Visits++
	for n.Parent != nil {
		wins := n.Eva
		n = n.Parent
		n.Visits++
		n.Eva += wins
	}
	return n
}

func (m *Mcts) rollout(n *Node, s *amazonsChess.State) float64 {
	//模拟
	var value float64 = 0
	res := &Node{
		State: n.State,
	}
	Sim := 0
	if n.State.CurrentPlayer == s.CurrentPlayer {
		if turns <= 20 {
			Sim = 2
		} else {
			Sim = 4
		}
	} else {
		if turns <= 20 {
			Sim = 1
		} else {
			Sim = 3
		}
	}
	for i := 0; i < Sim; i++ {
		if res.State.GameOver() == 0 {
			res.State, _ = res.State.RandomMove()
		}
	}
	value = evaluation(res)
	if s.CurrentPlayer == -1 {
		return value
	}
	return -value
}

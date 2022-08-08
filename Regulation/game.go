package amazonsChess

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
	"strconv"
	"time"
)

var num, num2 int
var str = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
var tt time.Duration = 0

type Game struct {
	CurrentPlayer int                    `json:"current_player,omitempty"`
	CurrentState  *State                 `json:"current_state,omitempty"`
	Winner        int                    `json:"winner,omitempty"`
	Ai1Handler    func(*State) ChessMove `json:"ai_1_handler,omitempty"`
	Ai2Handler    func(*State) ChessMove `json:"ai_2_handler,omitempty"`
}

func (g *Game) Reset(currentPlayer int) error {
	if currentPlayer != -1 && currentPlayer != 1 {
		return errors.New("wrong currentPlayer(need -1 or 1)")
	}
	g.CurrentPlayer = currentPlayer
	g.CurrentState = &State{
		Board:         NewBoard(),
		CurrentPlayer: currentPlayer,
	}
	g.Winner = 0
	return nil
}

func (g *Game) GetMove(state *State) ChessMove {
	if g.CurrentPlayer == -1 {
		if g.Ai1Handler == nil {
			return g.PersonMove(state)
		}
		return g.Ai1Handler(state)
	} else {
		if g.Ai2Handler == nil {
			return g.PersonMove(state)
		}
		return g.Ai2Handler(state)
	}
}

func (g *Game) Start(isShow bool) {
	var err error

	err = g.Reset(g.CurrentPlayer)
	if err != nil {
		log.Fatal(err)
	}
	for g.CurrentState.GameOver() == 0 {
		var err error
		t := time.Now()
		move := g.GetMove(g.CurrentState)
		elapsed := time.Now().Sub(t)
		fmt.Println("该函数执行完成耗时：", elapsed)
		if g.CurrentPlayer == 1 {
			tt += elapsed
		}
		fmt.Println("我们一共耗费：", tt)
		fmt.Println(move.Start, move.End, move.Obstacle)
		if move.Equal(ChessMove{}) {
			g.CurrentState, _ = g.CurrentState.RandomMove()
		} else {
			g.CurrentState, err = g.CurrentState.StateMove(move)
			if err != nil {
				log.Fatal(err)
			}
		}
		g.CurrentPlayer = g.CurrentState.CurrentPlayer
		num++
		rows := strconv.Itoa(10 - move.Start/10)
		cols := str[move.Start%10]
		rowe := strconv.Itoa(10 - move.End/10)
		cole := str[move.End%10]
		rowo := strconv.Itoa(10 - move.Obstacle/10)
		colo := str[move.Obstacle%10]
		filePath := "E:\\深度学习\\第一版 亚马逊棋(mcts)\\第二版AmazonsChess\\Run\\棋谱.txt"
		file, err := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE|os.O_APPEND, os.ModeAppend|os.ModePerm)
		//file, err := os.Create(filePath)
		if err != nil {
			fmt.Println("打开文件失败", err)
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(file)
		write := bufio.NewWriter(file)
		if num%2 != 0 {
			num2++
			num1 := strconv.Itoa(num2)
			write.WriteString(num1)
			write.WriteString(" ")
		}
		write.WriteString(cols)
		write.WriteString(rows)
		write.WriteString(cole)
		write.WriteString(rowe)
		write.WriteString("(")
		write.WriteString(colo)
		write.WriteString(rowo)
		write.WriteString(")")
		write.WriteString(" ")
		if num%2 == 0 {
			write.WriteString("\n")
		}
		write.Flush()
		if isShow {
			//fmt.Print("\x1b8")
			//fmt.Print("\x1b[2k") // 清空当前行的内容 擦除线<ESC> [2K
			g.CurrentState.PrintState()
			time.Sleep(50 * time.Millisecond)
		}
	}

	var playerStr string
	g.Winner = g.CurrentState.GameOver()
	if g.Winner == 1 {
		playerStr = color.New(color.FgHiRed).Sprintf("red")
	} else {
		playerStr = color.New(color.FgHiBlue).Sprintf("blue")
	}
	fmt.Printf("winner is: %s\n", playerStr)
}

func NewBoard() []int {
	board := make([]int, 100)
	board[3] = -1
	board[6] = -1
	board[30] = -1
	board[39] = -1
	board[60] = 1
	board[69] = 1
	board[93] = 1
	board[96] = 1
	return board
}

func (g *Game) PersonMove(state *State) ChessMove {
	for {
		var start int
		var end int
		var obstacle int
		var count int
		fmt.Println("请输入落子——起点，终点，障碍: ")
		_, err := fmt.Scanf("%d %d %d\n", &start, &end, &obstacle)
		if err != nil {
			continue
		}
		//判断起点合法
		for i := 0; i < 100; i++ {
			if state.CurrentPlayer == state.Board[i] && start == i {
				count++
			}
		}
		if count != 1 {
			continue
		}
		//判断终点合法
		endspace, err := state.GetActionSpace(start)
		if err != nil {
			continue
		}
		for _, i := range endspace {
			if i == end {
				count++
			}
		}
		if count != 2 {
			continue
		}
		//判断障碍合法
		board := make([]int, 100)
		_ = copy(board, state.Board)
		board[start] = 0
		board[end] = state.CurrentPlayer
		s := State{
			Board:         board,
			CurrentPlayer: state.CurrentPlayer,
		}
		obstacleSpace, err := s.GetActionSpace(end)
		if err != nil {
			continue
		}
		for _, i := range obstacleSpace {
			if i == obstacle {
				count++
			}
		}
		if count != 3 {
			continue
		}
		personMove := ChessMove{
			Start:    start,
			End:      end,
			Obstacle: obstacle,
		}
		return personMove
	}
}

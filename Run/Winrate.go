package main

import (
	"fmt"
	Mcts "github.com/PitterGao/MCTS"
	amazonsChess "github.com/PitterGao/Regulation"
	"time"
)

type winrate struct {
	WinsRed  float64
	WinsBlue float64
	RateRed  float64
	RateBlue float64
}

func (w *winrate) winsPrint() {
	fmt.Println("winsRed: ", w.WinsRed)
	fmt.Println("winsBlue: ", w.WinsBlue)
	fmt.Println("Winning rate of Red: ", w.RateRed*100, "%")
	fmt.Println("Winning rate of Blue:", w.RateBlue*100, "%")
}

func (w *winrate) wins(g int) {
	if g == 1 {
		w.WinsRed++
	} else {
		w.WinsBlue++
	}
}

func (w *winrate) Rate() {
	w.RateRed = w.WinsRed / (w.WinsRed + w.WinsBlue)
	w.RateBlue = w.WinsBlue / (w.WinsRed + w.WinsBlue)
}

func Ai1Handler(s *amazonsChess.State) amazonsChess.ChessMove {
	move := Mcts.AI(s)
	return move
}

func Ai2Handler(s *amazonsChess.State) amazonsChess.ChessMove {
	move := Mcts.AI(s)
	return move
}

func (w *winrate) play(t int) {
	for i := 0; i < t; i++ {
		g := amazonsChess.Game{
			CurrentPlayer: 1,
			//后手
			//Ai2Handler:    nil,
			//Ai1Handler:    Ai1Handler,
			//先手
			Ai1Handler: nil,
			Ai2Handler: Ai2Handler,
		}

		t := time.Now()
		g.Start(true)
		elapsed := time.Now().Sub(t)
		fmt.Println("该函数执行完成耗时：", elapsed)

		fmt.Printf("winner is: %v\n", g.Winner)
		if g.Ai1Handler != nil {
			fmt.Println("Ai1(blue) is running")
		}
		if g.Ai2Handler != nil {
			fmt.Println("Ai2(red) is running")
		}
		w.wins(g.Winner)
		w.Rate()
		w.winsPrint()
	}
}

func main() {
	w := winrate{}
	w.play(1)
}

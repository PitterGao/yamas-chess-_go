package Mcts

import "math"

var D = [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
var k1 = [3]float64{0.37, 0.25, 0.10}
var k2 = [3]float64{0.14, 0.30, 0.80}
var k3 = [3]float64{0.13, 0.20, 0.05}
var k4 = [3]float64{0.13, 0.20, 0.05}
var TmpChess1 [10][10]int
var TmpChess2 [10][10]int
var head int
var tail int

type node1 [10][10]struct {
	kb int
	kr int
	qb int
	qr int
}

var Map node1

type node2 [110]struct {
	x    int
	y    int
	step int
}

var que node2

func Queen(x int, n *Node) {
	tx := 0
	ty := 0
	head = 0
	tail = 0
	que[tail].x = x / 10
	que[tail].y = x % 10
	que[tail].step = 0
	tail++
	TmpChess2[x/10][x%10] = 1
	for head < tail {
		for k := 0; k < 8; k++ {
			tx = que[head].x
			ty = que[head].y
			for {
				tx = tx + D[k][0]
				ty = ty + D[k][1]
				if tx >= 0 && tx < 10 && ty >= 0 && ty < 10 && TmpChess2[tx][ty] == -1 {
					continue
				}
				if tx >= 0 && tx < 10 && ty >= 0 && ty < 10 && TmpChess2[tx][ty] == 0 {
					TmpChess2[tx][ty] = -1
					que[tail].x = tx
					que[tail].y = ty
					que[tail].step = que[head].step + 1
					tail++
				} else {
					break
				}
			}
		}
		head++
	}
	for k := 1; k < tail; k++ {
		if n.State.Board[x] == -1 && Map[que[k].x][que[k].y].qb > que[k].step {
			Map[que[k].x][que[k].y].qb = que[k].step
		}
		if n.State.Board[x] == 1 && Map[que[k].x][que[k].y].qr > que[k].step {
			Map[que[k].x][que[k].y].qr = que[k].step
		}
	}
}

func King(y int, n *Node) {
	tx := 0
	ty := 0
	head = 0
	tail = 0
	que[tail].x = y / 10
	que[tail].y = y % 10
	que[tail].step = 0
	tail++
	for head < tail {
		for k := 0; k < 8; k++ {
			tx = que[head].x + D[k][0]
			ty = que[head].y + D[k][1]
			if tx < 0 || tx > 9 || ty < 0 || ty > 9 || TmpChess1[tx][ty] != 0 {
				continue
			} else {
				TmpChess1[tx][ty] = 1
				que[tail].x = tx
				que[tail].y = ty
				que[tail].step = que[head].step + 1
				tail++
			}
		}
		head++
	}
	for k := 1; k < tail; k++ {
		if n.State.Board[y] == -1 && Map[que[k].x][que[k].y].kb > que[k].step {
			Map[que[k].x][que[k].y].kb = que[k].step
		}
		if n.State.Board[y] == 1 && Map[que[k].x][que[k].y].kr > que[k].step {
			Map[que[k].x][que[k].y].kr = que[k].step
		}
	}

}

func evaluation(n *Node) float64 {
	var queen float64 = 0
	var king float64 = 0
	var p1, p2 float64 = 0, 0
	var value float64
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			Map[i][j].kr = 10000
			Map[i][j].qr = 10000
			Map[i][j].kb = 10000
			Map[i][j].qb = 10000
		}
	}

	for i := 0; i < 100; i++ {
		if n.State.Board[i] == 1 || n.State.Board[i] == -1 {
			for j := 0; j < 100; j++ {
				if n.State.Board[j] == 0 {
					TmpChess1[j/10][j%10] = 0
					TmpChess2[j/10][j%10] = 0
				} else {
					TmpChess1[j/10][j%10] = 1
					TmpChess2[j/10][j%10] = 1
				}
			}
			Queen(i, n)
			King(i, n)
		}
	}

	for i := 0; i < 100; i++ {
		if n.State.Board[i] == 0 {
			if Map[i/10][i%10].kb == Map[i/10][i%10].kr && Map[i/10][i%10].kr != 10000 {
				if n.State.CurrentPlayer == -1 {
					king += 0.227
				} else {
					king -= 0.227
				}
			}
			if Map[i/10][i%10].kb < Map[i/10][i%10].kr {
				king++
			}
			if Map[i/10][i%10].kb > Map[i/10][i%10].kr {
				king--
			}
			if Map[i/10][i%10].qb == Map[i/10][i%10].qr && Map[i/10][i%10].qr != 10000 {
				if n.State.CurrentPlayer == -1 {
					queen += 0.227
				} else {
					queen -= 0.227
				}
			}
			if Map[i/10][i%10].qb < Map[i/10][i%10].qr {
				queen++
			}
			if Map[i/10][i%10].qb > Map[i/10][i%10].qr {
				queen--
			}
			//	position参数
			p1 = p1 + 2*(1.0/float64(Pow2(Map[i/10][i%10].qb))-1.0/float64(Pow2(Map[i/10][i%10].qr)))
			p2 = p2 + math.Min(1.0, math.Max(-1.0, (1.0/6.0)*float64(Map[i/10][i%10].kr-Map[i/10][i%10].kb)))
		}
	}
	if turns <= 20 {
		value = k1[0]*king + k2[0]*queen + k3[0]*p1 + k4[0]*p2
	} else if turns <= 40 {
		value = k1[1]*king + k2[1]*queen + k3[1]*p1 + k4[1]*p2
	} else {
		value = k1[2]*king + k2[2]*queen + k3[2]*p1 + k4[2]*p2
	}
	if turns < 25 {
		return queen + king
	}
	return value
}

func Pow2(num int) int {
	if num >= 31 {
		return 2147483647
	}
	return 1 << num
}

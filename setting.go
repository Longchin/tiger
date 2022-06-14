package golang

type Card struct {
	suits string
}

type WinInfo struct {
	Winner     string
	ComCard    string
	PlayerCard string
}
type Player struct {
	Money  int
	Bet    int
	points int
}
type Bingo struct {
	mul   int
	Bet   int
	five  []int
	four  []int
	three []int
}
type HitRate struct {
	Ng   int
	Fgt5 int
	Fgt6 int
	Fgt7 int
	Fg5  int
	Fg6  int
	Fg7  int
	r1   int
	r2   int
	r3   int
	r4   int
}

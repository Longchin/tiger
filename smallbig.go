package golang

// func SetBet() int {
// 	//玩家確認下注量100,200,500,1000
// 	// var b int
// 	fmt.Println("可下注額為:100,200,500,1000")
// 	// fmt.Print("請玩家確認下注量:")
// 	// fmt.Scanln(&b)
// 	// p.Bet = b
// 	p.Bet = 100
// 	p.Money = 10000000
// 	for i := 0; i < 500; i++ {
// 		p.Money -= p.Bet

// 	}
// 	return p.Bet
// }

var p Player

func Gamble(C int, r int) {
	//玩家確認下注量100,200,500,1000
	// var b int
	//fmt.Println("可下注額為:100,200,500,1000")
	// fmt.Print("請玩家確認下注量:")
	// fmt.Scanln(&b)
	// p.Bet = b
	p.Bet = C
	//一下注就減去下注量
	p.Money -= p.Bet
	//判斷贏分
	switch p.Bet {
	case 100:
		if r == 1 {
			p.points = 100 * 1.98
		} else {
			p.points = 0
		}

	case 200:
		if r == 1 {
			p.points = 200 * 1.98
		} else {
			p.points = 0
		}

	case 500:
		if r == 1 {
			p.points = 500 * 1.98
		} else {
			p.points = 0
		}
	case 1000:
		if r == 1 {
			p.points = 1000 * 1.98
		} else {
			p.points = 0
		}
	}
	p.Money += p.points
	//fmt.Println(p.Money)

}

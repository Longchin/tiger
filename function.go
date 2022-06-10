package golang

import (
	"math/rand"
	"time"
)

//產牌組
// func GetCards() []Card {
// 	cards := []Card{}
// 	var c Card
// 	for i := 1; i < 14; i++ {
// 		for j := 1; j < 5; j++ {
// 			c.number = i
// 			switch j {
// 			case 1:
// 				c.suits = "A"
// 			case 2:
// 				c.suits = "B"
// 			case 3:
// 				c.suits = "C"
// 			case 4:
// 				c.suits = "D"
// 			}
// 			cards = append(cards, c)

// 		}

// 	}
// 	return cards
// }

// // 抽4張牌
// func GetTakeCards(c []Card) (result []Card) {
// 	rand.Seed(time.Now().UnixNano())
// 	for i := 0; i < 4; i++ {
// 		result = append(result, c[rand.Intn(len(c))])
// 	}

// 	newSlice := []Card{} //排重
// 	copy(newSlice, result)
// 	wetherBreak := true
// 	for {
// 		for i, c1 := range result {
// 			for j, c2 := range newSlice {
// 				if i < j {
// 					if c1 == c2 {
// 						wetherBreak = false
// 						break
// 					}
// 				}
// 			}
// 			if !wetherBreak {
// 				break
// 			}
// 			if !wetherBreak {
// 				break
// 			}
// 		}
// 		if wetherBreak {
// 			return
// 		}
// 	} //排重

// 	// return
// }

//Longchin

// 抽4張牌
func GetTakeCards() []int {
	var Cards []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 4; i++ {
		Cards = append(Cards, rand.Intn(52))
	}
	return Cards
}

// 判斷牌組是否重複
func WetherSame(card []int) []int {
	for i := 0; i < 3; i++ {
		for j := i + 1; j < 4; j++ {
			if card[i] == card[j] {
				card = GetTakeCards()

				i = -1
				break
			}
		}
	}
	return card
}

//轉換花色並判斷大小
func Smallbig() int {

	var c Card
	var Result int
	var Cards, x []int
	var suit []string
	var xyy [4]int
	info := WinInfo{}
	takeCards := GetTakeCards()
	Cards = WetherSame(takeCards)
	for i := 0; i < 4; i++ {
		xyy[i] = Cards[i] % 4
		switch xyy[i] {
		case 0:
			c.suits = "D"
		case 1:
			c.suits = "C"
		case 2:
			c.suits = "B"
		case 3:
			c.suits = "A"
		}
		number := Cards[i]%13 + 1
		suit = append(suit, c.suits)
		x = append(x, number)

	}
	//判斷勝利並展示牌組
	// fmt.Println("電腦的牌組:  ", suit[0], x[0])
	// fmt.Println("玩家的牌組:  ", suit[1], x[1])
	// fmt.Println("玩家剩的牌組:", suit[2], x[2])
	// fmt.Println("玩家剩的牌組:", suit[3], x[3])

	for i := 0; i < 2; i++ {

		if x[i] == 1 {
			x[i] += 52
		}
	}

	if x[0] > x[1] {
		info.Winner = "computer"
		//fmt.Println("電腦獲勝")
		Result = 0
	} else if x[0] == x[1] {
		if Cards[0] > Cards[1] {
			info.Winner = "computer"
			//fmt.Println("電腦獲勝")
			Result = 0

		} else {
			info.Winner = "player"
			//fmt.Println("玩家獲勝")
			Result = 1
		}

	} else {
		info.Winner = "player"
		//fmt.Println("玩家獲勝")
		Result = 1

	}
	//fmt.Println("")
	return Result
}

// func turnsuit(card []int) {
// 	for i := 0; i < len(card); i++ {
// 		c.number = card[i] % 13

// 	}
// }

//----------------------------------------老虎機----------------------------------------------------
//產自然機率盤面
func Symbol(x int) []int {
	var Symbol []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		Symbol = append(Symbol, rand.Intn(x))
	}
	return Symbol
}

//NG滾輪回傳盤面
//0->FG
//1->SA
//2->SB
//3->SD
//4->SE
//5->S1
//6->S2
//7->S3
//8->S4
//9->WD
func NgSym(a []int) ([]int, []int, []int, []int, []int) {

	z := []int{0, 2, 5, 8, 7, 5, 8, 0, 8, 6, 7, 8, 2, 3, 7, 3, 4, 8, 4, 0, 7, 8, 5, 0, 8, 7, 8, 5, 8, 2, 7, 8, 6, 8, 6, 0, 8, 8, 7, 0, 6, 4, 8, 2, 6, 8, 7, 7, 8, 7, 8, 2, 8, 8, 4, 7, 8, 7, 8, 2, 8, 4, 8, 3, 8, 8, 7, 8, 7, 8, 3, 8, 1, 8, 7, 8, 0, 2}
	x := []int{0, 5, 6, 7, 1, 0, 2, 7, 1, 6, 8, 2, 6, 2, 3, 4, 6, 6, 9, 9, 9, 6, 5, 6, 3, 1, 3, 4, 0, 9, 7, 1, 1, 6, 0, 3, 1, 7, 0, 3, 4, 5, 6, 7, 1, 0, 7, 9, 6, 8, 0, 1, 7, 1, 8, 0, 2, 6, 5, 1, 4, 6, 0, 1, 1, 0, 1, 1, 7, 2, 6, 2, 3, 4, 6, 6, 0, 5}
	c := []int{0, 5, 6, 7, 1, 1, 0, 7, 3, 4, 2, 6, 2, 1, 0, 4, 3, 4, 0, 9, 9, 9, 4, 1, 1, 3, 4, 5, 6, 9, 8, 3, 0, 3, 5, 6, 7, 1, 1, 6, 3, 4, 7, 1, 7, 1, 5, 6, 7, 7, 8, 0, 6, 1, 6, 2, 6, 7, 8, 0, 4, 3, 4, 8, 0, 7, 3, 4, 2, 6, 7, 8, 0, 4, 3, 4, 0, 5}
	v := []int{6, 2, 5, 6, 9, 0, 6, 5, 6, 9, 6, 0, 6, 3, 0, 7, 6, 0, 2, 3, 4, 0, 9, 9, 9, 7, 4, 0, 6, 7, 8, 3, 5, 6, 7, 8, 0, 7, 7, 7, 0, 3, 7, 1, 6, 6, 8, 1, 6, 8, 0, 6, 5, 6, 9, 6, 0, 6, 3, 0, 7, 6, 6, 0, 6, 5, 6, 9, 6, 0, 6, 3, 6, 7, 6, 0, 6, 2}
	b := []int{0, 5, 6, 2, 0, 2, 8, 5, 6, 9, 5, 5, 3, 0, 0, 0, 7, 2, 3, 4, 5, 6, 7, 8, 4, 0, 6, 7, 8, 3, 3, 0, 8, 8, 4, 3, 4, 5, 6, 7, 5, 5, 3, 0, 1, 0, 5, 6, 8, 8, 2, 0, 5, 6, 9, 5, 5, 3, 0, 8, 0, 7, 2, 2, 0, 5, 6, 9, 5, 5, 3, 0, 0, 8, 7, 2, 0, 5}
	z = z[a[0] : a[0]+3]
	x = x[a[1] : a[1]+3]
	c = c[a[2] : a[2]+3]
	v = v[a[3] : a[3]+3]
	b = b[a[4] : a[4]+3]
	return z, x, c, v, b
}

//FG滾輪回傳盤面
func FgSym(a []int) ([]int, []int, []int, []int, []int) {

	z := []int{0, 7, 3, 0, 8, 8, 7, 3, 8, 4, 0, 7, 3, 4, 4, 4, 7, 7, 5, 0, 6, 5, 5, 6, 8, 7, 0, 7, 6, 7, 8, 0, 8, 8, 8, 8, 8, 8, 8, 0, 7, 7, 8, 8, 8, 8, 8, 0, 8, 8, 0, 8, 8, 7, 8, 0, 8, 8, 8, 7, 1, 8, 0, 8, 7, 8, 8, 8, 1, 7, 0, 8, 7, 8, 8, 7, 1, 8, 8, 8, 8, 8, 0, 7, 8, 2, 8, 8, 8, 8, 8, 0, 1, 8, 0, 8, 8, 8, 8, 7, 8, 8, 8, 8, 8, 7, 8, 8, 8, 8, 8, 8, 8, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 7, 8, 8, 2, 8, 8, 8, 8, 2, 7, 1, 8, 8, 8, 8, 8, 7, 0, 2, 8, 7, 8, 8, 7, 7, 0, 8, 8, 8, 8, 8, 8, 0, 2, 7, 8, 8, 8, 7, 8, 0, 8, 8, 8, 8, 8, 7, 0, 8, 8, 8, 7, 0, 7}
	x := []int{0, 7, 1, 6, 3, 4, 8, 8, 7, 8, 0, 8, 8, 4, 5, 7, 0, 8, 7, 8, 8, 4, 8, 8, 7, 0, 8, 5, 4, 0, 8, 7, 1, 8, 0, 7, 4, 5, 8, 1, 8, 1, 0, 4, 5, 8, 0, 8, 8, 7, 5, 4, 8, 1, 7, 8, 4, 5, 8, 7, 2, 0, 1, 7, 8, 0, 8, 7, 1, 8, 7, 0, 7, 5, 7, 7, 1, 7, 2, 3, 0, 5, 5, 7, 7, 2, 0, 4, 5, 7, 6, 5, 6, 7, 0, 4, 5, 6, 7, 0, 5, 3, 4, 7, 7, 7, 5, 7, 7, 0, 4, 5, 5, 7, 0, 4, 5, 6, 6, 5, 1, 6, 3, 4, 5, 6, 7, 7, 5, 2, 7, 4, 5, 6, 4, 5, 6, 7, 3, 4, 5, 6, 5, 5, 2, 3, 5, 5, 5, 7, 0, 5, 7, 3, 4, 0, 6, 5, 5, 2, 3, 4, 5, 6, 7, 3, 5, 7, 3, 4, 6, 6, 7, 4, 4, 4, 6, 7, 0, 7}
	c := []int{0, 6, 5, 8, 3, 4, 6, 6, 6, 8, 1, 8, 6, 4, 5, 0, 6, 8, 6, 8, 0, 4, 6, 0, 7, 7, 8, 0, 4, 6, 6, 1, 8, 6, 8, 8, 4, 5, 6, 2, 2, 1, 8, 4, 5, 6, 6, 0, 6, 8, 0, 4, 5, 8, 6, 0, 4, 5, 6, 6, 5, 0, 2, 1, 2, 0, 6, 2, 1, 2, 2, 0, 8, 5, 6, 6, 1, 8, 8, 3, 0, 5, 8, 7, 6, 2, 0, 4, 5, 6, 6, 8, 6, 8, 0, 4, 5, 6, 8, 0, 8, 3, 4, 8, 6, 7, 8, 6, 8, 0, 4, 5, 8, 7, 0, 4, 5, 6, 6, 8, 5, 6, 3, 4, 5, 6, 7, 0, 8, 2, 0, 4, 5, 6, 4, 0, 6, 8, 3, 4, 5, 6, 0, 8, 2, 3, 0, 5, 8, 7, 0, 8, 8, 3, 4, 0, 6, 8, 8, 2, 0, 4, 5, 6, 7, 0, 8, 8, 3, 4, 6, 6, 7, 8, 4, 0, 6, 7, 0, 6}
	v := []int{0, 5, 1, 8, 3, 4, 8, 6, 7, 8, 8, 8, 6, 4, 5, 0, 7, 8, 7, 8, 0, 4, 6, 0, 7, 7, 8, 0, 4, 7, 6, 7, 8, 6, 5, 5, 4, 5, 6, 4, 8, 8, 5, 4, 5, 6, 5, 0, 1, 5, 0, 4, 5, 1, 7, 0, 4, 5, 6, 7, 2, 0, 8, 5, 8, 0, 6, 5, 1, 8, 8, 0, 8, 5, 6, 7, 1, 5, 5, 3, 0, 5, 8, 7, 6, 2, 0, 4, 5, 6, 6, 8, 1, 5, 0, 4, 5, 6, 5, 0, 8, 3, 4, 5, 6, 7, 8, 1, 5, 0, 4, 5, 8, 7, 0, 4, 5, 6, 6, 8, 1, 6, 3, 4, 5, 6, 7, 0, 8, 2, 0, 4, 5, 6, 1, 0, 1, 5, 3, 4, 5, 6, 0, 8, 2, 3, 0, 5, 8, 7, 0, 8, 5, 3, 4, 0, 6, 8, 8, 2, 0, 4, 5, 6, 7, 0, 8, 5, 3, 4, 6, 6, 7, 8, 4, 0, 6, 7, 0, 5}
	b := []int{0, 4, 5, 8, 3, 4, 7, 6, 7, 8, 8, 8, 6, 4, 5, 0, 7, 8, 7, 8, 0, 4, 6, 0, 7, 7, 8, 0, 4, 7, 6, 7, 8, 6, 4, 2, 4, 5, 6, 8, 8, 8, 2, 4, 5, 6, 2, 0, 6, 2, 0, 4, 5, 8, 7, 0, 4, 5, 6, 7, 3, 0, 3, 2, 8, 0, 6, 2, 3, 4, 8, 0, 8, 5, 6, 7, 3, 2, 2, 3, 0, 5, 8, 7, 6, 2, 0, 4, 5, 6, 6, 8, 6, 2, 0, 4, 5, 6, 2, 0, 8, 3, 4, 2, 6, 7, 8, 6, 2, 0, 4, 5, 8, 7, 0, 4, 5, 6, 6, 8, 5, 6, 3, 4, 5, 6, 7, 0, 8, 2, 0, 4, 5, 6, 1, 0, 6, 2, 3, 4, 5, 6, 0, 8, 2, 3, 0, 5, 8, 7, 0, 8, 2, 3, 4, 0, 6, 8, 8, 2, 0, 4, 5, 6, 7, 0, 8, 2, 3, 4, 6, 6, 7, 8, 4, 0, 6, 7, 0, 4}
	z = z[a[0] : a[0]+3]
	x = x[a[1] : a[1]+3]
	c = c[a[2] : a[2]+3]
	v = v[a[3] : a[3]+3]
	b = b[a[4] : a[4]+3]
	return z, x, c, v, b
}

//判斷幾連線

var l Bingo
var y, FG int

func line(a []int, b []int, c []int, d []int, e []int) ([]int, []int, []int) {
	FG = 0
	l.five = []int{}
	l.four = []int{}
	l.three = []int{}
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)
	// fmt.Println(e)
	//判斷ＦＧ
	for j := 0; j < 3; j++ {
		if a[j] == 0 {
			for i := 0; i < 3; i++ {
				if b[i] == 0 {
					for j := 0; j < 3; j++ {
						if c[j] == 0 {
							for i = 0; i < 3; i++ {
								if d[i] == 0 {
									for j := 0; j < 3; j++ {
										switch e[j] {
										case 0:
											FG++
										}
									}

								}
							}
						}
					}
				}
			}
		}
	}
	for q := 0; q < 3; q++ {
		for w := 0; w < 3; w++ {
			if a[q] == b[w] || b[w] == 9 {
				if a[q] == 0 {
					continue
				}
				for r := 0; r < 3; r++ {
					if a[q] == c[r] || c[r] == 9 {
						for t := 0; t < 3; t++ {
							if a[q] == d[t] || d[t] == 9 {
								for y = 0; y < 3; y++ {
									if a[q] == e[y] || e[y] == 9 {
										// fmt.Println(a[q], "五連線")
										l.five = append(l.five, a[q])

									}
								}
							}
							// fmt.Println("四連線")
						}
					}
					// fmt.Println("三連線")

				}
			}
		}
	}
	// fmt.Println("五連線", l.five)
	for q := 0; q < 3; q++ {
		for w := 0; w < 3; w++ {
			if a[q] == b[w] || b[w] == 9 {
				if a[q] == 0 {
					continue
				}
				for r := 0; r < 3; r++ {
					if a[q] == c[r] || c[r] == 9 {
						for t := 0; t < 3; t++ {
							if a[q] == d[t] || d[t] == 9 {

								if a[q] == e[0] || e[0] == 9 {
									break
								} else if a[q] == e[1] || e[1] == 9 {
									break
								} else if a[q] == e[2] || e[2] == 9 {
									break
								} else {

									// fmt.Println(a[q], "四連線")
									l.four = append(l.four, a[q])
								}
							}

						}
					}

				}
			}
		}
	}
	// fmt.Println("四連線", l.four)

	for q := 0; q < 3; q++ {
		for w := 0; w < 3; w++ {
			if a[q] == b[w] || b[w] == 9 {
				if a[q] == 0 {
					continue
				}
				for r := 0; r < 3; r++ {
					if a[q] == c[r] || c[r] == 9 {
						if a[q] == d[0] || d[0] == 9 {
							break
						} else if a[q] == d[1] || d[1] == 9 {
							break
						} else if a[q] == d[2] || d[2] == 9 {
							break
						} else {
							// fmt.Println(a[q], "三連線")
							l.three = append(l.three, a[q])
						}
					}
				}
			}
		}
	}
	// fmt.Println("三連線", l.three)
	// fmt.Println("")
	return l.five, l.four, l.three
}

//賠率表計算盤面贏分
var h HitRate
var G Bingo

func Odds(five []int, four []int, three []int) int {
	// SAmap := make(map[int]int)
	// SBmap := make(map[int]int)
	// SDmap := make(map[int]int)
	// SEmap := make(map[int]int)
	// S1map := make(map[int]int)
	// S2map := make(map[int]int)
	// S3map := make(map[int]int)
	// S4map := make(map[int]int)
	// SAmap[3] = 80
	// SAmap[4] = 250
	// SAmap[5] = 500
	// SBmap[3] = 50
	// SBmap[4] = 200
	// SBmap[5] = 400
	// SDmap[3] = 40
	// SDmap[4] = 150
	// SDmap[5] = 300
	// SEmap[3] = 30
	// SEmap[4] = 100
	// SEmap[5] = 200
	// S1map[3] = 20
	// S1map[4] = 80
	// S1map[5] = 100
	// S2map[3] = 15
	// S2map[4] = 30
	// S2map[5] = 80
	// S3map[3] = 15
	// S3map[4] = 30
	// S3map[5] = 60
	// S4map[3] = 10
	// S4map[4] = 25
	// S4map[5] = 50
	var sum, sum1, sum2, sum3 int
	// var A = []int{3, 3, 4}
	G.mul = G.Bet / 100
	fiv := []int{0, 1000 * G.mul, 800 * G.mul, 600 * G.mul, 400 * G.mul, 200 * G.mul, 160 * G.mul, 120 * G.mul, 100 * G.mul}
	fou := []int{0, 500 * G.mul, 400 * G.mul, 300 * G.mul, 200 * G.mul, 160 * G.mul, 60 * G.mul, 60 * G.mul, 50 * G.mul}
	thr := []int{0, 160 * G.mul, 100 * G.mul, 80 * G.mul, 60 * G.mul, 40 * G.mul, 30 * G.mul, 30 * G.mul, 20 * G.mul}
	for _, points := range five {
		// fmt.Println(A+1, "五連線得分", fiv[points])
		sum1 += fiv[points]
	}
	for _, points := range four {
		// fmt.Println(B+1, "四連線得分", fou[points])
		sum2 += fou[points]

	}
	for _, points := range three {
		// fmt.Println(C+1, "三連線得分", thr[points])
		sum3 += thr[points]
	}
	sum = sum1 + sum2 + sum3
	// fmt.Println("盤面總得分為：", sum)
	if sum > 0 {
		h.Ng++
	}
	return sum
}

var times int = 13
var result int

func FreeGame5() int {
	j := 0
	mul := 1
	times = 13
	result = 0
	for i := 1; i < times; i++ {
		line(FgSym(Symbol(178)))
		score := Odds(l.five, l.four, l.three)
		score = score * mul * i
		result += score
		if j < 4 {
			if FG > 0 {
				times += 12
				j++
			}
		}
		if score > 0 {
			h.Fg5++
		}
		// fmt.Println("第幾局：", i)
		// fmt.Println("ＦＧ總得分", result)
	}
	h.Fgt5 += times
	return result
}

func FreeGame6() int {
	j := 0
	mul := 2
	times = 13
	result = 0
	for i := 1; i < times; i++ {
		line(FgSym(Symbol(178)))
		score := Odds(l.five, l.four, l.three)
		score = score * mul * i
		result += score
		if j < 4 {
			if FG > 0 {
				times += 12
				j++
			}
		}
		if score > 0 {
			h.Fg6++
		}
		// fmt.Println("第幾局：", i)
		// fmt.Println("ＦＧ總得分", result)
	}

	h.Fgt6 += times
	return result
}
func FreeGame7() int {
	j := 0
	mul := 3
	times = 13
	result = 0
	for i := 1; i < times; i++ {
		line(FgSym(Symbol(178)))
		score := Odds(l.five, l.four, l.three)
		score = score * mul * i
		result += score
		if j < 4 {
			if FG > 0 {
				times += 12
				j++
			}
		}
		if score > 0 {
			h.Fg7++
		}
		// fmt.Println ("第幾局：", i)
		// fmt.Println("ＦＧ總得分", result)

	}
	h.Fgt7 += times
	return result
}
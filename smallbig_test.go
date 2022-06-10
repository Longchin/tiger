package golang

import (
	"fmt"
	"testing"
)

func Test_smallbig(t *testing.T) {
	// 	RTP := 1.98
	// 	var Win float64
	// 	x := 0
	// 	var sum int
	// 	if p.Bet == 100 {

	// 	} else if p.Bet == 200 {

	// 	} else if p.Bet == 500 {

	// 	} else if p.Bet == 1000 {

	// 	} else {
	// 		return
	// 	}

	// 	for i := 0; i < 1000; i++ {
	// 		r := Smallbig()
	// 		Gamble(p.Bet, r)
	// 		sum += r
	// 		x++
	// 	}

	// 	Win = float64(sum) / float64(x)
	// 	fmt.Println("總測試次數:", x)
	// 	fmt.Println("玩家勝率:", Win)
	// 	fmt.Println("RTP:", RTP*float64(Win))

}

func TestShuffle(t *testing.T) {
	for i := 1; i < 11; i++ {
		fmt.Println(i)

		//計算測試結果
		var x, total, NGresult, FGresult, b, c, d, times, hit int
		var e, f, g float64
		x = 1000000

		G.Bet = 100
		total = x * G.Bet
		a := 0
		if G.Bet == 100 {

		} else if G.Bet == 200 {

		} else if G.Bet == 500 {

		} else if G.Bet == 1000 {

		} else if G.Bet == 2000 {

		} else if G.Bet == 3000 {

		} else if G.Bet == 5000 {

		} else if G.Bet == 10000 {

		} else {
			return
		}
		for i := 0; i < x; i++ {

			line(NgSym(Symbol(76)))
			if FG == 1 {
				// fmt.Println("--------------------------------------------------------------------")
				a = FreeGame5()
				b++
			} else if FG == 2 {
				// fmt.Println("--------------------------------------------------------------------")
				a = FreeGame6()
				c++
			} else if FG == 3 {
				// fmt.Println("--------------------------------------------------------------------")
				a = FreeGame7()
				d++
			}
			score := Odds(l.five, l.four, l.three)
			// fmt.Println(score)
			NGresult += score
			// result += a
			FGresult += a
			a = 0
			// fmt.Println(FGresult)

			// fmt.Println(NGresult)
		}
		hit = h.Fg5 + h.Fg6 + h.Fg7
		times = h.Fgt5 + h.Fgt6 + h.Fgt7
		fmt.Println("NG中獎率", float64(h.Ng)/float64(x))
		fmt.Println("FG中獎率", float64(hit)/float64(times))

		e = float64(b) / float64(x)
		f = float64(c) / float64(x)
		g = float64(d) / float64(x)

		fmt.Println("FG5hitrate", e)
		fmt.Println("FG6hitrate", f)
		fmt.Println("FG7hitrate", g)
		fmt.Println("NGRTP:", float64(NGresult)/float64(total))
		fmt.Println("FGRTP:", float64(FGresult)/float64(total))
		fmt.Println("ALLRTP:", float64(FGresult+NGresult)/float64(total))
		fmt.Println("")

	}
}

// // fmt.Println(Cryptorand())

// type human struct {
// 	address string
// 	info    []int
// }

// func TestExample(t *testing.T) {
// 	bizhen := human{
// 		address: "15B1",
// 		info:    []int{200, 100},
// 	}
// 	longchin := human{
// 		address: "2D3",
// 		info:    []int{300, 50},
// 	}
// 	winnerCity := []human{}
// 	winnerCity = append(winnerCity, bizhen, longchin)
// 	for index, human := range winnerCity {
// 		fmt.Println(index+1, human)
// 	}

//sessionId函数用来生成一个session ID，即session的唯一标识符
// func sessionId() string {
// 	b := make([]byte, 32)
// 	//ReadFull从rand.Reader精确地读取len(b)字节数据填充进b
// 	//rand.Reader是一个全局、共享的密码用强随机数生成器
// 	if _, err := io.ReadFull(rand.Reader, b); err != nil {
// 		return ""
// 	}
// 	fmt.Println(b)                              //[62 186 123 16 209 19 130 218 146 136 171 211 12 233 45 99 80 200 59 20 56 254 170 110 59 147 223 177 48 136 220 142]
// 	return base64.URLEncoding.EncodeToString(b) //将生成的随机数b编码后返回字符串,该值则作为session ID
// }

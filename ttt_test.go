package golang

import (
	cryptoRand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/big"
	mathRand "math/rand"
	"testing"
)

// func Test_main(t *testing.T) {
// 	main()
// }

func BenchmarkCryptorand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Mathrand()
		Cryptorand()
	}

}
func Cryptorand() (x *big.Int) {

	for i := 0; i < 4; i++ {
		x, err := cryptoRand.Int(cryptoRand.Reader, big.NewInt(52))
		if err != nil {
			return x
		}
		// fmt.Println(crypto.Int(crypto.Reader, big.NewInt(52)))
		return x
	}
	return
}
func Mathrand() {
	// mathRand.Seed(time.Now().UnixNano())
	for i := 0; i < 4; i++ {
		mathRand.Intn(100)
	}

}
func RandNumber(num int) int {
	bytes := make([]byte, 4)
	_, err := cryptoRand.Read(bytes)
	if err != nil {
		return 0
	}
	return int(binary.BigEndian.Uint32(bytes)) % num
}

func GetRangeRand(min, max int) int {

	// max 要+1才有辦法取到最後一個數字
	max = max + 1

	// max-min 小於0會panic
	if max-min < 0 {
		return min
	}

	bytes := make([]byte, 4)
	_, err := cryptoRand.Read(bytes)
	if err != nil {
		return 0
	}
	seed := int(binary.BigEndian.Uint32(bytes))
	mathRand.Seed(int64(seed))

	return min + mathRand.Intn(max-min)
}

func TestExample1(t *testing.T) {
	// card := []int{}
	// for i := 0; i < 3; i++ {
	// 	randNum := RandNumber(52) + 1
	// 	card = append(card, randNum)
	// }
	// fmt.Println(card)
	for i := 0; i < 1000000; i++ {
		num := GetRangeRand(3, 5)
		if num < 3 || num > 5 {
			fmt.Println("error")
		}
	}
}

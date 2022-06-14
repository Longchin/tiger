package golang

import (
	"fmt"
	"testing"
)

type Line struct {
	symbol string
	place  []int
	hit    int
	odds   int
}

func TestGetTakeCards(t *testing.T) {
	// fmt.Println(tiger())
	// // Wline()
	lines := []Line{}
	l := Line{
		symbol: "SA",
		place:  []int{1, 5, 9},
		hit:    3,
		odds:   10,
	}
	lines = append(lines, l, l)
	fmt.Println(lines[0].place)

}

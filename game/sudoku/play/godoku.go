package main

import (
	"math"
	"github.com/gobasic/game/sudoku"
	"fmt"
)

func getSize(bstr string) int {
	return int(math.Sqrt(float64(len(bstr))))
}

func main() {
	text := "1...79.8." + ".5.413.6." + ".6......." +
			"..1......" + "8.5......" + "....3...7" +
			".94.5...." + "7..8.63.9" + ".....74.."
	// text := "......54." + ".278...31" + "........2" +
	//   ".7862...." + "........6" + ".....37.." +
	//   ".432.7..8" + "1..45...3" + "..5.8...."

	// Create the sudoku board
	sz := getSize(text)
	board := sudoku.NewBoard(sz)

	for i, ch := range text {
		r := i / sz + 1 // rows
		c := i % sz + 1 // columns
		fmt.Println(r)
		fmt.Println(c)
		// ch : character in string literal in text range
		if ch >= '0' && ch <= '9' {
			ch.(string)
			//  board.Set(r,c, ch - '0')
//			number, _ := strconv.Atoi(ch)
			//board.Set(r, c, ch - '0')
		}
	}
	board.Print()
	if !board.Solved() {
		board.PrintAll()
	}
}

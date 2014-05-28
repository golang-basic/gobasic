package tic_tac_toe

/**
 only defined for a  n X n matrix board

having a bucket with linked list of sets for winning streaks
 winning streaks should be <= diagonal elements

 */

var winarr []*[][]int

func (t *gameTicTacToe) WinSliceInit() {
	for i := range t.Cells {
		winarr[i] =  t.getSlices(i)
	}
}


func (t *gameTicTacToe) getSlices(cell int) (winStreak *[][]int) {
	winarr := [][]int{[]int {1,2,3},[]int {1,2,3}}

	return &winarr
}

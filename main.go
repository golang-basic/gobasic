package main

import (
	"fmt"
	"github.com/golang-basic/tic_tac_toe"
	"github.com/golang-basic/guess_number"
	"github.com/golang-basic/tower_of_hanoi"
//	"github.com/golang-basic/game/tic_tac_toe"
//	"github.com/golang-basic/game/guess_number"
//	"github.com/golang-basic/game/tower_of_hanoi"
	//"github.com/gobasic/game/get_blog_stats"
)

func main() {
	choice := 4
	fmt.Println("Enter the choice.\n 1. Tic-r®Tac-Toe \n 2. Guess Number\n 3. Tower of Hanoi ")
	fmt.Scanf("%d", &choice)
	switch {
	case choice == 1:
		fmt.Println("Playing Tic-Tac-Toe")
		tictactoe := tic_tac_toe.NewTicTacToe()
		tictactoe.Init()
		tictactoe.Play()
	case choice == 2:
		fmt.Println("Playing Guess Game")
		guess_number.TestGuessGame()
		gg := guess_number.GuessGame {1, "", 0}
		gg.Greet()
		gg.Choose_Number()
		gg.Play()
	case choice == 3:
		fmt.Println("Playing Tower of hanoi")
		tower_of_hanoi.NewTowerOfHanoi(3)
		tower_of_hanoi.SimulateGame()
	case choice == 4:
		fmt.Println("Playing Blog runs")
	//	get_blog_stats.ExecuteBlog()
	}
}

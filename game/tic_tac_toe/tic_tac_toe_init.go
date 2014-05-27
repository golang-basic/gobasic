package tic_tac_toe

import (
	"strconv"
	"math/rand"
	"fmt"
)

func (tictactoe *gameTicTacToe) Init() {
	fmt.Println("Welcome to TicTacToe!!!")
	tictactoe.showLayout()
	tictactoe.askUserSymbol()
	turn := tictactoe.randomSelectTurn()
	fmt.Println("Game ON!! User ",turn, " will play now.")
}

func (tictactoe *gameTicTacToe) showLayout() {
	fmt.Println("Game board and all Postions")
	tictactoe.printCellPositions()
	fmt.Println("Your Game Board")
	tictactoe.layout()
}

func (tictactoe *gameTicTacToe) layout() {
	blockLine := " __ __ __ __ __ __"
	endLine := " |"
	line := []string{endLine}
	fmt.Println(blockLine)
	for i := 0; i <= 8; i++ {
		line = append(line, tictactoe.cells[i], endLine, tictactoe.cells[i+1], endLine, tictactoe.cells[i+2], endLine)
		fmt.Println(line)
		i++
		i++
		fmt.Println(blockLine)
		line = []string{endLine}
	}
	tictactoe.SeenLayout = true
}

func (tictactoe *gameTicTacToe) printCellPositions() {
	blockLine := " __ __ __ __ __"
	endLine := "|"
	line := []string{endLine}
	fmt.Println(blockLine)
	for i := 0; i <= 8; i++ {
		line = append(line, strconv.Itoa(i), endLine, strconv.Itoa(i+1), endLine, strconv.Itoa(i+2), endLine)
		fmt.Println(line)
		i++
		i++
		fmt.Println(blockLine)
		line = []string{endLine}
	}
	tictactoe.SeenPosition = true
}

func (tictactoe *gameTicTacToe) askUserSymbol() {
	if (tictactoe.checkUserSymbol()) {
		tmpsymbol := "tmp"
	Loop:
		for i := 0; i < tictactoe.chances; i++ {
			fmt.Println("Please select a symbol (O or X) :")
			fmt.Scanf("%s", &tmpsymbol)
			if !(tmpsymbol == "X" || tmpsymbol == "O") {
				continue
			}
			tictactoe.UserSymbol = tmpsymbol
			switch {
			case tictactoe.UserSymbol == "X" :
				tictactoe.AISymbol = "O"
				break Loop

			case tictactoe.UserSymbol == "O" :
				tictactoe.AISymbol = "X"
				break Loop
			}
		}
		if !(tictactoe.UserSymbol == "X" || tictactoe.UserSymbol == "O") {
			fmt.Println("Failed to get a User Symbol. Choosing symbol O for User.")
			tictactoe.AISymbol = "X"
			tictactoe.UserSymbol = "O"
		}
	}
}

func (tictactoe *gameTicTacToe) checkUserSymbol() bool {
	if (tictactoe.UserSymbol == "X" || tictactoe.UserSymbol == "O") {
		return false
	}
	return true
}

func (tictactoe *gameTicTacToe) randomSelectTurn() (turn string) {
	if rand.Intn(101)%2 == 0 {
		tictactoe.Turn = Players[rand.Intn(101)%2]
	} else {
		tictactoe.Turn = Players[rand.Intn(101)%2]
	}
	return tictactoe.Turn
}

package tic_tac_toe

import (
	"fmt"
	"strconv"
	"os"
)

type gameTicTacToe struct {
	Turn       string
	UserSymbol string
	AISymbol   string
	cells      []string
	cellbytes  []byte
	GameOn	   bool
	SeenPosition bool
	SeenLayout bool
}

var Players = [2]string{"Player", "AI"}

func NewTicTacToe() (tictactoe *gameTicTacToe) {
	tictactoe = new(gameTicTacToe)
	tictactoe.cells = []string {" ", " ", " ", " ", " ", " ", " ", " ", " "}
	tictactoe.cellbytes = []byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '}
	tictactoe.GameOn = true
	return
}

func (tictactoe *gameTicTacToe) Init() {
	fmt.Println("Welcome to TicTacToe!!!")
	tictactoe.showLayout()
	tictactoe.askUserSymbol()
//	tictactoe.randomSelectTurn()
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

func (tictactoe *gameTicTacToe) Play(){
	for(tictactoe.GameOn){
		tictactoe.GameOn = false
	}
	fmt.Println("Game end !!")
	os.Exit(3)
}

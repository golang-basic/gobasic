package tic_tac_toe

import "fmt"

type gameTicTacToe struct {
	Turn       string
	UserSymbol string
	AISymbol   string
	cells      []string
	cellbytes  []byte
	GameOn	   bool
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
	tictactoe.randomSelectTurn()
}

func (tictactoe *gameTicTacToe) showLayout() {
	fmt.Println("Cells and Postions")
	tictactoe.printCellPositions()
	tictactoe.layout()
}

func (tictactoe *gameTicTacToe) layout() {
	blockLine := " __ __ __ "
	endLine := "|"
	fmt.Println(blockLine)
	for i := 0; i <= 8; i++ {
		line := append(endLine, tictactoe.cells[i], endLine, tictactoe.cells[i+1], endLine, tictactoe.cells[i+2], endLine)
		fmt.Println(line)
		i++
		i++
		fmt.Println(blockLine)
	}
}

func (tictactoe *gameTicTacToe) printCellPositions() {
	blockLine := " __ __ __ "
	endLine := "|"
	fmt.Println(blockLine)
	for i := 0; i <= 8; i++ {
		line := append(endLine, string(i), endLine, string(i+1), endLine, string(i+2), endLine)
		fmt.Println(line)
		i++
		i++
		fmt.Println(blockLine)
	}
}

func (tictactoe *gameTicTacToe) Play(){
	for(tictactoe.GameOn){

	}
}

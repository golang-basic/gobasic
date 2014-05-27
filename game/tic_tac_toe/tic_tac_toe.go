package tic_tac_toe

import (
	"fmt"
	"os"
)

type gameTicTacToe struct {
	Turn         string
	UserSymbol   string
	AISymbol     string
	cells        []string
	cellbytes    []byte
	GameOn       bool
	SeenPosition bool
	SeenLayout   bool
	chances      int
}

var Players = [2]string{"Human", "AI"}

func NewTicTacToe() (tictactoe *gameTicTacToe) {
	tictactoe = new(gameTicTacToe)
	tictactoe.cells = []string {" ", " ", " ", " ", " ", " ", " ", " ", " "}
	tictactoe.cellbytes = []byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '}
	tictactoe.GameOn = true
	tictactoe.chances = 2
	return
}


func (tictactoe *gameTicTacToe) Play() {
	Loop:
	for (tictactoe.GameOn) {
//		tictactoe.GameOn = false
		switch {
			//human plays
		case tictactoe.Turn == Players[0]:

			// AI plays
		case tictactoe.Turn == Players[1]:

		}
		break Loop //TODO
	}
	fmt.Println("Game end !!")
	os.Exit(3)
}

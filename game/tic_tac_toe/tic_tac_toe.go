package tic_tac_toe

import (
	"fmt"
	"os"
	"math"
)

type gameTicTacToe struct {
	Turn         string
	UserSymbol   string
	AISymbol     string
	Cells        []string
	cellbytes    []byte
	GameOn       bool
	SeenPosition bool
	SeenLayout   bool
	chances      int
}

var Players = [2]string{"Human", "AI"}
var size int
var lengthOfSide int
var corners []int

func NewTicTacToe() (tictactoe *gameTicTacToe) {
	tictactoe = new(gameTicTacToe)
	tictactoe.Cells = []string {" ", " ", " ", " ", " ", " ", " ", " ", " "}
	tictactoe.cellbytes = []byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '}
	tictactoe.GameOn = true
	tictactoe.chances = 2
	tictactoe.getCorners()
	return
}

func (tictactoe *gameTicTacToe) getCorners() {
	size = len(tictactoe.Cells)
	lengthOfSide = int(math.Sqrt(float64(size)) - 1)
	corners = []int {0, 0+lengthOfSide, size-1-lengthOfSide, size-1}
}

func (tictactoe *gameTicTacToe) Play() {
	//	Loop:
	for (tictactoe.GameOn) {
		//		tictactoe.GameOn = false
		switch {
			//human plays
		case tictactoe.Turn == Players[0]:
			newtictactoe := tictactoe.PlayersMove()
			if(newtictactoe != nil){
				fmt.Println("Getting NEW TIC TAC TOAAA")
				tictactoe.Cells = newtictactoe.Cells
			}
			// AI plays
		case tictactoe.Turn == Players[1]:
			tictactoe.AIMove()
		}
		//		break Loop //TODO
	}
	fmt.Println("Game end !!")
	//	os.Exit(3)
}
func (tictactoe *gameTicTacToe) AIMove() (newTicTacToe *gameTicTacToe){
	move := tictactoe.NextAIMove()
	tictactoe.Cells[move] = tictactoe.AISymbol

	//check if User wins after making move
	if tictactoe.IsWinner(tictactoe.AISymbol) {
		tictactoe.showLayout()
		fmt.Println("Yooo!!! AI won !!")
		newTicTacToe = tictactoe.AskPlayerToPlayOneMoreTime()
	} else {
		// if player does not win, Select the next computer move
		// set the next player to AI
		if (tictactoe.isCellsFull()) {
			tictactoe.tieGame()
		}else {
			tictactoe.Turn = Players[0]
		}
	}
	return
}
func (tictactoe *gameTicTacToe) NextAIMove() (move int) {
	tictactoe.Turn = Players[0]

	move = tictactoe.AIChoosesMove(tictactoe.AISymbol)
	if move != 99 {
		return
	}

	move = tictactoe.AIChoosesMove(tictactoe.UserSymbol)
	if move != 99 {
		return
	}

//	aiPlaces := tictactoe.GetAllAIPlaces()
	available, move := tictactoe.AISelectsCenterPosition()
	if available {
		return
	}
	available, move = tictactoe.AISelectsCornerPosition()
	if available {
		return
	}

	move = tictactoe.AIChoosesMiddleElement([]int{ 1, 3, 5, 7})

	return
}

func (tictactoe *gameTicTacToe) AIChoosesMiddleElement(arr []int) int {
	for _, index := range arr {
		if (tictactoe.isFreeSpace(index)) {
			return index
		}
	}
	return 99
}

func (tictactoe *gameTicTacToe) AIChoosesMove(symbol string) (int) {
	for i := range tictactoe.Cells {
		copy := tictactoe
		if copy.isFreeSpace(i) {
			copy.Cells[i] = symbol
			if copy.IsWinner(symbol) {
				return i
			} else {
				copy.Cells[i] = " "
			}
		}
	}
	return 99
}

func (tictactoe *gameTicTacToe) GetAllAIPlaces() (aiPlaces []int) {
	for i := range tictactoe.Cells {
		if (tictactoe.Cells[i] == tictactoe.AISymbol) {
			aiPlaces = append(aiPlaces, i)
		}
	}
	return
}

/**
	select the center empty space
*/
func (tictactoe *gameTicTacToe) AISelectsCenterPosition() (midavailable bool, midposition int) {
	midavailable = true
	midposition = len(tictactoe.Cells)/2
	if (tictactoe.isFreeSpace(int(midposition))) {
		tictactoe.Cells[midposition] = tictactoe.AISymbol
		return midavailable, midposition
	} else {
		midavailable = false
		midposition = 0
	}
	return
}

/**
	if center is not available select one of the free corners
 */
func (tictactoe *gameTicTacToe) AISelectsCornerPosition() (corneravailable bool, cornerPosition int) {
	corneravailable = true
	// corners will always be 4 in size it its 2X2 square
	tictactoe.layout()

	cornerPosition = corners[0]
	for _, cornerIndex := range corners {
		if (tictactoe.isFreeSpace(cornerIndex)) {
			cornerPosition = cornerIndex
			return
		}
	}

	return false, 0
}



func (tictactoe *gameTicTacToe) PlayersMove() (newTicTacToe *gameTicTacToe) {
	tictactoe.showLayout()
	tictactoe.AskUserMove()
	//check if User wins after making move
	if tictactoe.IsWinner(tictactoe.UserSymbol) {
		tictactoe.showLayout()
		tictactoe.congratulatePlayer()
		newTicTacToe = tictactoe.AskPlayerToPlayOneMoreTime()
	} else {
		// if player does not win, Select the next computer move
		// set the next player to AI
		if (tictactoe.isCellsFull()) {
			tictactoe.tieGame()
		}else {
			tictactoe.Turn = Players[1]
		}
	}
	return
}

func (tictactoe *gameTicTacToe) AskUserMove() (move int) {
	tmpMove := 99
Loop:
	for i := 0; i < tictactoe.chances ; i++ {
		fmt.Println("Player ", tictactoe.Turn, ", select your move (0/8) :")
		fmt.Scanf("%d", &tmpMove)
		if (tmpMove >= 0 && tmpMove <= 8) {
			if (tictactoe.isFreeSpace(tmpMove)) {
				tictactoe.Cells[tmpMove] = tictactoe.UserSymbol
				fmt.Println("Good Move!")
				break Loop
			}
		}
		fmt.Println("Try Again !!")
		tictactoe.layout()
		tmpMove = 99
	}
	if (tmpMove == 99) {
		tmpMove = tictactoe.randomlySelectFirstFreeSpace()
		if (tmpMove == -1) {
			fmt.Println("No Free moves available. Exiting Game.")
			os.Exit(1)
		}
		fmt.Printf("\nChances over. Selecting your first available choice at %d\n", tmpMove)
		tictactoe.Cells[tmpMove] = tictactoe.UserSymbol
	}
	return tmpMove
}

func (tictactoe *gameTicTacToe) isFreeSpace(position int) bool {
	return tictactoe.Cells[position] == " "
}

func (tictactoe *gameTicTacToe) isCellsFull() bool {
	for _, cell := range tictactoe.Cells {
		if cell == " " {
			return false
		}
	}
	return true
}

func (tictactoe *gameTicTacToe) IsWinner(symbol string) bool {
	return ((tictactoe.Cells[0] == symbol && tictactoe.Cells[3] == symbol && tictactoe.Cells[6] == symbol) || //first vertical line
			(tictactoe.Cells[1] == symbol && tictactoe.Cells[4] == symbol && tictactoe.Cells[7] == symbol) || //second vertical line
			(tictactoe.Cells[2] == symbol && tictactoe.Cells[5] == symbol && tictactoe.Cells[8] == symbol) || //third vertical line
			(tictactoe.Cells[0] == symbol && tictactoe.Cells[1] == symbol && tictactoe.Cells[2] == symbol) || //first Horizontal line
			(tictactoe.Cells[3] == symbol && tictactoe.Cells[4] == symbol && tictactoe.Cells[5] == symbol) || //second Horizontal line
			(tictactoe.Cells[6] == symbol && tictactoe.Cells[7] == symbol && tictactoe.Cells[8] == symbol) || //third Horizontal line
			(tictactoe.Cells[2] == symbol && tictactoe.Cells[4] == symbol && tictactoe.Cells[6] == symbol) || //first diagonal line
			(tictactoe.Cells[0] == symbol && tictactoe.Cells[4] == symbol && tictactoe.Cells[8] == symbol))  //second diagonal line
}

func (tictactoe *gameTicTacToe) randomlySelectFirstFreeSpace() (int) {
	for i, _ := range tictactoe.Cells {
		if (tictactoe.isFreeSpace(i)) {
			return i
		}
	}
	return -1
}

func (tictactoe *gameTicTacToe) congratulatePlayer() {
	fmt.Println("Congratulations!! You won the game!!!")
}

func (tictactoe *gameTicTacToe) AskPlayerToPlayOneMoreTime() (newTicTacToe *gameTicTacToe) {
	oneMore := "n"
	fmt.Println("Play Game Again (y/n):")
	fmt.Scanf("%s", &oneMore)
	if (oneMore == "y" || oneMore == "Y") {
		fmt.Println("Cool!! lets play again!!")
		newTicTacToe = NewTicTacToe()
		newTicTacToe.Init()
		tictactoe.GetNew(newTicTacToe)
		return
	} else {
		tictactoe.GameOn = false
	}
	return
}

func (tictactoe *gameTicTacToe) GetNew(newtictactoe *gameTicTacToe) {
	tictactoe.Cells = newtictactoe.Cells
	tictactoe.AISymbol = newtictactoe.AISymbol
	tictactoe.UserSymbol = newtictactoe.UserSymbol
	tictactoe.Turn = newtictactoe.Turn
}

func (tictactoe *gameTicTacToe) tieGame() {
	fmt.Println("Game is a tie !!!")
	tictactoe.AskPlayerToPlayOneMoreTime()
}



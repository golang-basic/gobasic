package tic_tac_toe_test

import (
	. "github.com/gobasic/game/tic_tac_toe"
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TicTacToe", func() {
	var tictactoe = NewTicTacToe()

	BeforeEach(func() {
		tictactoe.Init()
	})
	Describe("Initialization", func() {
		Context("setup game for users to start playing", func() {

			It("sets the game-on to true", func() {
				Expect(tictactoe.GameOn).Should(BeTrue())
			})
			It("sets the layout on screen", func() {
				Expect(tictactoe.GameOn).Should(BeTrue())
				Expect(tictactoe.SeenPosition).Should(BeTrue())
				Expect(tictactoe.SeenLayout).Should(BeTrue())
			})
			It("lets user pick up a symbol and chooses symbol for AI", func(){
				Expect(tictactoe.GameOn).Should(BeTrue())
				fmt.Println(tictactoe.UserSymbol)
				if !(tictactoe.UserSymbol != "X" || tictactoe.UserSymbol != "O") {
					Fail("Failed to get user Symbol")
				}
				if !(tictactoe.UserSymbol != "X" || tictactoe.UserSymbol != "O") {
					Fail("Failed to set AI Symbol")
				}
			})
			It("randomly selects who should play first turn", func() {
				Expect(tictactoe.GameOn).Should(BeTrue())
				if! (tictactoe.Turn != Players[0] || tictactoe.Turn != Players[1]){
					Fail("Failed to select the turn")
				}
			})
		})
	})
	Describe("Play", func() {
		Context("Plays game by choosing moves based on other players moves", func() {

		})
	})
})

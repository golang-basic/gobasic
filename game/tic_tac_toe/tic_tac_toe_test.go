package tic_tac_toe_test

import (
	. "github.com/gobasic/game/tic_tac_toe"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TicTacToe", func() {
	Describe("Initialization", func() {
		Context("setsup game for users to start playing", func() {
			var tictactoe = NewTicTacToe()

			BeforeEach(func() {
				tictactoe.Init()
			})
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
				if tictactoe.UserSymbol != "X" || tictactoe.UserSymbol != "O" {
					Fail("Failed to get user Symbol")
				}
				if tictactoe.UserSymbol != "X" || tictactoe.UserSymbol != "O" {
					Fail("Failed to set AI Symbol")
				}

			})
			It("randomly selects who should play first turn", func() {
				Expect(tictactoe.GameOn).Should(BeTrue())

			})
		})
	})
})

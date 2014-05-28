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
			It("lets user pick up a symbol and chooses symbol for AI", func() {
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
				if !(tictactoe.Turn != Players[0] || tictactoe.Turn != Players[1]) {
					Fail("Failed to select the turn")
				}
			})
		})
	})
	Describe("Play", func() {
		Context("player goes first && plays choosing moves", func() {
			BeforeEach(func() {
				tictactoe.Turn = Players[0]
			})
			It("player gets to make a move && selects first free space for user if no input given", func() {
				move := tictactoe.AskUserMove()
				Ω(tictactoe.Cells[move]).To(Equal(tictactoe.UserSymbol))
			})

			It("checks if player wins after selecting 3 moves", func() {
				move1 := tictactoe.AskUserMove()
				move2 := tictactoe.AskUserMove()
				move3 := tictactoe.AskUserMove()

				Ω(tictactoe.Cells[move1]).To(Equal(tictactoe.UserSymbol))
				Ω(tictactoe.Cells[move2]).To(Equal(tictactoe.UserSymbol))
				Ω(tictactoe.Cells[move3]).To(Equal(tictactoe.UserSymbol))
				Ω(tictactoe.IsWinner(tictactoe.UserSymbol)).Should(BeTrue())
			})
			It("sees layout, selects move, on winning sets gameon to false", func() {
				tictactoe = NewTicTacToe()
				tictactoe.Init()
				tictactoe.Cells[1] = tictactoe.UserSymbol
				tictactoe.Cells[2] = tictactoe.UserSymbol
				/**
				Next autoselected move will be at position 0 and user will win
				 */
				tictactoe.PlayersMove()
				Ω(tictactoe.GameOn).ToNot(BeTrue())
			})

			It("sees layout, selects move, does not win and sets the next turn for player 2", func() {
				tictactoe = NewTicTacToe()
				tictactoe.Init()
				tictactoe.Cells[2] = tictactoe.UserSymbol

				tictactoe.PlayersMove()
				Ω(tictactoe.Turn).Should(Equal(Players[1]))
			})
		})
		Context("AI goes first ", func() {
			It("gets all the positions occupied by AI symbol", func() {
				tictactoe.Cells[2] = tictactoe.AISymbol
				tictactoe.Cells[4] = tictactoe.AISymbol
				tictactoe.Cells[7] = tictactoe.AISymbol
				tictactoe.Cells[3] = tictactoe.AISymbol
				tictactoe.Cells[0] = tictactoe.AISymbol

				aipos := tictactoe.GetAllAIPlaces()
				Ω(aipos).To(HaveLen(5))
				Ω(aipos).Should(ContainElement(2))
				Ω(aipos).Should(ContainElement(4))
				Ω(aipos).Should(ContainElement(7))
				Ω(aipos).Should(ContainElement(3))
				Ω(aipos).Should(ContainElement(0))
			})
			It("selects center position if there are no AI positions occupied ", func() {
				tictactoe = NewTicTacToe()
				tictactoe.Init()
				aipos := tictactoe.GetAllAIPlaces()
				Ω(aipos).To(HaveLen(0))
				move := tictactoe.NextAIMove()
				Ω(move).To(Equal(4))
			})
			It("selects corner positions if there are no AI positions occupied & center position is occupied ", func() {
				tictactoe = NewTicTacToe()
				tictactoe.Init()
				tictactoe.Cells[4] = tictactoe.UserSymbol
				tictactoe.Cells[0] = tictactoe.UserSymbol
				tictactoe.Cells[8] = tictactoe.AISymbol

				aipos := tictactoe.GetAllAIPlaces()
				Ω(aipos).To(HaveLen(1))
				move := tictactoe.NextAIMove()
				Ω(move).To(Equal(2))
			})
			It("AI selects the move by which it can win the game", func() {
				tictactoe = NewTicTacToe()
				tictactoe.Init()
				tictactoe.Cells[4] = tictactoe.AISymbol
				tictactoe.Cells[0] = tictactoe.AISymbol

				move := tictactoe.NextAIMove()
				Ω(move).To(Equal(8))

			})
			It("AI selects the move by which User can win the next game", func() {
				tictactoe = NewTicTacToe()
				tictactoe.Init()
				tictactoe.Cells[3] = tictactoe.UserSymbol
				tictactoe.Cells[0] = tictactoe.UserSymbol

				move := tictactoe.NextAIMove()
				Ω(move).To(Equal(6))

			})
//			It("")
		})
	})
})

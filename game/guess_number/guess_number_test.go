package guess_number_test

import (
	. "github.com/gobasic/game/guess_number"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GuessNumber", func() {
	var guessGame GuessGame

	BeforeEach(func() {
		guessGame = GuessGame{
			Guesses_made: 1,
			Name: "",
			Number: 0,
		}
	})

	Describe("Names the User", func() {
		Context("greets and asks the user for Name", func() {
			It("Name should be Guesser if user does not input Name", func() {
				Ω(guessGame.Name).Should(Equal(""))
				guessGame.Greet()
				Ω(guessGame.Name).ShouldNot(BeNil())
				Ω(guessGame.Name).Should(Equal(GuessName))
			})
		})
	})
	Describe("Random Number", func() {
		Context("Choosen random Number", func() {
			It("should be less than 20", func() {
				Ω(guessGame.Number).Should(Equal(int32(0)))
				guessGame.Choose_Number()
				Ω(guessGame.Number).Should(BeNumerically("<", 20))
				Ω(guessGame.Number).Should(BeNumerically(">=", 1))
			})
			It("should not be reinitialized to another Number during the game", func() {
				guessGame.Choose_Number()
				num_once := guessGame.Number
				for i := 0; i < 10; i++ {
					guessGame.Choose_Number()
					Ω(guessGame.Number).Should(Equal(num_once))
				}
			})
		})
	})
	Describe("Play Game", func() {
		JustBeforeEach(func() {
			guessGame.Choose_Number()
		})
		Context("Guessing", func() {
			It("a number to see if its greater or less than choosen number", func() {
				guessGame.Play()
			})
			It("makes guesses for a maximum of 6 times", func() {

			})
		})
	})
})

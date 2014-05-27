package guess_number

import (
	"fmt"
	"math/rand"
)

type GuessGame struct {
	Guesses_made int
	Name         string
	Number       int32
}

const GuessName = "Guesser"
const GuessingChances = 6
const MaxNumber = 19

func TestGuessGame() {
	gg := GuessGame {1, "", 0}
	gg.Greet()
	gg.Choose_Number()
	gg.Play()
}

func (gg *GuessGame) Choose_Number() {
	if (gg.Number == 0) {
		gg.Number = rand.Int31n(MaxNumber)+1
	}
}
func (gg *GuessGame) Greet() {
	fmt.Println("Hello! What is your Name?\n ")
	fmt.Scanf("%s", &gg.Name)
	if (gg.Name == "") {
		gg.Name = GuessName
	}
	fmt.Printf("Well, %s, I am thinking of a number between 1 and 20.\n", gg.Name)
}

func (gg *GuessGame) makeGuess() (bool) {
	var guess int32

	gg.askUser(&guess)

	if guess < gg.Number {
		fmt.Printf("Your %d is too low.\n", guess)
	}
	if guess > gg.Number {
		fmt.Printf("Your %d is too high.\n", guess)
	}
	if guess == gg.Number {
		fmt.Printf("Good job, %s! You guessed my number in %d guesses!", gg.Name, gg.Guesses_made)
		return false
	}

	return true
}

func (gg *GuessGame) askUser(guess *int32) {
	fmt.Println("\nTake a guess !!")
	fmt.Scanf("%d", *guess)
}

func (gg *GuessGame) Play() {
	for gg.Guesses_made <= GuessingChances && gg.makeGuess() {
		gg.Guesses_made++
	}
	if (gg.Guesses_made == GuessingChances) {
		fmt.Printf("Nope. The number I had in mind was %d\n", gg.Number)
	} else {
		fmt.Println("\nGood Job!!")
	}
}

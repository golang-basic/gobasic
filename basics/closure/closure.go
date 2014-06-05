package main

import (
	"fmt"
)

func main() {
	/**
	Closure revursive fibonacci Sequence
	 */
	fmt.Println(GenerateFibonacciSeries(0))
	fmt.Println(GenerateFibonacciSeries(1))
	fmt.Println(GenerateFibonacciSeries(2))
	fmt.Println(GenerateFibonacciSeries(5))

	/**
	Get basic closure
	 */

	addSeries := GetAnnonymousAddSeries()
	squareSeries := GetAnnonymousSquareSeries()
	tylorSeries := GetAnnonymousTylorSeries()
	fmt.Println("Number series")
	for i := 0; i < 10; i++ {
		fmt.Printf(" %d", addSeries())
	}
	fmt.Println("Square series")

	for i := 0; i < 10; i++ {
		fmt.Printf(" %d", squareSeries())
	}
	fmt.Println("Tylor series")

	for i := .999; i > 0.5; i = i-.2 {
		fmt.Printf(" %f", tylorSeries(i))
	}

	/**
	Closure with goroutine
	 */
	fmt.Println("\nClosures and GoRoutines")
	WrongClosureWithRoutine()
	CorrectClosureWithRoutine()
}

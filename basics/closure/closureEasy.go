package main

import "fmt"

func GetAnnonymousAddSeries() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func GetAnnonymousSquareSeries() func() int {
	x := 1
	return func() int {
		i := x*x
		x++
		return i
	}
}

func GetAnnonymousTylorSeries() func(x float64) float64 {

	// x  belongs to open interval 1,-1
	return func(x float64) float64 {
		if (x > 1 || x < -1) {
			fmt.Println("Not a Tylor series number. Pass number between a and -1")
			return -9999999.99999
		} else {
			x = 1/(1-x)
			return x
		}
	}
}

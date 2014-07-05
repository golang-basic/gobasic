package main

import "fmt"

// A Month specifies a month of the year (January = 1, ...).
type Month int

const (
	January Month = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

var months = [...]string{
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
}

// String returns the English name of the month ("January", "February", ...).
func (m Month) String() string { return months[m-1] }

func mainForMonth() {
	month := December
	if month == December {
		fmt.Println("Found a December")
	}
	// %!v(PANIC=runtime error: index out of range)
	month = month + Month(2)
	//fmt.Println(month)

	month = January + Month(2)
	fmt.Println(month)

	month++
	fmt.Println(month)

	day := 34
	month = Month(day % 31)
	fmt.Println(month)

	val := int(month) + 4
	fmt.Println(val)

	month = Month(val) + 1
	fmt.Println(month)
}

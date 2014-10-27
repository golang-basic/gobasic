package main

import "fmt"

// for enums define a type of the enum that one want
// we want days be of type int where int ranges from 1 to 7
// Another reason is we want to be able to perform operations on Day
// adding 1/2..7 or subtracting 1/2..7 from the variables of type Day
type Day int

// next we define the constants of type Day
// we use iota to provide int values to each of the contant types
const (
	MONDAY Day = 1 + iota
	TUESDAY
	WEDNESDAY
	THURSDAY
	FRIDAY
	SATURDAY
	SUNDAY
)

// Enums should be able to printout as strings
// so we declare the next best thing, a slice of strings
// for eg. the string value will be used in the println
var days = [...]string {
	"MONDAY",
	"TUESDAY",
	"WEDNESDAY",
	"THURSDAY",
	"FRIDAY",
	"SATURDAY",
	"SUNDAY",
}

// String() function will return the english name
// that we want out constant Day be recognized as
func (day Day) String() string {
	return days[day - 1]
}

func mainDayTest() {
	firstDay := MONDAY
	tellItLikeItIs(firstDay)

	thirdDay := WEDNESDAY
	tellItLikeItIs(thirdDay)

	fifthDay := FRIDAY
	tellItLikeItIs(fifthDay)

	sixthDay := SATURDAY
	tellItLikeItIs(sixthDay)

	seventhDay := SUNDAY
	tellItLikeItIs(seventhDay)

}

func tellItLikeItIs(day Day) {
	switch (day){
	case MONDAY:
		fmt.Println("Mondays are bad.")

	case FRIDAY:
		fmt.Println("Fridays are better.")

	case SATURDAY, SUNDAY:
		fmt.Println("Weekends are best.")

	case TUESDAY, WEDNESDAY, THURSDAY :
		fmt.Println("Midweek days are so-so.")

	}
}

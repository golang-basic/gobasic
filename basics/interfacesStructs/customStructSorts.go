package main

import (
	"fmt"
	"sort"
)

type class struct {
	className        string
	classSize        int
	ClassRollNumbers Series
}

//a struct defined as an array of ints
type Series []int

// Methods required by sort.Interface.
func (s Series) Len() int {
	return len(s)
}
func (s Series) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s Series) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Method for printing - sorts the elements before printing.
func (s Series) String() string {

	str := "["
	for i, elem := range s {
		if i > 0 {
			str += " "
		}
		str += fmt.Sprint(elem)
	}
	return str + "]"
}

func main() {
	var series1 Series
	series1 = []int{22, 34, 23, 12, 45, 35, 14, 5, 1, 56, 15, 3, 1, 6, 8, 9, 10, 11}

	fmt.Println(series1)
	sort.Sort(series1)
	fmt.Println(series1)
}

package main

import (
	"fmt"
	//	"sort"
	"sort"
	"strconv"
)

var Ar2D [][]int

func mutateValue(pos *int) {
	*pos = *pos*100
}

/**
Here you are assigning this pointer to a new position so the value at &Ar2d[1][1] is not going to change
 */
func mutateValueAndPosition(pos *int) {
	a := 32
	pos = &(a)
}
func changeValueAtAnyIndexInArray(Ar2d [][]int) {
	fmt.Print(" rows && pointer")
	fmt.Println(Ar2d[1])
	fmt.Println(&Ar2d[1])
	fmt.Println(&Ar2d[1][1])
	mutateValue(&Ar2d[1][1])
	fmt.Println("After Mutation")
	fmt.Print(" rows && pointer")
	fmt.Println(Ar2d[1])
	fmt.Println(&Ar2d[1])
	fmt.Println(&Ar2d[1][1])
	mutateValueAndPosition(&Ar2d[1][1])
	fmt.Println("After Incorrect Mutation of position")
	fmt.Print(" rows && pointer")
	fmt.Println("Here you are assigning this pointer to a new position so the value at &Ar2d[1][1] is not going to change	")
	fmt.Println(Ar2d[1])
	fmt.Println(&Ar2d[1])
	fmt.Println(&Ar2d[1][1])
}

func main() {
	Ar2d := [][]int{
		{10, 200, 13},
		{4, 52, 6},
		{61, 7, 80},
	}
	changeValueAtAnyIndexInArray(Ar2d)
	fmt.Println("*********************************************")
	/**
	Tower of Hanoi specific
	 */
	UseOnlyInBuiltSort()
	fmt.Println("Unsorted Colummns")
	fmt.Println(Ar2d)
	exchangeValuesOnlyOfColumns_BasedOnAscendingMagnitute(&Ar2d)
	fmt.Println("Sorted Colummns")
	fmt.Println(Ar2d)
}



func exchangeValuesOnlyOfColumns_BasedOnAscendingMagnitute(Ar2d *[][]int) {

	sortOneColumn(Ar2d, 0)
	sortOneColumn(Ar2d, 1)
	sortOneColumn(Ar2d, 2)


	//	row1 := intP {&(*Ar2d)[0][0], &(*Ar2d)[1][0], &(*Ar2d)[2][0]}
	//	row1.printP()
	//	row1.print()
	//	row1 = sortPointers(&row1)
	//	row1.printP()
	//	row1.print()
	//	fmt.Println("")
	//
	//	/*
	//	Now perform exchange for values stored in pointers at row1
	//	 */
	//	exchange(row1, intP {&(*Ar2d)[0][0], &(*Ar2d)[1][0], &(*Ar2d)[2][0]})
	//	/**
	//	similarly for other rows
	//	 */
	//	row2 := sortPointers(&(intP {&(*Ar2d)[0][1], &(*Ar2d)[1][1], &(*Ar2d)[2][1]}))
	//	exchange(row2, intP {&(*Ar2d)[0][1], &(*Ar2d)[1][1], &(*Ar2d)[2][1]})
	//	row3 := sortPointers(&(intP {&(*Ar2d)[0][2], &(*Ar2d)[1][2], &(*Ar2d)[2][2]}))
	//	exchange(row3, intP {&(*Ar2d)[0][2], &(*Ar2d)[1][2], &(*Ar2d)[2][2]})
}

func sortOneColumn(Ar2d *[][]int, row int) {
	row2 := sortPointers(&(intP {&(*Ar2d)[0][row], &(*Ar2d)[1][row], &(*Ar2d)[2][row]}))
	exchange(row2, intP {&(*Ar2d)[0][row], &(*Ar2d)[1][row], &(*Ar2d)[2][row]})
	return
}

/**
arrange values in r2's pointed mem addresses per suggested by r1
 */
func exchange(r1 intP, r2 intP) {
	tmp := []int {}
	for _, a := range r1 {
		tmp = append(tmp, *a)
	}
	for i, a := range r2 {
		*a = tmp[i]
	}
}

/**
compare will take 3 pointers & return all 3 in sorted ascending order only
It wont sort the elements which is crapy
 */
func sortPointers(arP *intP) (intP) {
	sort.Sort(*arP)
	return *arP
}

//Using an example of GOlang
// func (a ByAge) Len() int           { return len(a) }
type intP []*int

func (a intP) print() {
	fmt.Print("[")
	for _, a := range a {
		fmt.Print(" " + strconv.Itoa(*a) + ",")
	}
	fmt.Print("]")
}

func (a intP) printP() {
	fmt.Print("[")
	for _, a := range a {
		fmt.Print(" ")
		fmt.Print(a)
		fmt.Print(",")
	}
	fmt.Print("]")
}

func (a intP) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a intP) Less(i, j int) bool {
	return *a[i] < *a[j]
}
func (a intP) Len() int {
	return len(a)
}

/**
DOES NOT WORK
 */
var a = 100
var b = 20
var c = 33
var rows []*int

/**
DOES NOT WORK
 */
func UseOnlyInBuiltSort() {
	rows = []*int{&a, &b, &c}
	//sort.Sort(rows) // fails []*int does not implement sort.Interface (missing Len method)
	//sort.Sort(sort.IntSlice(rows)) // fails cannot convert rows (type []*int) to type sort.IntSlice
	// line 122 has an implementation of it which does not work

	//	fmt.Println(rows)
}

/**
	This thing didnt work coz looks like I might have to implement a C library ( looked only for 2-3 mins on go source at

https://code.google.com/p/go/source/browse/#hg%2Fsrc%2Fliblink%253Fstate%253Dclosed
So we will have to implement the sort now
	 */

//func (a []*int) Len() int{
//	return len(a)
//}

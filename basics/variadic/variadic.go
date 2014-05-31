package main

import (
	"fmt"
	"strconv"
)

func multiply(nums ...int) (result int) {
	result = 1
	for _, i := range nums {
		result = result*i
	}
	return
}

var num []int
var users int

var nums, pool = []int{1, 2, 3, 4}, []int {5, 6, 7, 8}

func main() {
	res := multiply(nums...)
	fmt.Println("Result : " + strconv.Itoa(res))

	res = multiply(pool...)
	fmt.Println("Result : " + strconv.Itoa(res))

	fmt.Println("Enter the numbers ")
	fmt.Scanf("%d", &num)
	fmt.Println("Number ", num)
	res = multiply(num...)
	fmt.Println("Result : " + strconv.Itoa(res))


	fmt.Println("Enter the numbers ")
	fmt.Scanf("%d", &users)
	fmt.Println("Number ", users)
	res = multiply(users)
	fmt.Println("Result : " + strconv.Itoa(res))
}

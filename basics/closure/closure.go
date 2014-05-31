package main

import (
	"fmt"
	"strconv"
)

var f0, f1 int = 0, 1

type T func() (*[]int)

func getSeries(arr []int, n int) T {
//	fmt.Println(len(arr))
	fmt.Println(cap(arr))
	switch {
	case n == 1 :
		return func() (*[]int) {
			arr[0] = f0
			return &arr
		}

	case n == 2 :
		return func() (*[]int) {
			arr[0] = f0
			arr[1] = f1
			return &arr
		}

	case n > 2 :
		return func() (*[]int) {
			arr[0] = f0
			arr[1] = f1
			for i := 2; i < n-1; i++ {
				fmt.Println("adding"+ strconv.Itoa((arr[i-1]+arr[i])))
				arr[i] = (arr[i-1]+arr[i])
			}
			return &arr
		}
	}
	return func() (*[]int) {
		return &arr
	}
}

func mult(args ...int) (res int) {
	res = 1
	for _, i := range args {
		res = res*i
	}
	return
}


func fibbonaci0(n int) (res []int) {
	res = make([]int, n, n)
	p := getSeries(res, n)
	p()
	return
}

func main() {
	fmt.Println(fibbonaci0(0))
	fmt.Println(fibbonaci0(1))
	fmt.Println(fibbonaci0(2))
	fmt.Println(fibbonaci0(5))
}

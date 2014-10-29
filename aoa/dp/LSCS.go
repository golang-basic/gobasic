package main

import "fmt"

/**
Largest Sum Contiguous Subarray
Write an efficient program to find the sum of contiguous subarray
within a one-dimensional array of numbers which has the largest sum.
 */

func main() {
	a := []int{-2, -3, 4, -1, -2, 1, 5, -3}
	res := LSCS(a)
	fmt.Println("Sum : ", res[0], ", i : ", res[1], ", j : ", res[2])
	res = LSCS_Kadane(a)
	fmt.Println("Sum : ", res[0], ", i : ", res[1], ", j : ", res[2])

	a = []int{-2, -3, -4, -1, -2, -1, -5, -3}
	res = LSCS(a)
	fmt.Println("Sum : ", res[0], ", i : ", res[1], ", j : ", res[2])
	res = LSCS_Kadane(a)
	fmt.Println("Sum : ", res[0], ", i : ", res[1], ", j : ", res[2])
}

func LSCS(a []int) (res [3]int) {
	res = [3]int{}
	var pre int
	for i := range a {
		pre = a[i]
		for j := i + 1; j < len(a); j++ {
			pre = pre+a[j]
			if (res[0] < pre) {
				res[0] = pre
				res[1] = i
				res[2] = j
			}
		}
	}
	return
}

func LSCS_Kadane(a []int) (res [3]int) {
	max_ending_here := a[0] // compute the new max with a new window starting at i which is 0 in this case.
	max_so_far := a[0] // max consecutive sum which is greatest so far
	max_start := 0
	max_end := 0
	for i := range a {
		max_ending_here = max(a[i], max_ending_here+a[i])
		if (max_ending_here == a[i]) {
			max_start = i
		}
		max_so_far = max(max_so_far, max_ending_here)
		if (max_so_far == max_ending_here) {
			max_end = i
		}
	}
	return [3]int{max_so_far, max_start, max_end}
}

func max(a int, b int) int {
	if a > b {
		return a
	}else {
		return b
	}
}

func PrintMatrix(M [][]int) {
	for i := range M {
		fmt.Println()
		for j := range M[0] {
			fmt.Print("  ", M[i][j])
		}
	}
}

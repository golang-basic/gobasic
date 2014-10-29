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
}

func LSCS(a []int) (res [3]int) {
	res = [3]int{}
	/**
	res[0] = max sum
	res[1] = max start index
	res[2] = max end index
	S[i][j] where the value S[i][j] represents the value
	 of sum starting from i to j
	 */
//	S := make([][]int, len(a))
//	for i := range S {
//		S[i] = make([]int, len(a))
//	}
	//	max_sum := a[0]
	res[0] = 0
	var pre int
	for i := range a {
//		S[i][0] = a[i]
		pre = a[i]
		for j := i+1; j < len(a); j++ {
//			S[i][j] = S[i][j-1]+a[j]
			pre = pre+a[j]
			if (res[0] < pre) {
				res[0] = pre
				res[1] = i
				res[2] = j
			}

		}
	}
//	PrintMatrix(S)
	return

}


func PrintMatrix(M [][]int) {
	for i := range M {
		fmt.Println()
		for j := range M[0] {
			fmt.Print("  ", M[i][j])
		}
	}
}

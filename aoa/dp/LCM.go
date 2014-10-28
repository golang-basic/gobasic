package main

import ("fmt"

)

func main() {
	M := [][]byte{[]byte{0, 1, 1, 0, 1},
		[]byte{1, 1, 0, 1, 0},
		[]byte{0, 1, 1, 1, 0},
		[]byte{1, 1, 1, 1, 0},
		[]byte{1, 1, 1, 1, 1},
		[]byte{0, 0, 0, 0, 0}}

	res := LCM(M)
	fmt.Println("Order of Largest square matrix with 1's ", res[0])
	PrintMatrixBackwards(M, int(res[0]), int(res[1]), int(res[2]))
}

func LCM(M [][]byte) []byte {
	S := make([][]byte, len(M))
	for i := range S {
		S[i] = make([]byte, len(M[0]))
	}
	S[0] = M[0]
	for i := range S[0] {
		S[i][0] = M[i][0]
	}
	res := []byte{0, 0, 0}// value, index, index
	for i := range S {
		if (i == 0) {
			continue
		}
		for j := range S[0] {
			if (j == 0) {
				continue
			}
			if M[i][j] == 1 {
				S[i][j] = min(S[i-1][j-1], S[i][j-1], S[i-1][j])+1
				if (res[0] < S[i][j]) {
					res[0] = S[i][j]
					res[1] = byte(i)
					res[2] = byte(j)
				}
			}else {
				S[i][j] = 0
			}
		}
	}
	return res
}

func min(a byte, b byte, c byte) byte {
	m := a
	if (a > b) {
		m = b
	}
	if (b > c) {
		m = c
	}
	return m
}

func PrintMatrix(M [][]byte) {
	for i := range M {
		fmt.Println()
		for j := range M[0] {
			fmt.Print(" ", M[i][j])
		}
	}
}

func PrintMatrixBackwards(M [][]byte, size int, row int, col int) {
	for i := row; i > row-size; i-- {
		for j := col; j > col-size; j-- {
			fmt.Print(" ", M[i][j])
		}
		fmt.Println()
	}
}



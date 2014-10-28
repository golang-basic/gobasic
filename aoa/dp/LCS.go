package main

import "fmt"

/**
Let Xi be the subsequence from X1 to Xi and X0 = empty
Similarly Yi
C is 2-D matrix responsible for holding the length of longest common subsequence LCS

C[i,j] is the length of the longest common subsequence between Xi and Yj
C[0,i] and C[0,j] == 0 .. no common subsequence if one of the sequence is empty
C[i,j] = 0 if i =0 or j =0
C[i,j] = C[i-1,j-1] + 1 C[i-1,j-1]-> previous diagonal element if xi = yj
C[i,j] = Max(C[i,j-1],C[i-1,j]) if xi !=yi i,j>0
I is another 2 D char matrix having I[0,i] and I[0,j] == nil
It holds either r or c or d for row, column and diagonal respectively.
It will be use for back tracking from I[n-1][n-1] back to top, to get the LCS

Once there is an entry in C[n,m] the LCS is obtained by tracing the path from c[n,m]
yo top row or leftmost column

The value of C[n,m] is the length of LCS
 */


func main() {
	fmt.Println("Finding Longest Common Subsequence in given two subsequences")
	X := [...]string{"", "a", "h", "p", "e", "u", "l", "a", "l", "g", "o", "w", "c", "t"}
	Y := [...]string{"", "h", "v", "e", "a", "l", "q", "l", "p", "u", "o", "c", "t"}
	var longestLength int
	if (len(X) > len(Y)) {
		longestLength = len(X)
	}else {
		longestLength = len(Y)
	}
	//	var LCS [longestLength]string
	fmt.Println(longestLength)
	C, I := LCS(X[:], Y[:], len(X), len(Y))
	fmt.Printf("%d", len(C))
	fmt.Printf("%d", len(I))

	fmt.Printf("Longest Sequence length %d", C[len(X)-1][len(Y)-1])
	PrintLCS(I, X[:], len(X)-1, len(Y)-1)
}

func LCS(X []string, Y []string, rows int, columns int) ([][]int, [][]string) {

	C := make([][]int, rows)
	I := make([][]string, rows)

	for i := range C {
		C[i] = make([]int, columns)
		I[i] = make([]string, columns)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if (i == 0 || j == 0) {
				C[i][j] = 0
				I[i][j] = ""
				continue
			}
			if X[i] == Y[j] {
				C[i][j] = C[i-1][j-1]+1
				I[i][j] = "d"
			}else if (X[i] != Y[j]) {
				if (C[i][j-1] > C[i-1][j]) {
					C[i][j] = C[i][j-1]
					I[i][j] = "c"
				}else {
					C[i][j] = C[i-1][j]
					I[i][j] = "r"
				}
			}
		}
	}
	return C, I
}

func PrintLCS(I [][]string, X []string, i int, j int) {
	if i == 0 || j == 0 {
		return
	}
	if I[i][j] == "d" {
		PrintLCS(I, X, i-1, j-1)
		fmt.Print(" ", X[i])
	}
	if I[i][j] == "r" {
		PrintLCS(I, X, i-1, j)
	}
	if I[i][j] == "c" {
		PrintLCS(I, X, i, j-1)
	}

}

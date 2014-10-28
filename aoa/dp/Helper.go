package dp

import "fmt"

func PrintMatrix(M [][]byte) {
	for i := range M {
		for j := range M[0] {
			fmt.Println(M[i][j])
		}
	}
}

package main

import "fmt"

func main() {

	var n, m, q int
	fmt.Scan(&n, &m, &q)
	var D = make([][]int, n)
	var F = make([][]string, n)
	for i := 0; i < n; i++ {
		D[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&D[i][j])
		}
	}
	for i := 0; i < n; i++ {
		F[i] = make([]string, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&F[i][j])
		}
	}
	fmt.Printf("digraph {\n" + "\trankdir = LR\n")
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("\t%d -> %d [label = \"%s(%s)\"]\n", i, D[i][j], string(rune(j+'a')), F[i][j])
		}
	}
	fmt.Println("}")
}

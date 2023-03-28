package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var k = 0

func DFS(D [][]int, renumeration []int, q int) {
	renumeration[q] = k
	k++
	for _, v := range D[q] {
		if renumeration[v] == -1 {
			DFS(D, renumeration, v)
		}
	}
}

func main() {
	var n, m, q int
	bufstdin := bufio.NewReader(os.Stdin)
	bufstdout := bufio.NewWriter(os.Stdout)
	fmt.Fscan(bufstdin, &n, &m, &q)
	var D = make([][]int, n)
	var F = make([][]string, n)
	var renumeration = make([]int, n)
	var reRenumeration = make([]int, n)
	for i := 0; i < n; i++ {
		renumeration[i] = -1
		D[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(bufstdin, &D[i][j])
		}
	}
	for i := 0; i < n; i++ {
		F[i] = make([]string, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(bufstdin, &F[i][j])
		}
	}
	DFS(D, renumeration, q)
	for i, x := range renumeration {
		if x != -1 {
			reRenumeration[x] = i
		} else {
			n--
		}
	}
	bufstdout.WriteString(strconv.Itoa(n) + "\n" + strconv.Itoa(m) + "\n" + strconv.Itoa(0) + "\n")
	for i := 0; i < n; i++ {
		k := reRenumeration[i]
		for j := 0; j < m; j++ {
			bufstdout.WriteString(strconv.Itoa(renumeration[D[k][j]]) + " ")
		}
		bufstdout.WriteString("\n")
	}
	for i := 0; i < n; i++ {
		k := reRenumeration[i]
		for j := 0; j < m; j++ {
			bufstdout.WriteString(F[k][j] + " ")
		}
		bufstdout.WriteString("\n")
	}
	bufstdout.Flush()
}

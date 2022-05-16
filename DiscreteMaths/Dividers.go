package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	dividers := make([]int, 0)
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			if i*i == n {
				dividers = append(dividers, i)
			} else {
				dividers = append(dividers, i)
				dividers = append(dividers, n/i)
			}
		}
	}
	sort.Ints(dividers)
	fmt.Printf("graph {\n")
	for _, elem := range dividers {
		fmt.Printf("\t%d\n", elem)
	}
	length := len(dividers)
	for i := length - 1; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if dividers[i]%dividers[j] == 0 {
				for k := i - 1; k >= j; k-- {
					if k == j {
						fmt.Printf("\t%d--%d\n", dividers[i], dividers[j])
					} else {
						if dividers[k]%dividers[j] == 0 && dividers[i]%dividers[k] == 0 {
							break
						}
					}
				}
			}
		}
	}
	fmt.Printf("}")
}

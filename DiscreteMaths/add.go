package main

import "fmt"

func add(a, b []int32, p int) []int32 {
	var length, over int
	over = 0
	if len(a) < len(b) {
		length = len(b)
	} else {
		length = len(a)
	}
	c := make([]int32, length)
	for i := 0; i < length; i++ {
		if i >= len(a) {
			c[i] = b[i] + int32(over)
		} else if i >= len(b) {
			c[i] = a[i] + int32(over)
		} else {
			c[i] = a[i] + b[i] + int32(over)
		}
		if c[i] >= int32(p) {
			over = 1
			c[i] = c[i] - int32(p)
		} else {
			over = 0
		}
	}
	return c
}

func main() {
	a := make([]int32, 5)
	b := make([]int32, 7)
	for i := 0; i < 5; i++ {
		fmt.Scan(&a[i])
	}
	for i := 0; i < 7; i++ {
		fmt.Scan(&b[i])
	}
	p := 10
	add(a, b, p)
	fmt.Print(add(a, b, p))
}

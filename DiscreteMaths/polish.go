package main

import (
	"bufio"
	"fmt"
	"os"
)

func Scan1() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	return in.Text()
}

var i int

func Calculate(s string) int {
	i++
	if s[0] != '(' {
		return int(s[0] - '0')
	}
	if i >= len(s) {
		return 0
	}
	if s[i] == '*' {
		return Calculate(s) * Calculate(s)
	}
	if s[i] == '+' {
		return Calculate(s) + Calculate(s)
	}
	if s[i] == '-' {
		return Calculate(s) - Calculate(s)
	}
	if s[i] == '(' || s[i] == ')' || s[i] == ' ' {
		return Calculate(s)
	}
	return int(s[i] - '0')
}

func main() {
	var s string
	s = Scan1()
	fmt.Printf("%d\n", Calculate(s))
}

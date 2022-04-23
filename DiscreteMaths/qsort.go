package main

import "fmt"

var arr [10]int

func less(i, j int) bool {
	return arr[i] < arr[j]
}

func swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func quicksort(low, high int, less func(i, j int) bool, swap func(i, j int)) {
	i := low
	j := high
	pivot := (i + j) / 2
	for i <= j {
		for less(pivot, j) {
			j--
		}
		for less(i, pivot) {
			i++
		}
		if i <= j {
			swap(i, j)
			i++
			j--
		}
	}
	if low < j {
		quicksort(low, j, less, swap)
	}
	if i < high {
		quicksort(i, high, less, swap)
	}
}

func qsort(n int, less func(i, j int) bool, swap func(i, j int)) {
	quicksort(0, n-1, less, swap)
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Scan(&arr[i])
	}
	qsort(10, less, swap)
	for i := 0; i < 10; i++ {
		fmt.Print(arr[i], " ")
	}
}

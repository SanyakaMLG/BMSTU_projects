package main

import "fmt"

var arr = []int{
	-503, -521, -587, -610, -515, -510, -533, -549, -600, -557,
	-533, -571, -597, -534, -516, -544, -601, -550, -605, -566,
	-523, -534, -540, -581, -500, -561,
}

func less(i, j int) bool {
	return arr[i] < arr[j]
}

func swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func partition(low, high int, less func(i, j int) bool, swap func(i, j int)) int {
	i := low
	for j := low; j < high; j++ {
		if less(j, high) {
			swap(i, j)
			i++
		}
	}
	swap(i, high)
	return i
}

func quicksort(low, high int, less func(i, j int) bool, swap func(i, j int)) {
	if low < high {
		q := partition(low, high, less, swap)
		quicksort(low, q-1, less, swap)
		quicksort(q+1, high, less, swap)
	}
}

func qsort(n int, less func(i, j int) bool, swap func(i, j int)) {
	quicksort(0, n-1, less, swap)
}

func main() {
	qsort(26, less, swap)
	for i := 0; i < 26; i++ {
		fmt.Print(arr[i], " ")
	}
}

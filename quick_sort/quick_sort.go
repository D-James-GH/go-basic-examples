package main

import "fmt"

func partition(arr *[]int, lo int, hi int) int {
	a := *arr
	pivot := a[hi]
	idx := lo - 1
	for i := lo; i < hi; i++ {
		if a[i] <= pivot {
			idx++
			tmp := a[i]
			a[i] = a[idx]
			a[idx] = tmp
		}
	}

	idx++
	a[hi] = a[idx]
	a[idx] = pivot

	return idx
}
func qs(arr *[]int, lo int, hi int) {
	if lo >= hi {
		return
	}
	pivotIdx := partition(arr, lo, hi)

	qs(arr, lo, pivotIdx-1)

	qs(arr, pivotIdx+1, hi)
}

func edit(arr *[]int) {
	(*arr)[0] = 333
}

func main() {
	arr := []int{9, 44, 4, 7, 69, 7, 420, 32}

	qs(&arr, 0, len(arr)-1)
	// edit(&arr)
	fmt.Println(arr)
}

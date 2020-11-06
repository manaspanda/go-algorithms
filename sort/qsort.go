package main

import "fmt"

func partition(a []int, start, end int) int {
	pivot := a[end]
	pIndex := start

	for i := start; i < end; i++ {
		if a[i] < pivot {
			a[i], a[pIndex] = a[pIndex], a[i]
			pIndex++
		}
	}
	a[pIndex], a[end] = a[end], a[pIndex]
	fmt.Printf("q-sort: %v\n", a)
	return pIndex
}

func qsort(a []int, start, end int) {
	if start >= end {
		return
	}

	pivot := partition(a, start, end)
	qsort(a, start, pivot-1)
	qsort(a, pivot+1, end)
}

func test_qsort() {

	fmt.Println("-------- Quicksort ---------")
	a := []int{5, 10, 1, 4, 9, 3, 6, 5, 7}
	fmt.Printf("%v\n", a)
	qsort(a, 0, len(a)-1)
}

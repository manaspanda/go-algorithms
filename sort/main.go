package main

import "fmt"

func main() {
	a := []int{5, 10, 1, 4, 9, 3, 6, 5, 7}
	b := []int{5, 10, 1, 4, 9, 3, 6, 5, 7}
	c := []int{5, 10, 1, 4, 9, 3, 6, 5, 7}
	//a := []int{5, 10, 1, 4, 9, 3, 6, 5, 7}
	//a := []int{5, 10, 1, 4, 9, 3, 6, 5, 7, 20, 15, 100, 11}
	fmt.Printf("-------- selection-sort --------\n%v\n", a)
	selectionSort(a)
	fmt.Printf("-------- insertion-sort --------\n%v\n", b)
	insertionSort(b)
	fmt.Printf("-------- bubble-sort --------\n%v\n", c)
	bubbleSort(c)

	test_qsort()

	test_sortpkg()
}

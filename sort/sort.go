package main

import (
	"fmt"
)

// Insertion sort
// 	{5, 10, 1, 4, 9, 3, 6, 5, 7}
func insertionSort(a []int) {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0; j-- {
			if a[j-1] > a[j] {
				temp := a[j]
				a[j] = a[j-1]
				a[j-1] = temp
			}
		}
		fmt.Printf("i-sort: %v\n", a)
	}
}

// Selection sort
func selectionSort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[i] {
				temp := a[j]
				a[j] = a[i]
				a[i] = temp
			}
		}
		fmt.Printf("s-sort: %v\n", a)
	}
}

// Bubble sort
func bubbleSort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-1; j++ {
			if a[j] > a[j+1] {
				temp := a[j]
				a[j] = a[j+1]
				a[j+1] = temp
			}
		}
		fmt.Printf("b-sort: %v\n", a)
	}

}

func main() {
	a := []int{5, 10, 1, 4, 9, 3, 6, 5, 7}
	b := []int{5, 10, 1, 4, 9, 3, 6, 5, 7}
	c := []int{5, 10, 1, 4, 9, 3, 6, 5, 7}
	//a := []int{5, 10, 1, 4, 9, 3, 6, 5, 7}
	//a := []int{5, 10, 1, 4, 9, 3, 6, 5, 7, 20, 15, 100, 11}
	fmt.Printf("unsorted: %v\n", a)
	selectionSort(a)
	fmt.Printf("selection-sorted: %v\n", a)
	insertionSort(b)
	fmt.Printf("insertion-sorted: %v\n", b)
	bubbleSort(c)
	fmt.Printf("bubble-sorted: %v\n", c)

	test_qsort()
}

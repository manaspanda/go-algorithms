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
				a[j], a[j-1] = a[j-1], a[j]
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
				a[j], a[i] = a[i], a[j]
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
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
		fmt.Printf("b-sort: %v\n", a)
	}

}

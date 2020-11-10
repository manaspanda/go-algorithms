package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}
type Persons []*Person
type ByAge struct{ Persons }

func (p Persons) Len() int         { return len(p) }
func (p Persons) Swap(i, j int)    { p[i], p[j] = p[j], p[i] }
func (p ByAge) Less(i, j int) bool { return p.Persons[i].Age < p.Persons[j].Age }

func PrintPerson(p Persons) {
	for i := 0; i < len(p); i++ {
		fmt.Printf("%v, ", *p[i])
	}
}

func test_sortpkg() {
	fmt.Println("\n-------- go sort pkg --------")

	// Sort int array
	inta := []int{5, 10, 1, 4, -9, 3, 6, -5, 7}
	fmt.Printf("\nSort %v\n", inta)
	sort.Ints(inta)
	fmt.Printf("==> %v\n", inta)

	// Sort string array
	stra := []string{"one", "two", "three", "four", "five", "six"}
	fmt.Printf("\nSort %v\n", stra)
	sort.Strings(stra)
	fmt.Printf("==> %v\n", stra)

	// Sort Persons array
	family := []*Person{
		{"Seema", 48}, {"Anshul", 19}, {"Manas", 50}, {"Nola", 2}, {"Richa", 15},
	}
	fmt.Printf("\nSort ")
	PrintPerson(family)
	sort.Sort(ByAge{family})
	fmt.Printf("\n==> ")
	PrintPerson(family)
	fmt.Printf("\n")
}

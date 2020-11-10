package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}
type Persons struct {
	list []*Person
}

func (p Persons) Len() int           { return len(p.list) }
func (p Persons) Swap(i, j int)      { p.list[i], p.list[j] = p.list[j], p.list[i] }
func (p Persons) Less(i, j int) bool { return p.list[i].Age < p.list[j].Age }

func (p Persons) Print() {
	for i := 0; i < len(p.list); i++ {
		fmt.Printf("%v, ", *p.list[i])
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
	family := Persons{
		list: []*Person{
			{"Sandy", 48}, {"Andrew", 19}, {"Manuel", 50}, {"Nola", 2}, {"Rachel", 15},
		},
	}
	fmt.Printf("\nSort ")
	family.Print()
	sort.Sort(family)
	fmt.Printf("\n==> ")
	family.Print()
	fmt.Printf("\n")
}

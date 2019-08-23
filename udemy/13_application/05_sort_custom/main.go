package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByAge []Person

func (a ByAge) Len() int {
	return len(a)
}
func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func main() {
	People := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}
	fmt.Println("People1 =", People)
	sort.Sort(ByAge(People))
	fmt.Println("People2 =", People)
	// There are two ways to sort a slice. First, one can define
	// a set of methods for the slice type, as with ByAge, and
	// call sort.Sort. In this first example we use that technique.

	// =================================================================================
	// The other way is to use sort.Slice with a custom Less
	// function, which can be provided as a closure. In this
	// case no methods are needed. (And if they exist, they
	// are ignored.) Here we re-sort in reverse order: compare
	// the closure with ByAge.Less.
	fmt.Println("=================================================================================")
	sort.Slice(People, func(i, j int) bool {
		return People[i].Age < People[j].Age
	})
	fmt.Println("People other way sort =", People)
	fmt.Println("=================================================================================")

	sort.Slice(People, func(i, j int) bool {
		return People[i].Name < People[j].Name
	})
	fmt.Println("People Name sort =", People)
}

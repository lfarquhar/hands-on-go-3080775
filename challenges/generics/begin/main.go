// challenges/generics/begin/main.go
package main

import (
"fmt"
"golang.org/x/exp/constraints"
)

// Part 1: print function refactoring

// non-generic print functions
func print[T any](val T) { fmt.Println(val) }

type numeric interface {
	constraints.Integer | constraints.Float
}



// Part 2 sum function refactoring
func sum[T numeric](numbers ...T) T {
	var s T
	for _, n := range numbers {
		s += n
	}
	return s
}

func main() {
	// call generic printAny function for each value above
	// print("Hello")
	// print(42)
	// print(true)
	
	// call generics sumAny function
	fmt.Println(sum(1, 2, 3))
}

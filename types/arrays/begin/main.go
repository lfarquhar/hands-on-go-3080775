// types/arrays/begin/main.go
package main

import "fmt"

func main() {
	// declare an array of integers - like other languages it remains the same size
	var a [3] int

	// print the array
	fmt.Println(a)

	// set the first element to 1
	a[0] = 1

	// print the array
	fmt.Println(a)
}

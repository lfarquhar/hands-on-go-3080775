// flow-control/defer-panic-recover/begin/main.go
package main

import "fmt"

func cleanup(msg string) {
	fmt.Println(msg)
}

// kind of like try catch finally but go wants us to handle errors like values so this should be used sparingly

func main() {
	// defer function calls
	defer cleanup("first")
	defer cleanup("second")
	
	fmt.Println("do work")

	// defer recovery
	defer func ()  {
		if err := recover(); err != nil {
			fmt.Println("recovered from:", err)			
		}
	}()

	// panic
	panic("bad stuff")
}

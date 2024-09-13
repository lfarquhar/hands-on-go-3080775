// challenges/types-primitive/begin/main.go
package main

import "fmt"

// use type inference to create a package-level variable "val" and assign it a value of "global"
var val = "24"

func main() {
	// create a local variable "val" and assign it an int value	
	val := 42

	// print the value of the local variable "val"
	fmt.Println(val)

	// print the value of the package-level variable "val"
	fmt.Println(getGlobalVal())

	// update the package-level variable "val" and print it again
	setGlobalVal("foo")
	fmt.Println(getGlobalVal())
	
	// print the pointer address of the local variable "val"
	fmt.Println(&val)

	// assign a value directly to the pointer address of the local variable "val" and print the value of the local variable "val"
	*(&val) = 100 // * is deference
	fmt.Println("local val = ", val)
}

func getGlobalVal() string {
	return val
}

func setGlobalVal(value string) {
	val = value
}

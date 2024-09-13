// functions/begin/main.go
package main

import (
	"errors"
	"fmt"
)

// simple greet function
//

// greetWithName returns a greeting with the name
//

// greetWithName returns a greeting with the name and age of the person
//

// divide divides two numbers and returns the result
// if the second number is zero, it returns  error
//

func foo(b bool)(string, error){
	
	if (b){
		return "", errors.New("ouch!")
	}
	
	return "foo", nil
}

func main() {
	
	resultStr, err := foo(false)

	fmt.Println(resultStr, err)


	resultStr, err = foo(true)

	fmt.Println(resultStr, err)

	// invoke greet function
	// fmt.Println(greet())

	// invoke greetWithName function
	// fmt.Println(greetWithName("Toni"))

	// invoke divide function
	// fmt.Println(divide(10, 2))

	// invoke divide function with zero denominator to get an error
	// fmt.Println(divide(5, 0))
}

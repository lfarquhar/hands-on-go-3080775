// interfaces/type-switch/begin/main.go
package main

import "fmt"

// define whatAmI which takes in an argument of any type and returns inforamtion about the underlying value's type
func whatAmI(val interface{}) string {
	switch val.(type){
	case int:
		return fmt.Sprint("%v is an int", val)
	case string:
		return fmt.Sprint("%v is a string", val)
	default:
		return fmt.Sprintf("%v is a %T", val, val)
	}
}

type Book struct{
	name string
}

func main() {
	// invoke whatAmI function
	fmt.Println(whatAmI(1))
	fmt.Println(whatAmI("hello"))
	fmt.Println(whatAmI(true))
	
	b := Book{ name: "foo"}
	fmt.Println(whatAmI(b))
}

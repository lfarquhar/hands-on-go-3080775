// interfaces/basics/begin/main.go
package main

import "fmt"

// define a humanoid interface with speak and walk methods returning string
type humanoid interface {
	speak()
	walk()
}

// define a person type that implements humanoid interface
type person struct {
	name string
}

func (p person) speak() {
	fmt.Printf("%s doesn't have much to say... ", p.name)
}

func (p person) walk() {
	fmt.Printf("%s is walking...", p.name)
}

func (p person) String() string {
	return fmt.Sprintf("Hello, my name is %s", p.name)
}

type dog struct{}

func (d dog) walk() { fmt.Println("Dog walking") }

// implement the Stringer interface for the person type

// define a dog type that can walk but not speak

func main() {
	p := person{name: "John"}

	// invoke with a person
	// doHumanThings(p)

	// d := dog{}
	// can we invoke with a dog?
	// doHumanThings(d)

	fmt.Println(p)
}

func doHumanThings(h humanoid) {
	h.speak()
	h.walk()
}

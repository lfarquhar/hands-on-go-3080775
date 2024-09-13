// concurrency/channels/begin/main.go
package main

import (
	"fmt"
)

// sum calculates and prints the sum of numbers
func sum(nums []int, ch chan<- int) {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	ch <- sum
}

func main() {
	nums := []int{1, 2, 3, 4, 5}

	ch := make(chan int)

	// invoke the sum function as a goroutine
	go sum(nums, ch) // unbuffered channel will block until there is a receiver

	result := <-ch

	fmt.Println("Result: ", result)

	ch2 := make(chan string, 2)

	ch2 <- "one"
	ch2 <- "two"

	fmt.Println(<- ch2, <- ch2)
}

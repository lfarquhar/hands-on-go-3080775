// concurrency/channel-non-blocking/begin/main.go
package main

import (
	"fmt"
	"time"
)

// unbuffered channel will block until there is a receiver

func main() {
	// declare a signal channel to read from
	timeChan1 := time.After(200 * time.Millisecond)
	timeChan2 := time.After(200 * time.Millisecond)

	for{
		select {
		case <-timeChan1:
			fmt.Println("timeChan1 returned")
			return
		case <-timeChan2:
			fmt.Println("timeChan2 returned")
			return
		default:
			fmt.Println("default case")
			time.Sleep(250 * time.Millisecond)
		}
	}
}

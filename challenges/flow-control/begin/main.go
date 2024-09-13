// challenges/flow-control/begin/main.go
package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	// handle any panics that might occur anywhere in this function
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("don't panic. I got you.")
		}
	}()

	// use os package to read the file specified as a command line argument
	fileName := os.Args[1]

	fileBytes, err := os.ReadFile(fileName)

	if err != nil {
		panic(fmt.Errorf("failed to read file: %s", err))
	}
	// convert the bytes to a string
	data := string(fileBytes)

	// I got the data spew.Dump(data)
	// no errors fmt.Println(err)

	// initialize a map to store the counts
	m := make(map[string]int)

	// err := json.Unmarshal(input, &m)

	// use the standard library utility package that can help us split the string into words
	words := strings.Fields(data)

	// capture the length of the words slice
	m["words"] = len(words)

	// use for-range to iterate over the words slice and for each word, count the number of letters, numbers, and symbols, storing them in the map
	for _, word := range words {
		for _, char := range word {
			if unicode.IsLetter(char) {
				m["letters"]++
			} else if unicode.IsNumber(char) {
				m["numbers"]++
			} else {
				m["symbols"]++
			}
		}
	}

	// dump the map to the console using the spew package
	spew.Dump(m)
}

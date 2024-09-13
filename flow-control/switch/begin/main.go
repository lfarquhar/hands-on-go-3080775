// flow-control/switch/begin/main.go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	os := runtime.GOOS

	// switch os := runtime.GOOS; os {

	switch os {
	case "linux", "darwin", "unix":
		fmt.Println("*nix variant")
	case "windows":
		fmt.Println("windows")
	default:
		fmt.Printf("%s.\n", os)

	}
}

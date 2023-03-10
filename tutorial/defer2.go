package main

import "fmt"

func main() {
	// It is possible to call defer multiple times in a function.
	defer fmt.Println("Begin")
	defer fmt.Println("End")
}

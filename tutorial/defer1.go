package main

import "fmt"

func main() {
	fmt.Println("Begin")
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
	fmt.Println("End")
}

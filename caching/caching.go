package main

import "fmt"

func prettyPrint(f float64) {
	fmt.Printf("> %f", f)
}

func main() {
	prettyPrint(.25)
}
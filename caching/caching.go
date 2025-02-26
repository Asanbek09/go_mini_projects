package main

import "fmt"

func prettyPrint(f float64) {
	fmt.Printf("> %f", f)
}

func prettyPrint2[T any](t T) {
	fmt.Printf("> %v", t)
}

func main() {
	prettyPrint(.25)

	prettyPrint2[float64](.25)
	prettyPrint2[string]("pockets")
}
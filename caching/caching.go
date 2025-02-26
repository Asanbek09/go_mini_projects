package main

import "fmt"

type Printable interface {
	PrettyPrint()
}

type Cloud struct {
	Value float64
}

func (c Cloud) PrettyPrint() {
	fmt.Printf("> %f\n", c.Value)
}

type Group[T Printable] []T

func (g Group[T]) PrettyPrint() {
	for _, v := range g {
		v.PrettyPrint()
	}
}

func main() {
	g := Group[Cloud]{
		{Value: 1.23},
		{Value: 4.56},
	}
	g.PrettyPrint()
}
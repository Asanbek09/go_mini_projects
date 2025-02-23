package main

import (
	"flag"
	"fmt"
)

func main() {
	from := flag.String("from", "", "source currency, required")
	to := flag.String("to", "EUR", "target currency")

	flag.Parse()

	fmt.Println(*from, *to)
}
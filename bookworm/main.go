package main

import (
	"fmt"
	"os"
)

func main() {
	bookworms, err := loadBookworms("database/data.json")
	if err != nil {
		_, _ =fmt.Fprintf(os.Stderr, "failed to load bookworms: %s\n", err)
		os.Exit(1)
	}

	fmt.Println(bookworms)
}
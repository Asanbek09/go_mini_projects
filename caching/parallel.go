package main

import (
	"fmt"
	"time"
)

func printEverySecond(msg string) {
	for i := 0; i < 10; i++ {
		fmt.Println(msg)
		time.Sleep(time.Second)
	}
}

func main() {
	go printEverySecond("Hello")
	go printEverySecond("World")

	var input string
	fmt.Scanln(&input)
}

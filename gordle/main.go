package main

import (
	"fmt"
)

func main() {
    fmt.Println("Hello, 世界")
    fmt.Println(len("Hello, 世界"))
    fmt.Println(len([]rune("Hello, 世界")))
}
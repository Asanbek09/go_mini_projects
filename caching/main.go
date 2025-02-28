package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go cookRice(wg)
	go cookCurry(wg)

	wg.Wait()
}

func cookRice(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Cooking rice...")
}

func cookCurry(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Preparing curry sauce...")
}
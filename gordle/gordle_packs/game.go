package gordlepacks

import (
	"bufio"
	"fmt"
	"io"
)

type Game struct {
	reader *bufio.Reader
}

func New(playerInput io.Reader) *Game {
	g := &Game{
		reader: bufio.NewReader(playerInput),
	}

	return g
}

func (g *Game) Play() {
	fmt.Println("Welcome to Gordle!")

	fmt.Printf("Enter a guess : \n")
}
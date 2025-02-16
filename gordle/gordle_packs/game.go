package gordlepacks

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

const solutionLength = 5

func (g *Game) ask() []rune {
	fmt.Printf("Enter a %d-character guess: \n", solutionLength)

	for {
		playerInput, _, err := g.reader.ReadLine()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Gordle failed to read your guess: %\n", err.Error())
			continue
		}

		guess := []rune(string(playerInput))

		if len(guess) != solutionLength {
			_, _ = fmt.Fprintf(os.Stderr, "Your attempt is invalid with Gordle's solution! Expected %d characters, got %d\n", solutionLength, len(guess))
		} else {
			return guess
		}
	}
}
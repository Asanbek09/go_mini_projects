package gordlepacks

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Game struct {
	reader *bufio.Reader
	solution []rune
	maxAttempts int
}

func New(playerInput io.Reader, solution string, maxAttempts int) *Game {
	g := &Game{
		reader: bufio.NewReader(playerInput),
		solution: splitToUppercaseCharacters(solution),
		maxAttempts: maxAttempts,
	}

	return g
}

func (g *Game) Play() {
	fmt.Println("Welcome to Gordle!")

	guess := g.ask()

	fmt.Printf("Your guess is: %s\n", string(guess))
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

		err = g.validateGuess(guess)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Your attempt is invalid with gordle's solution: %s\n", err.Error())
		} else {
			return guess
		}
	}
}

var errInvalidWordLength = fmt.Errorf("Invalid guess, word doesn't have the same number of charaters as the solution")

func (g *Game) validateGuess(guess []rune) error {
	if len(guess) != solutionLength {
		return fmt.Errorf("Expected %d, got %d, %w", solutionLength, len(guess), errInvalidWordLength)
	}

	return nil
}

func splitToUppercaseCharacters(input string) []rune {
	return []rune(strings.ToUpper(input))
}
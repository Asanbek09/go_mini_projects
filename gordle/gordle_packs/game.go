package gordlepacks

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
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

	for currentAttempt := 1; currentAttempt <= g.maxAttempts; currentAttempt++ {
		guess := g.ask()

		if slices.Equal(guess, g.solution) {
			fmt.Printf("You won! You found it in %d guess(es)! The word has: %s\n", currentAttempt, string(g.solution))
			return
		}
	}

	fmt.Printf("You have lost! The solution was: %s\n", string(g.solution))
}

const solutionLength = 5

func (g *Game) ask() []rune {
	fmt.Printf("Enter a %d-character guess: \n", len(g.solution))

	for {
		playerInput, _, err := g.reader.ReadLine()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Gordle failed to read your guess: %\n", err.Error())
			continue
		}

		guess := splitToUppercaseCharacters(string(playerInput))

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
	if len(guess) != len(g.solution) {
		return fmt.Errorf("Expected %d, got %d, %w", len(g.solution), len(guess), errInvalidWordLength)
	}

	return nil
}

func splitToUppercaseCharacters(input string) []rune {
	return []rune(strings.ToUpper(input))
}
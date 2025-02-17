package gordlepacks

import (
	"fmt"
	"os"
	"strings"
)

type hint byte
type feedback []hint

const (
	absentCharacter hint = iota
	wrongPosition
	correctPosition
)

func (h hint) String() string {
	switch h {
	case absentCharacter:
		return "🎉"
	case wrongPosition:
		return "🤡"
	case correctPosition:
		return "🥳"
	default:
		return "❤️"
	}
}

func (fb feedback) StringConcat() string {
	var output string
	for _, h := range fb {
		output += h.String()
	}

	return output
}

func (fb feedback) String() string {
	sb := strings.Builder{}
	for _, h := range fb {
		sb.WriteString(h.String())
	}

	return sb.String()
}

func computeFeedback(guess, solution []rune) feedback {
	result := make(feedback, len(guess))
	used := make([]bool, len(solution))

	if len(guess) != len(solution) {
		_, _ = fmt.Fprintf(os.Stderr, "Internal error! Guess and solution have different lengths: %d vs %d", len(guess), len(solution))
		return result
	}

	for posInGuess, character := range guess {
		if character == solution[posInGuess] {
			result[posInGuess] = correctPosition
			used[posInGuess] = true
		}
	}

	for posInGuess, character := range guess {
		if result[posInGuess] != absentCharacter {
			continue
		}

		for posInSolution, target := range solution {
			if used[posInSolution] {
				continue
			}
			if character == target {
				result[posInGuess] = wrongPosition
				used[posInSolution] = true
				break
			}
		}
	}
	return result
}

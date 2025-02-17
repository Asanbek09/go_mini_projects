package gordlepacks

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
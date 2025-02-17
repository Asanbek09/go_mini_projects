package gordlepacks

type hint byte

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
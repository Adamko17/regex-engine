package rgx

type tokenType uint8

const (
	group           tokenType = iota
	bracket         tokenType = iota // []
	or              tokenType = iota // |
	repeat          tokenType = iota
	literal         tokenType = iota
	groupUncaptured tokenType = iota
)

type token struct {
	tokenType tokenType
	value     interface{}
}

type parseContext struct {
	pos    int
	tokens []token
}
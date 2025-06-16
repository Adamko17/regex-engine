package rgx

type state struct {
	start       bool
	terminal    bool
	transitions map[uint8][]*state
}
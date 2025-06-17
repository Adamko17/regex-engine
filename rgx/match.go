package rgx

const (
	startOfText uint8 = 1
	endOfText   uint8 = 2
)

func getChar(input string, pos int) uint8 {
	if pos >= len(input) {
		return endOfText
	}

	if pos < 0 {
		return startOfText
	}

	return input[pos]
}

func (s *state) Check(input string, pos int) bool {
	if pos > len(input) {
		return false
	}

	ch := getChar(input, pos)

	if ch == endOfText && s.terminal {
		return true
	}

	if states := s.transitions[ch]; len(states) > 0 {
		nextState := states[0]
		if nextState.Check(input, pos+1) {
			return true
		}
	}

	for _, state := range s.transitions[epsilonChar] {
		if state.Check(input, pos) {
			return true
		}

		if ch == startOfText && state.Check(input, pos+1) {
			return true
		}
	}

	return false
}

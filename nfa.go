package rgx

type state struct {
	start       bool
	terminal    bool
	transitions map[uint8][]*state
}

const epsilonChar uint8 = 0 // empty char

func toNfa(ctx *parseContext) *state {
	startState, endState := tokenToNfa(&ctx.tokens[0])

	for i := 1; i < len(ctx.tokens); i++ {
		startNext, endNext := tokenToNfa(&ctx.tokens[i])
		endState.transitions[epsilonChar] = append(
			endState.transitions[epsilonChar],
			startNext,
		)
		endState = endNext
	}

	start := &state{
		transitions: map[uint8][]*state{
			epsilonChar: {startState},
		},
		start: true,
	}
	end := &state{
		transitions: map[uint8][]*state{},
		terminal:    true,
	}

	endState.transitions[epsilonChar] = append(
		endState.transitions[epsilonChar],
		end,
	)

	return start
}

// returns (start, end)
func tokenToNfa(t *token) (*state, *state) {
	start := &state{
		transitions: map[uint8][]*state{},
	}
	end := &state{
		transitions: map[uint8][]*state{},
	}

	switch t.tokenType {
	case literal:
	case or:
	case bracket:
	case group, groupUncaptured:
	case repeat:
	default:
		panic("unknown type of token")
	}

	return start, end
}
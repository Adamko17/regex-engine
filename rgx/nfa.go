package rgx

type state struct {
	start       bool
	terminal    bool
	transitions map[uint8][]*state
}

const epsilonChar uint8 = 0 // empty char

func ToNfa(ctx *parseContext) *state {
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
		ch := t.value.(uint8)
		start.transitions[ch] = []*state{end}
	case or:
		values := t.value.([]token)
		left := values[0]
		right := values[1]

		s1, e1 := tokenToNfa(&left)
		s2, e2 := tokenToNfa(&right)

		start.transitions[epsilonChar] = []*state{s1, s2}
		e1.transitions[epsilonChar] = []*state{end}
		e2.transitions[epsilonChar] = []*state{end}
	case bracket:
		literals := t.value.(map[uint8]bool)

		// transitions from start to end state
		for l := range literals {
			start.transitions[l] = []*state{end}
		}
	case group, groupUncaptured:
		tokens := t.value.([]token)
		start, end = tokenToNfa(&tokens[0])
		for i := 1; i < len(tokens); i++ {
			ts, te := tokenToNfa(&tokens[i])
			end.transitions[epsilonChar] = append(
				end.transitions[epsilonChar],
				ts,
			)
			end = te
		}
	case repeat:
		p := t.value.(repeatPayLoad)

		if p.min == 0 {
			start.transitions[epsilonChar] = []*state{end}
		}

		var copyCount int

		if p.max == repeatInfinity {
			if p.min == 0 {
				copyCount = 1
			} else {
				copyCount = p.min
			}
		} else {
			copyCount = p.max
		}

		from, to := tokenToNfa(&p.value)
		start.transitions[epsilonChar] = append(
			start.transitions[epsilonChar],
			from,
		)

		for i := 2; i <= copyCount; i++ {
			s, e := tokenToNfa(&p.value)

			to.transitions[epsilonChar] = append(
				to.transitions[epsilonChar],
				s,
			)

			from = s
			to = e

			if i > p.min {
				s.transitions[epsilonChar] = append(
					s.transitions[epsilonChar],
					end,
				)
			}
		}

		to.transitions[epsilonChar] = append(
			to.transitions[epsilonChar],
			end,
		)

		if p.max == repeatInfinity {
			end.transitions[epsilonChar] = append(
				end.transitions[epsilonChar],
				from,
			)
		}

	default:
		panic("unknown type of token")
	}

	return start, end
}
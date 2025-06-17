package rgx

const (
	startOfText uint8 = 1
	endOfText   uint8 = 2
)

func getChar(input string, pos int) uint8 {
	if pos > len(input) {
		return endOfText
	}

	if pos < 0 {
		return startOfText
	}

	return input[pos]
}
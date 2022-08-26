package lexer

const LPAREN = "{"
const RPAREN = "}"

func Tokenize(input string) []string {
	var tokenArray []string
	position := 0
	midJson := true

	for midJson {
		if position == (len(input) - 1) {
			break
		}
		char := input[position]
		switch char {
		case '{':
			tokenArray = append(tokenArray, LPAREN)
			position += 1
		case '}':
			tokenArray = append(tokenArray, RPAREN)
			position += 1
		default:
			if isLetter(char) {
				currentPosition := position
				for isLetter(input[position]) {
					position += 1
				}
				tokenArray = append(tokenArray, input[currentPosition:position])
			}

		}
	}
	return tokenArray

}

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z'
}

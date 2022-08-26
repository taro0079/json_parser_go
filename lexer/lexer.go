package lexer

const LPAREN = "{"
const RPAREN = "}"
const LBRAKET = "["
const RBRAKET = "]"
const COLON = ":"
const COMMA = ","

func Tokenize(input string) []string {
	var tokenArray []string
	position := 0
	midJson := true

	for midJson {
		char := input[position]
		switch char {
		case '{':
			tokenArray = append(tokenArray, LPAREN)
			position += 1
		case '}':
			tokenArray = append(tokenArray, RPAREN)
			position += 1
		case '[':
			tokenArray = append(tokenArray, LBRAKET)
			position += 1
		case ']':
			tokenArray = append(tokenArray, RBRAKET)
			position += 1
		case ':':
			tokenArray = append(tokenArray, COLON)
			position += 1
		case ',':
			tokenArray = append(tokenArray, COMMA)
			position += 1

		default:
			if isLetter(char) {
				currentPosition := position
				readPosition := position + 1
				for isLetter(input[position]) {
					if readPosition > len(input) {
						readPosition = 0
					}
					if isLetter(input[readPosition]) {
						position += 1
						readPosition += 1

					} else {
						position += 1
						tokenArray = append(tokenArray, input[currentPosition:position])
						break
					}

				}
			}

		}
		if position >= len(input) {
			midJson = false
		}
	}
	return tokenArray

}

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z'
}

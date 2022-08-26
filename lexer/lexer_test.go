package lexer

import (
	"testing"
)

func TestTokenizer(t *testing.T) {
	testInput := "{test}"
	expectedArray := []string{"{", "test", "}"}
	got := Tokenize(testInput)
	for i := 0; i < len(got); i++ {
		if got[i] != expectedArray[i] {
			t.Errorf("wrong")
		}
	}
}

package lexer

import (
	"testing"
)

func TestTokenizer(t *testing.T) {
	testInput := "{test:test1}"
	expectedArray := []string{"{", "test", ":", "test1", "}"}
	got := Tokenize(testInput)
	for i := 0; i < len(got); i++ {
		if got[i] != expectedArray[i] {
			t.Errorf("wrong")
		}
	}
}

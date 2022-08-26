package main

import (
	"fmt"
	"json_parser/lexer"
)

func main() {
	l := lexer.Tokenize("{test}")
	fmt.Println(l)

}

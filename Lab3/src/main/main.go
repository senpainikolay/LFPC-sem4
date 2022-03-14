package main

import (
	"fmt"
	"lexer"
	"os"
	"token"
)

func main() {
	file_input, _  := os.ReadFile("code.txt") 
	lexer_input := string(file_input)
	l := lexer.New(lexer_input)

	for {
		tok := l.NextToken()
		fmt.Println(tok) 
		if tok.Type == token.EOF {
			break
		}
	}
}

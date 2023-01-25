package main

import (
	"bible-compiler/lexer"
)

func main() {
	// get arg after "go run main.go"
	code := "int aa = 12"
	l := lexer.New(code)
	for {
		token := l.NextToken()
		println(token.Type, " ", token.Literal)
		if l.Ch == 0 {
			break
		}
	}
}

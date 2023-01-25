package lexer

import (
	"bible-compiler/tokenizer"
)

type Lexer struct {
	Source       string //Source code
	Position     int    //Current Cursor Position
	ReadPosition int    //Next Reading Position
	Ch           byte   //Current Character
}

// Declare new lexer
func New(Source string) *Lexer {
	return &Lexer{Source: Source}
}

// Consume char and to to next
func (lexer *Lexer) ReadChar() {
	if len(lexer.Source) > lexer.ReadPosition {
		lexer.Ch = lexer.Source[lexer.ReadPosition]
	} else {
		lexer.Ch = 0
	}
	lexer.Position = lexer.ReadPosition
	lexer.ReadPosition++
}

// Look at next char
func (lexer *Lexer) PeekChar() byte {
	if len(lexer.Source) > lexer.ReadPosition {
		return lexer.Source[lexer.ReadPosition]
	}
	return 0
}

// Assign token once found
func (lexer *Lexer) NextToken() tokenizer.Token {
	var token tokenizer.Token
	lexer.ReadChar()
	lexer.eatWhiteSpace()
	switch lexer.Ch {
	// basic tokens
	case '(':
		token = tokenizer.Token{Type: tokenizer.PARENTL, Literal: string(lexer.Ch)}
	case ')':
		token = tokenizer.Token{Type: tokenizer.PARENTR, Literal: string(lexer.Ch)}
	case '{':
		token = tokenizer.Token{Type: tokenizer.CURLYL, Literal: string(lexer.Ch)}
	case '}':
		token = tokenizer.Token{Type: tokenizer.CURLYR, Literal: string(lexer.Ch)}
	case '[':
		token = tokenizer.Token{Type: tokenizer.BRACKETL, Literal: string(lexer.Ch)}
	case ']':
		token = tokenizer.Token{Type: tokenizer.BRACKETR, Literal: string(lexer.Ch)}
	case ',':
		token = tokenizer.Token{Type: tokenizer.COMMA, Literal: string(lexer.Ch)}
	case '+':
		token = tokenizer.Token{Type: tokenizer.PLUS, Literal: string(lexer.Ch)}

	// number tokens
	case '-':
		token = tokenizer.Token{Type: tokenizer.MINUS, Literal: string(lexer.Ch)}
	case '/':
		token = tokenizer.Token{Type: tokenizer.DIVIDE, Literal: string(lexer.Ch)}
	case '%':
		token = tokenizer.Token{Type: tokenizer.MODULO, Literal: string(lexer.Ch)}

	default:

		if lexer.Ch == '!' {
			if lexer.PeekChar() == '=' {
				lexer.ReadChar()
				token = tokenizer.Token{Type: tokenizer.NEQ, Literal: "!="}
			} else {
				token = tokenizer.Token{Type: tokenizer.NOT, Literal: "!"}
			}
		} else if lexer.Ch == '&' {
			if lexer.PeekChar() == '&' {
				lexer.ReadChar()
				token = tokenizer.Token{Type: tokenizer.AND, Literal: "&&"}
			} else {
				token = tokenizer.Token{Type: tokenizer.ADDR, Literal: ""}
			}

		} else if lexer.Ch == '|' && lexer.PeekChar() == '|' {
			lexer.ReadChar()
			token = tokenizer.Token{Type: tokenizer.OR, Literal: "||"}

		} else if lexer.Ch == '=' {
			if lexer.PeekChar() == '=' {
				lexer.ReadChar()
				token = tokenizer.Token{Type: tokenizer.EQ, Literal: "=="}
			} else {
				token = tokenizer.Token{Type: tokenizer.ASSIGN, Literal: "="}
			}
		} else if lexer.Ch == '>' {
			if lexer.PeekChar() == '=' {
				lexer.ReadChar()
				token = tokenizer.Token{Type: tokenizer.MOREEQ, Literal: ">="}
			} else {
				token = tokenizer.Token{Type: tokenizer.MORETHAN, Literal: string(lexer.Ch)}
			}

		} else if lexer.Ch == '<' {
			if lexer.PeekChar() == '=' {
				lexer.ReadChar()
				token = tokenizer.Token{Type: tokenizer.LESSEQ, Literal: "<="}
			} else {
				token = tokenizer.Token{Type: tokenizer.LESSTHAN, Literal: string(lexer.Ch)}
			}

		} else if lexer.Ch == '*' {
			if lexer.PeekChar() == '*' {
				lexer.ReadChar()
				token = tokenizer.Token{Type: tokenizer.POWER, Literal: "**"}
			} else {
				token = tokenizer.Token{Type: tokenizer.MULTIPLY, Literal: string(lexer.Ch)}
			}
		} else if isValidNumber(lexer.Ch) {
			nbr := lexer.ReadNumber()
			var nbr_type tokenizer.TokenType = "number"
			token = tokenizer.Token{Type: nbr_type, Literal: nbr}
		} else if isValidLetter(lexer.Ch) { //int[ ]a Token type int lit int, Toe
			id := lexer.ReadIdentifier()
			id_type := evaluateIDType(id)
			token = tokenizer.Token{Type: id_type, Literal: id}
		} else {
			token = tokenizer.Token{Type: tokenizer.ERROR, Literal: "ERROR"}
		}
	}
	return token
}

func (lexer *Lexer) eatWhiteSpace() {
	for lexer.Ch == ' ' || lexer.Ch == '\t' || lexer.Ch == '\n' {
		lexer.ReadChar()
	}
}

func evaluateIDType(id string) tokenizer.TokenType {
	switch id {
	case tokenizer.INT:
		return tokenizer.INT
	case tokenizer.FLOAT:
		return tokenizer.FLOAT
	case tokenizer.BOOL:
		return tokenizer.BOOL
	case tokenizer.CHAR:
		return tokenizer.CHAR
	case tokenizer.STRING:
		return tokenizer.STRING
	case tokenizer.LONG:
		return tokenizer.LONG
	case tokenizer.DOUBLE:
		return tokenizer.DOUBLE
	default:
		return tokenizer.ID
	}
}

func (lexer *Lexer) ReadNumber() string {
	startPos := lexer.Position
	for isValidNumber(lexer.Ch) {
		lexer.ReadChar()
	}
	return lexer.Source[startPos:lexer.Position]
}

func (lexer *Lexer) ReadIdentifier() string {
	startPos := lexer.Position
	for isValidLetter(lexer.Ch) || isValidNumber(lexer.Ch) {
		lexer.ReadChar()
	}
	return lexer.Source[startPos:lexer.Position]
}

func isValidLetter(c byte) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || c == '_' || c == '-'
}

func isValidNumber(c byte) bool {
	return '0' <= c && c <= '9'
}

// a45a4ze
// 564zea

//int a = 12
//  ^
//Source : "int a = 12"
//postion : 2
//ReadPosition : 3
//Ch : t

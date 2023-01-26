package tokenizer

type TokenType string

// Structure that holds a token
type Token struct {
	Type    TokenType
	Literal string
}

// Declare all types of tokens
const (
	//AST specific Tokens
	EOF   = "EOF"
	ERROR = "ERROR"

	//Statement Tokens
	ASSIGN   = "="
	COMMA    = ","
	PARENTL  = "("
	PARENTR  = ")"
	BRACKETL = "["
	BRACKETR = "]"
	CURLYL   = "{"
	CURLYR   = "}"

	//Arithmetic Operators
	PLUS     = "+"
	MINUS    = "-"
	MULTIPLY = "*"
	DIVIDE   = "/"
	MODULO   = "%"
	POWER    = "**"
	MORETHAN = ">"
	LESSTHAN = "<"
	MOREEQ   = ">="
	LESSEQ   = "<="

	//Boolean Operators
	AND  = "&&"
	OR   = "||"
	NOT  = "!"
	NAND = "!&"
	NOR  = "!|"
	XOR  = "^"
	EQ   = "=="
	NEQ  = "!="

	//Typedef
	INT    = "int"
	FLOAT  = "float"
	BOOL   = "bool"
	CHAR   = "char"
	STRING = "string"
	LONG   = "long"
	DOUBLE = "double"

	// Identifier
	ID = "id"

	// PTR
	PTR  = "#"
	ADDR = "&"
)

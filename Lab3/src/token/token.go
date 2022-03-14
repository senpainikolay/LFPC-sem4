package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENTIFIER"
	INT   = "INT"

	// Operators
	ASSIGN                  = "="
	PLUS                    = "+"
	MINUS                   = "-"
	DIVISION          =""
	MUTIPLICATION = "*"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="

	// Delimiters
	COMMA

	LPAREN = "LEFT_PAREN"
	RPAREN = "RIGHT_PAREN"
	LBRACE = "LEFT_BRACE"
	RBRACE = "RIGHT_BRACE"

	// Keywords
	FUNCTION = "FUNCTION_DECLARION"
	VAR      = "VAR_DECLARATION"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"def":  FUNCTION,
	"var":    VAR,
	"True":   TRUE,
	"False":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

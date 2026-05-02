package token

type TokenType int

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

const (
	IDENTIFIER TokenType = iota // 0
	INT                         // 1

	// Keywords
	LET // 2

	// Operators
	ASSIGN // 3
	PLUS   // 4

	// Delimiters
	COMMA     // 5
	SEMICOLON // 6
	LPAREN    // 7
	RPAREN    // 8
	LBRACE    // 9
	RBRACE    // 10
)

var Keywords = map[string]TokenType{
	"let": LET,
}

func isKeyword(str string) TokenType {
	if keyword, ok := Keywords[str]; ok {
		return keyword
	} else {
		return IDENTIFIER
	}
}

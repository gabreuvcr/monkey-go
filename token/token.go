package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	Illegal    TokenType = "Illegal"
	Eof        TokenType = "Eof"
	Ident      TokenType = "Ident"
	Int        TokenType = "Int"
	Assign     TokenType = "Assign"
	Plus       TokenType = "Plus"
	Minus      TokenType = "Minus"
	Bang       TokenType = "Bang"
	Asterisk   TokenType = "Asterisk"
	Slash      TokenType = "Slash"
	LessThan   TokenType = "LessThan"
	GreatThan  TokenType = "GreatThan"
	Equal      TokenType = "Equal"
	NotEqual   TokenType = "NotEqual"
	Comma      TokenType = "Comma"
	Semicolon  TokenType = "Semicolon"
	LeftParen  TokenType = "LeftParen"
	RightParen TokenType = "RightParen"
	LeftBrace  TokenType = "leftBrace"
	RightBrace TokenType = "RightBrace"
	
	Function   TokenType = "Function"
	Let        TokenType = "Let"
	True       TokenType = "True"
	False      TokenType = "False"
	If         TokenType = "If"
	Else       TokenType = "Else"
	Return     TokenType = "Return"
)

var keywords = map[string]TokenType{
	"fn":     Function,
	"let":    Let,
	"true":   True,
	"false":  False,
	"if":     If,
	"else":   Else,
	"return": Return,
}

func New(tokenType TokenType, literal string) Token {
	return Token{Type: tokenType, Literal: literal}
}

func Keyword(ident string) (TokenType, bool) {
	tok, ok := keywords[ident]
	return tok, ok
}

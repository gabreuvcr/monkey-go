package token

type Token struct {
	Type    TokenType
	Literal string
}

type TokenType string

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

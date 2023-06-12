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
	Assign     TokenType = "="
	Plus       TokenType = "+"
	Minus      TokenType = "-"
	Bang       TokenType = "!"
	Asterisk   TokenType = "*"
	Slash      TokenType = "/"
	LessThan   TokenType = "<"
	GreatThan  TokenType = ">"
	Equal      TokenType = "=="
	NotEqual   TokenType = "!="
	Comma      TokenType = ","
	Semicolon  TokenType = ";"
	LeftParen  TokenType = "("
	RightParen TokenType = ")"
	LeftBrace  TokenType = "{"
	RightBrace TokenType = "}"
	
	Function   TokenType = "Function"
	Let        TokenType = "Let"
	True       TokenType = "True"
	False      TokenType = "False"
	If         TokenType = "If"
	Else       TokenType = "Else"
	Return     TokenType = "Return"
)

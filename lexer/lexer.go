package lexer

import "monkey/token"

type Lexer struct {
	input        string
	currPosition int
	nextPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '+':
		tok = token.New(token.Plus, "+")
	case '-':
		tok = token.New(token.Minus, "-")
	case '=':
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.New(token.Equal, "==")
		} else {
			tok = token.New(token.Assign, "=")
		}
	case '!':
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.New(token.NotEqual, "!=")
		} else {
			tok = token.New(token.Bang, "!")
		}
	case '/':
		tok = token.New(token.Slash, "/")
	case '*':
		tok = token.New(token.Asterisk, "*")
	case '<':
		tok = token.New(token.LessThan, "<")
	case '>':
		tok = token.New(token.GreatThan, ">")
	case ';':
		tok = token.New(token.Semicolon, ";")
	case '(':
		tok = token.New(token.LeftParen, "(")
	case ')':
		tok = token.New(token.RightParen, ")")
	case ',':
		tok = token.New(token.Comma, ",")
	case '{':
		tok = token.New(token.LeftBrace, "{")
	case '}':
		tok = token.New(token.RightBrace, "}")
	case 0:
		tok = token.New(token.Eof, "")
	default:
		if isLetter(l.ch) {
			var literal = l.readIdentifier()
			if keywordToken, ok := token.Keyword(literal); ok {
				tok = token.New(keywordToken, literal)
			} else {
				tok = token.New(token.Ident, literal)
			}
			return tok
		} else if isDigit(l.ch) {
			var literal = l.readInt()
			return token.New(token.Int, literal)
		} else {
			tok = token.New(token.Illegal, string(l.ch))
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readChar() {
	if l.nextPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.nextPosition]
	}
	l.currPosition = l.nextPosition
	l.nextPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.nextPosition >= len(l.input) {
		return 0
	}
	return l.input[l.nextPosition]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readIdentifier() string {
	currPosition := l.currPosition
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[currPosition:l.currPosition]
}

func (l *Lexer) readInt() string {
	currPosition := l.currPosition
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[currPosition:l.currPosition]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

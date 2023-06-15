package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
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
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			var literal = string(ch) + string(l.ch)
			tok = createToken(token.Equal, literal)
		} else {
			tok = createToken(token.Assign, string(l.ch))
		}
	case '+':
		tok = createToken(token.Plus, string(l.ch))
	case '-':
		tok = createToken(token.Minus, string(l.ch))
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			var literal = string(ch) + string(l.ch)
			tok = createToken(token.NotEqual, literal)
		} else {
			tok = createToken(token.Bang, string(l.ch))
		}
	case '/':
		tok = createToken(token.Slash, string(l.ch))
	case '*':
		tok = createToken(token.Asterisk, string(l.ch))
	case '<':
		tok = createToken(token.LessThan, string(l.ch))
	case '>':
		tok = createToken(token.GreatThan, string(l.ch))
	case ';':
		tok = createToken(token.Semicolon, string(l.ch))
	case '(':
		tok = createToken(token.LeftParen, string(l.ch))
	case ')':
		tok = createToken(token.RightParen, string(l.ch))
	case ',':
		tok = createToken(token.Comma, string(l.ch))
	case '{':
		tok = createToken(token.LeftBrace, string(l.ch))
	case '}':
		tok = createToken(token.RightBrace, string(l.ch))
	case 0:
		tok = createToken(token.Eof, "")
	default:
		if isLetter(l.ch) {
			var literal = l.readIdentifier()
			switch literal {
			case "fn":
				tok = createToken(token.Function, literal)
			case "let":
				tok = createToken(token.Let, literal)
			case "true":
				tok = createToken(token.True, literal)
			case "false":
				tok = createToken(token.False, literal)
			case "if":
				tok = createToken(token.If, literal)
			case "else":
				tok = createToken(token.Else, literal)
			case "return":
				tok = createToken(token.Return, literal)
			default:
				tok = createToken(token.Ident, literal)
			}
			return tok
		} else if isDigit(l.ch) {
			var literal = l.readInt()
			return createToken(token.Int, literal)
		} else {
			tok = createToken(token.Illegal, string(l.ch))
		}
	}

	l.readChar()
	return tok
}

func createToken(tokenType token.TokenType, ch string) token.Token {
	return token.Token{Type: tokenType, Literal: ch}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readInt() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

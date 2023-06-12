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
			tok = createTokenLiteral(token.Equal, literal)
		} else {
			tok = createToken(token.Assign, l.ch)
		}
	case '+':
		tok = createToken(token.Plus, l.ch)
	case '-':
		tok = createToken(token.Minus, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			var literal = string(ch) + string(l.ch)
			tok = createTokenLiteral(token.NotEqual, literal)
		} else {
			tok = createToken(token.Bang, l.ch)
		}
	case '/':
		tok = createToken(token.Slash, l.ch)
	case '*':
		tok = createToken(token.Asterisk, l.ch)
	case '<':
		tok = createToken(token.LessThan, l.ch)
	case '>':
		tok = createToken(token.GreatThan, l.ch)
	case ';':
		tok = createToken(token.Semicolon, l.ch)
	case '(':
		tok = createToken(token.LeftParen, l.ch)
	case ')':
		tok = createToken(token.RightParen, l.ch)
	case ',':
		tok = createToken(token.Comma, l.ch)
	case '{':
		tok = createToken(token.LeftBrace, l.ch)
	case '}':
		tok = createToken(token.RightBrace, l.ch)
	case 0:
		tok = createTokenLiteral(token.Eof, "")
	default:
		if isLetter(l.ch) {
			var literal = l.readIdentifier()
			switch literal {
			case "fn":
				tok = createTokenLiteral(token.Function, literal)
			case "let":
				tok = createTokenLiteral(token.Let, literal)
			case "true":
				tok = createTokenLiteral(token.True, literal)
			case "false":
				tok = createTokenLiteral(token.False, literal)
			case "if":
				tok = createTokenLiteral(token.If, literal)
			case "else":
				tok = createTokenLiteral(token.Else, literal)
			case "return":
				tok = createTokenLiteral(token.Return, literal)
			default:
				tok = createTokenLiteral(token.Ident, literal)
			}
			return tok
		} else if isDigit(l.ch) {
			var literal = l.readInt()
			return createTokenLiteral(token.Int, literal)
		} else {
			tok = createToken(token.Illegal, l.ch)
		}
	}

	l.readChar()
	return tok
}

func createToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func createTokenLiteral(tokenType token.TokenType, ch string) token.Token {
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

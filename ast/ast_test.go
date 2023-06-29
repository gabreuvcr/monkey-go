package ast

import (
	"monkey/token"
	"testing"
)

func TestString(t *testing.T) {
	input := "let myVar = anotherVar;"
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.Let, Literal: "let"},
				Name: &IdentifierExpression{
					Token: token.Token{Type: token.Ident, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &IdentifierExpression{
					Token: token.Token{Type: token.Ident, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != input {
		t.Errorf("program.String() wrong, expected=%q, got=%q", input, program.String())
	}
}

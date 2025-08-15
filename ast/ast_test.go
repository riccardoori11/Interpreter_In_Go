package ast

import (

	"interpreter_go/token/token"
	"testing"

)


func TestString(t *testing.T){

	program := &Program{

		Statements: []Statement{

			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: 	token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myvar",
				
				},


				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},

					Value: "anotherVar",
					},

				},
			},
	}

		if program.String() != "let myVar = anotherVar"{

			t.Errorf("Wrong. We got %q", program.String())

		}

	}

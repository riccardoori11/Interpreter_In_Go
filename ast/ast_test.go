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
				Value: "myVar",


			},

			Value: &Identifier{

				Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
				Value: "anotherVar"
			}

		},


	}

}

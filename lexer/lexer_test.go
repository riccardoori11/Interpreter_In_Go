package lexer

import (
	
	"interpreter_go/token/token"
	"testing"
)



func TestNextToken(t *testing.T){

		input := ` 5 + 3 nOT= 7 
				
			fn add(x,y){
				
				return x + y;
			}

			let five = /5/;
			<>
		`

		tests := []struct{
		expectedType token.TokenType
		expectedLiteral string
		}{
			{token.INT, "5"},
			{token.PLUS, "+"},
			{token.INT, "3"},
			{token.NEQ, "nOT="},
			{token.INT, "7"},
			{token.FUNCTION, "fn"},
			{token.IDENT, "add"},
			{token.LPAR, "("},
			{token.IDENT, "x"},
			{token.COMMA, ","},
			{token.IDENT, "y"},
			{token.RPAR, ")"},
			{token.LBRAC, "{"},
			{token.RETURN, "return"},
			{token.IDENT, "x"},
			{token.PLUS, "+"},
			{token.IDENT, "y"},
			{token.SEMICOLON, ";"},
			{token.RBRAC, "}"},
			{token.LET, "let"},
			{token.IDENT, "five"},
			{token.ASSIGN, "="},
			{token.SLASH, "/"},
			{token.INT, "5"},
			{token.SLASH, "/"},
			{token.SEMICOLON, ";"},
			{token.LTHAN, "<"},
			{token.GTHAN, ">"},
			{token.EOF, ""},


						}

	l := New(input)

	for i,tt := range tests{


		tok := l.NextToken()
		
		if tok.Type != tt.expectedType{
			
			t.Fatalf("tests[%d] - tokentype wrong. expected = %q, got = %q", i, tt.expectedType, tok.Type )

		}
		if tok.Literal != tt.expectedLiteral{
				
				t.Fatalf("tests[%d] - tokenliteral wrong. expected = %q, got = %q ", i, tt.expectedLiteral, tok.Literal)

		}
	}}







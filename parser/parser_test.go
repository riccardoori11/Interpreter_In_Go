package parser

import (
	"interpreter_go/token/ast"
	"interpreter_go/token/lexer"
	"testing"

)
func testLetStatement(t *testing.T, s ast.Statement, name string)bool{


		if s.TokenLiteral() != "let"{
			
			t.Errorf("s.TokenLiteral not 'let' instead we got %q", s.TokenLiteral())
			return false
		}

		letStmt, ok := s.(*ast.LetStatement)
		if !ok{

			t.Errorf("s not *ast.LetStatement. instead we had %T", s)
			return false
		}
		if letStmt.Name.Value != name{
			t.Errorf("letStmt.Name.Value not '%s'. Instead we got %s", name, letStmt.Name.Value)
			return false	
		}
		if letStmt.Name.TokenLiteral() != name{
			t.Errorf("s.Name not %s, got %s", name, letStmt.Name)
			return false	
		}

		return true
}

func checkParserErrors(t *testing.T, p *Parser){

	
	errors := p.Errors()

	if len(errors) == 0{
		return 
	}

	t.Errorf("parser has %d errors", len(errors) )

	for _,msg := range errors{

		t.Errorf("parser error: %q", msg)

	}

	t.FailNow()

}





func TestReturnStatements(t *testing.T){


	
		input := `

		return 5;
		return 100;
		return 1230321;

		`

		l := lexer.New(input)
		p := New(l)


		program := p.ParserProgram()

		checkParserErrors(t,p)

		Number_Of_Statements_Return := 3;
			
		if len(program.Statements) != Number_Of_Statements_Return{

		t.Fatalf("Expected length %d instead we got %d", Number_Of_Statements_Return ,len(program.Statements))	
		}

		for _,stmt := range program.Statements{

			returnstmt, ok := stmt.(*ast.ReturnStatement)
			if !ok {

				t.Errorf("smt not *ast.Statement. got = %T", stmt)
			}

			if returnstmt.TokenLiteral() != "return"{
				
				t.Errorf("return stmtTokenLiteral not return instead we got %q", returnstmt.TokenLiteral())

			} 

		}


}

func TestLetStatement(t *testing.T){


	input := `
		
	let x = 5;
	let y = 10;
	let foobar = 838383;
	`


	l := lexer.New(input)
	p := New(l)


	program := p.ParserProgram()
	checkParserErrors(t,p)


	if len(program.Statements) != 3{
		t.Fatalf("The program does not contains 3 statemetns instead had %d", len(program.Statements))	
	}

	tests := []struct{
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i,tt := range tests{
		stmt := program.Statements[i]
		if !testLetStatement(t,stmt, tt.expectedIdentifier){
			return
		}
	}

}



func TestIdentifierExpression(t *testing.T){
		
	input := "foobar;"
		
	l := lexer.New(input)

	p := New(l)

	program := p.ParserProgram()
	checkParserErrors(t,p)

	
	if len(program.Statements) != 1{

		t.Fatalf("Expected 1 statemenet, instead we got %d ", len(program.Statements))
	}
	
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok{

		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got %T", program.Statements[0])

	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	
	if !ok{


		t.Fatalf("stmt.Expression is not an identifier. Instead it is %q", stmt.Expression)

	}
	
	
	if ident.Value != "foobar"{

		t.Errorf("ident.Value is not foobar. Instead we got %s", ident.Value)
	}
	
	if ident.TokenLiteral() != "foobar"{

		t.Errorf("ident.TokenLiteral is not foobar, we got %s", ident.TokenLiteral())
	}



}



func TestIntegerLiteralExpression(t *testing.T){
		
	input := "5;"
	l := lexer.New(input)
	
	p := New(l)
	program := p.ParserProgram()
	checkParserErrors(t,p)
	
	if len(program.Statements) != 1{
		
		t.Fatalf("Expected 1 statement, got %d", len(program.Statements))

	}
	
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok{

		t.Fatalf("Statement was not an ast.ExpressionStatement %T", program.Statements[0] )
	}

	literal,ok := stmt.Expression.(*ast.IntegerLiteral)	
	
	if !ok{
		
		t.Fatalf("stmt.Expression is not an integer literal %T", stmt.Expression)

	}
	
	if literal.Value != 5{
		t.Errorf("Wrong input got %d", literal.Value)
	}
	
	if literal.Token.Literal != "5"{

		t.Errorf("literal.Token.Literal got %s", literal.Token.Literal)
	}
	

}

















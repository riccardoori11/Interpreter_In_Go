package parser

import (
	"fmt"
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

		t.Errorf("ident.TokenLiteral is not %s, we got %s","foobar", ident.TokenLiteral())
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

		t.Errorf("literal.Token.Literal expected %s instead we got %s", "5", literal.Token.Literal)
	}
	

}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64)bool{

	integ, ok := il.(*ast.IntegerLiteral)	
	
	if !ok{

		t.Errorf("Integ is not of type ast.IntegerLiteral. Instead got %s", il)
		return false
	}
	
	if integ.Value != value{

		t.Errorf("The value is wrong. Expected %d and got %d", value,integ.Value)
		return false
	}
	
	if integ.TokenLiteral() != fmt.Sprintf("%d", value){

		t.Errorf("Expected %d instead we got %s", value, integ.TokenLiteral())
		return false
	}
	return true

}


func TestParsingPrefix(t *testing.T){

	
	prefixTests := []struct{
		input 	string
		operator 	string
		integerValue int64

	}{
		{"/5", "/", 5},
		{"-15", "-", 15},

	}
	for _,tt := range prefixTests{

		l := lexer.New(tt.input)
		p := New(l)
		
		program := p.ParserProgram()
		
		checkParserErrors(t,p)
		
		if len(program.Statements) != 1{

			t.Fatalf("Expected 1 statement. Instead we got %d", len(program.Statements))	
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok{
				
			t.Fatalf("Expeced a *ast.ExpressionStatemenet instead got %T", program.Statements[0])

		}
		
		exp,ok := stmt.Expression.(*ast.PrefixExpression)
		
		if !ok{
			
			t.Fatalf("Error not a PrefixExpression. Got %T", stmt.Expression)

		}
		
		if exp.Operator != tt.operator{

			t.Fatalf("exp.Operator is not %s, instead we got %s", tt.operator, exp.Operator)
		}
		
		if !testIntegerLiteral(t, exp.Right, tt.integerValue){
			
			return

		}
		

	}

}

func TestInfixParseExpression(t *testing.T){

	
	infix_tests := []struct{

		input string
		left_value int64
		operator string
		right_value int64


	}{
		{"5 + 5", 5, "+",5 },
{"5 - 5", 5, "-",5 },
{"5 nOT= 5", 5, "nOT=",5 },
{"5 / 5", 5, "/",5 },
{"5 > 5", 5, ">",5 },
{"5 < 5", 5, "<",5 },
}

for _,tt := range infix_tests{

	l := lexer.New(tt.input)
	p := New(l)
	
	program := p.ParserProgram()
	checkParserErrors(t,p)
	
	if len(program.Statements) != 1{

		t.Fatalf("Expected 1 statement instead we got %d", len(program.Statements))
	}
	
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	
	if !ok{
		
		t.Fatalf("stmt is not a *ast.ExpressionStatement. Instead we got %T", program.Statements[0])
	}
	
	exp, ok := stmt.Expression.(*ast.InfixExpression)

}




}






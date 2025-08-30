package parser

import (
	"fmt"
	"interpreter_go/token/ast"
	"interpreter_go/token/lexer"
	"interpreter_go/token/token"
	"strconv"
)


type Parser struct{



	l *lexer.Lexer
errors []string

	currToken token.Token
	peekToken token.Token

	prefixParseFns	map[token.TokenType]prefixParseFn
	infixParseFns map[token.TokenType]infixParseFn

}

const (

	_ int = iota 	
	
	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL



)

var precedence  = map[token.TokenType]int{


	token.ASSIGN: EQUALS,
token.NEQ: EQUALS,
token.LTHAN: LESSGREATER,
token.GTHAN: LESSGREATER,
token.PLUS: SUM,
token.MINUS: SUM,
token.ASTERIK: PRODUCT,

token.SLASH: PRODUCT,






} 




func (p *Parser)nextToken(){
		

	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()

}

func (p *Parser)parsePrefixExpression() ast.Expression{
		
	expression := &ast.PrefixExpression{
			
			Token: p.currToken,
			Operator: p.currToken.Literal,


	}
	
	p.nextToken()
	expression.Right = p.parseExpression(PREFIX)
	
	return expression


}

func (p *Parser) parseBoolean() ast.Expression{
	
	return &ast.Boolean{Token: p.currToken,Value: p.currTokenIs(token.TRUE)}

}


func New(l *lexer.Lexer) *Parser{

		
		p := &Parser{
			l : l,
			errors: []string{},
	}
		
	
		
		p.registerPrefix(token.TRUE,p.parseBoolean)
p.registerPrefix(token.FALSE,p.parseBoolean)
		p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
		p.registerPrefix(token.SLASH, p.parsePrefixExpression)	
		p.registerPrefix(token.MINUS, p.parsePrefixExpression)
		p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.ParseIntegerLiteral)

	p.infixParseFns = make(map[token.TokenType]infixParseFn)

		p.registerInfix(token.PLUS, p.parseInfixExpression)
p.registerInfix(token.MINUS, p.parseInfixExpression)
p.registerInfix(token.SLASH, p.parseInfixExpression)
p.registerInfix(token.ASTERIK, p.parseInfixExpression)
p.registerInfix(token.ASSIGN, p.parseInfixExpression)
p.registerInfix(token.NEQ, p.parseInfixExpression)
p.registerInfix(token.LTHAN, p.parseInfixExpression)
p.registerInfix(token.GTHAN, p.parseInfixExpression)



		p.nextToken()
		p.nextToken()


		return p
}

func (p *Parser) Errors() []string{
	
	return p.errors

}

func (p *Parser)peekError(t token.TokenType){


	msg := fmt.Sprintf("Expected met token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)

}


func (p *Parser)parseIdentifier() ast.Expression{

	return &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}

}


func (p *Parser) ParserProgram() *ast.Program{

	program := &ast.Program{}
	program.Statements = []ast.Statement{}


	for p.currToken.Type != token.EOF{


		stmt := p.parseStatement()
		if stmt != nil{
			program.Statements = append(program.Statements, stmt) 

		}

		p.nextToken()
	}
	
	return program

}

func (p *Parser) currTokenIs(t token.TokenType)bool{
		
	return p.currToken.Type == t

}
func (p *Parser) peekTokenIs(t token.TokenType)bool{
		
	return p.peekToken.Type == t

}

func (p *Parser) expectPeek(t token.TokenType)bool{
		
	if p.peekTokenIs(t){
		
		p.nextToken()
		return true
	}else{
		p.peekError(t)
		return false
	}

}

func (p *Parser) parseLetStatement() *ast.LetStatement{

	stmt := &ast.LetStatement{Token: p.currToken}

	if !p.expectPeek(token.IDENT){
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}

	if !p.expectPeek(token.ASSIGN){
		return nil
	}



	//TODO WE ARE SKIPPING THE EXPRESSIONS UNTIL WE ENCOUNTER A SEMI COLON

	for !p.currTokenIs(token.SEMICOLON){
		p.nextToken()
	}

	return stmt
}


func (p *Parser) parserReturnStatement() *ast.ReturnStatement{

		stmt := &ast.ReturnStatement{Token: p.currToken}

		p.nextToken()

		for !p.currTokenIs(token.SEMICOLON){
			p.nextToken()
		}

		return stmt

}



func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement{

		stmt := &ast.ExpressionStatement{Token: p.currToken}
		
		stmt.Expression = p.parseExpression(LOWEST)
	
		if p.peekTokenIs(token.SEMICOLON){
			
			p.nextToken()

		}

		return stmt
}

func (p *Parser) parseStatement() ast.Statement{

	switch p.currToken.Type{


	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parserReturnStatement()
	


	default:
		return p.parseExpressionStatement()

	}


}


type (

	prefixParseFn func() ast.Expression
	infixParseFn func(ast.Expression) ast.Expression

)




func (p *Parser)registerPrefix(tokenType token.TokenType, fn prefixParseFn){


	p.prefixParseFns[tokenType] = fn


}



func (p *Parser)registerInfix(tokenType token.TokenType, fn infixParseFn){

	p.infixParseFns[tokenType] = fn


}


func (p *Parser)ParseIntegerLiteral() ast.Expression{

	lit := &ast.IntegerLiteral{Token: p.currToken}
	
	value,err := strconv.ParseInt(p.currToken.Literal,0,64)
	
	if err != nil{


		msg := fmt.Sprintf("could not parse %q as an integer", p.currToken.Literal)
	
		p.errors = append(p.errors, msg)
		
		return nil
	}
	
	lit.Value = value
	
	return lit
	

}


func (p *Parser)noPrefixParseFnError(t token.TokenType){

		msg := fmt.Sprintf("no prefix parse function for %s found", t)
		
		p.errors = append(p.errors, msg)
}


func (p *Parser)parseExpression(precedence int) ast.Expression{

 
	prefix := p.prefixParseFns[p.currToken.Type]
	if prefix == nil{
		
		p.noPrefixParseFnError(p.currToken.Type)
		return nil

	}
	
	leftExp := prefix()	
	
	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence(){

		infix := p.infixParseFns[p.peekToken.Type]
		
		if infix == nil{
			
			return leftExp

		}
		
		p.nextToken()

		leftExp = infix(leftExp)
	} 	

	return leftExp


}

func (p *Parser)peekPrecedence()int{
	
	if p,ok := precedence[p.peekToken.Type]; ok{
		return p
	}
	
	return LOWEST


}

func (p *Parser) curPrecedence()int{
	
	if p,ok := precedence[p.currToken.Type]; ok{
		return p
	}
	
	return LOWEST
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression{

	expression := &ast.InfixExpression{

		Token: 	p.currToken,
		Operator: p.currToken.Literal,
		Left: left,

	}
	
	precedence := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)
	
	return expression


}




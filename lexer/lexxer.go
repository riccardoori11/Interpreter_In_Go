package lexer

import (
	"interpreter_go/token/token"
)

type Lexer struct {

	input string
	position int
	readPosition int
	ch byte
}



// position gives us the where we last read the input and readPosition gives us the next input we are going to read
// this only works on ASCII elements, if we want full UTF-8 it would be different 
func (l *Lexer)readChar(){

		if l.readPosition >= len(l.input){
			l.ch = 0
		} else{
		l.ch = l.input[l.readPosition]

		}
	l.position = l.readPosition
	l.readPosition += 1

}

func (l *Lexer)skipWhiteSpace(){
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' || l.ch == '!'{
		l.readChar()
	}
}



func isLetter(ch byte)bool{

	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'  

}

func isDigit(ch byte)bool{
		
	return '0' <= ch && ch <= '9'

}

func (l *Lexer)readIdentifier() string{

	position := l.position

	for isLetter(l.ch){

		l.readChar()
	}
	return l.input[position:l.position]
}

func New(input string) *Lexer{
		
	l := &Lexer{input: input}
	l.readChar()
	return l

}

func NewToken(tokenType token.TokenType, ch byte ) token.Token{
		
	return token.Token{Type: tokenType, Literal: string(ch) }
}


func (l *Lexer)PeekWord(length_peek int) string{
	
	end := l.position + length_peek


	if end > len(l.input){
		
		end = len(l.input)

	}

	return l.input[l.position:end]


	}








func (l *Lexer)readNumber()string{

		
	position := l.position
	for isDigit(l.ch){
		
		l.readChar()
	}
	return l.input[position:l.position]


} 





func (l *Lexer)NextToken() token.Token{
		
		var tok token.Token
		l.skipWhiteSpace()	
		



		
		switch l.ch{

			case '=':
				tok = NewToken(token.ASSIGN, l.ch)
			case '+':
				tok = NewToken(token.PLUS, l.ch)
			case '*':
				tok = NewToken(token.ASTERIK, l.ch)
			case '-':
				tok = NewToken(token.MINUS, l.ch)
			case ')':
				tok = NewToken(token.RPAR, l.ch)
			case '(':
				tok = NewToken(token.LPAR, l.ch)
			case '}':
				tok = NewToken(token.RBRAC, l.ch)
			case '{':
				tok = NewToken(token.LBRAC, l.ch)
			case ',':
				tok = NewToken(token.COMMA, l.ch)
			case ';':
				tok = NewToken(token.SEMICOLON, l.ch)
			case '/':
				tok = NewToken(token.SLASH, l.ch)
			case '>':
				tok = NewToken(token.GTHAN, l.ch)

			case '<':
				tok = NewToken(token.LTHAN, l.ch)

			case 0:
				tok.Literal = ""
				tok.Type = token.EOF
			default:
				if isLetter(l.ch){
				if l.ch == 'n' && l.PeekWord(4) == "nOT="{

			tok = token.Token{Type:token.NEQ, Literal: "nOT=" }
			l.readChar()
			l.readChar()
			l.readChar()
			l.readChar()
			
			return tok;}else{tok.Literal = l.readIdentifier()
					tok.Type = token.Key_identifier(tok.Literal)  
					return tok;}
				}else if isDigit(l.ch){
					tok.Type = token.INT
					tok.Literal = l.readNumber()
					return tok
				}else{

					tok = NewToken(token.ILLEGAL,l.ch)
			}
			} 
		l.readChar()
		return tok

}



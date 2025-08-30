package token



type TokenType string

type Token struct{

	Type TokenType

	Literal string 
}


const(

	//Special characters
	ILLEGAL= "illegal"
	EOF = "eof"

	//Identifiers + literals

	IDENT = "ident"
	INT = "integers"	

	//operatos
	ASSIGN = "="
	PLUS = "+"
	MINUS  ="-"

		
	//delimiters
	LPAR = "Lparent"
	RPAR = "Rparent"
	LBRAC = "Lbrac"
	RBRAC = "Rbrac"
	COMMA = "Comma"
	SEMICOLON = "Semicolon"
	SLASH = "Slash" 

	//keywords
	LET = "LET"
	FUNCTION = "FUNCTION"
	RETURN = "RETURN"

	// Compairaison
	GTHAN = "GThan"
	LTHAN = "LThan"
	NEQ = "NOT_Equal"

	
	// BOOL
	


	// If else
	IF = "IF" 
	ELSE = "ELSE"

	//bool
	TRUE = "TRUE" 
	FALSE = "FALSE"
	
	ASTERIK = "*"

)


var keywords = map[string]TokenType{
	"fn" : FUNCTION,
	"if" : IF,
	"else" : ELSE,
	"let" : LET,
	"true" : TRUE,
	"false" : FALSE,
	"return" : RETURN,

}

// if ident found in keywords return its value else returns IDENT indicating a user variable name
func Key_identifier(ident string) TokenType{


	if tok, ok := keywords[ident]; ok{
			return tok
	}

	
	return IDENT
} 




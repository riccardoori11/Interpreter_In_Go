package repl

import (


	"bufio"
	"fmt"
	"io"
	"interpreter_go/token/lexer"
 	"interpreter_go/token/token"
) 


const Prompt = ">> \n"

func Start(in io.Reader, out io.Writer){

		scanner := bufio.NewScanner(in)

		for {


		fmt.Printf(Prompt)	

		scanned := scanner.Scan()
		if !scanned{

			return 	


		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken(){

			fmt.Printf("%+v", tok)

		}
		}


} 




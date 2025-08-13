package main

import(

	"fmt"
	"os"
	"os/user"
	"interpreter_go/token/repl"
)

func main(){


	user,err := user.Current()
	if err != nil{
		panic(err)
	}

	fmt.Printf("Hello %s, THis is ricc programming language! \n", user.Username)
	
	fmt.Printf("Start typing in commands\n")

	repl.Start(os.Stdin, os.Stdout)


}

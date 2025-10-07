package main

import (
	"arg-calc/lexer"
	"arg-calc/parse"
	"fmt"
	"os"
)

func main() {
	input := os.Args[1:len(os.Args)]
	Lexer := lexer.New(&input)
	lexer.Lex(&Lexer)
	parser := parse.Init_Parser(Lexer.Token_Buffer)
	output := parser.Parse()
	fmt.Println(output)
}

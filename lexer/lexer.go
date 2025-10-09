package lexer

import (
	"fmt"
	"os"
)

const (
	ADD    = "add"
	SUB    = "sub"
	MULT   = "mult"
	DIV    = "div"
	NUMBER = "num"
)

type Token struct {
	T      string
	Raw    string
	Weight int
}
type Lexer struct {
	input         *[]string
	Current_Token string
	Current_Pos   int
	Token_Buffer  []Token
}

func New(input *[]string) (self Lexer) {
	self.input = input
	self.Token_Buffer = make([]Token, 0, 10)
	self.Current_Token = (*self.input)[self.Current_Pos]
	self.Current_Pos++
	return
}
func check_if_valid_number(lexer *Lexer) {
	for _, ch := range lexer.Current_Token[1:len(lexer.Current_Token)] {
		if ch < '0' || ch > '9' {
			fmt.Printf("Character: %c, is not a valid number character", ch)
			if ch == '-' {

			} else {
				os.Exit(1)
			}
		}
	}
}
func Lex(lexer *Lexer) bool {
	if lexer.Current_Token[0] >= '0' && lexer.Current_Token[0] <= 'r' {
		check_if_valid_number(lexer)
		new_token := Token{
			T:      NUMBER,
			Raw:    lexer.Current_Token,
			Weight: 0,
		}
		lexer.Token_Buffer = append(lexer.Token_Buffer, new_token)
	} else if lexer.Current_Token[0] == '-' && lexer.Current_Token[len(lexer.Current_Token)-1] >= '0' && lexer.Current_Token[len(lexer.Current_Token)-1] <= '9' {
		check_if_valid_number(lexer)
		new_token := Token{
			T:      NUMBER,
			Raw:    lexer.Current_Token,
			Weight: 0,
		}
		lexer.Token_Buffer = append(lexer.Token_Buffer, new_token)
	} else {
		switch lexer.Current_Token {
		case "+":
			new_token := Token{
				T:      ADD,
				Raw:    "+",
				Weight: 1,
			}
			lexer.Token_Buffer = append(lexer.Token_Buffer, new_token)
		case "-":
			new_token := Token{
				T:      SUB,
				Raw:    "-",
				Weight: 1,
			}
			lexer.Token_Buffer = append(lexer.Token_Buffer, new_token)
		case "x":
			new_token := Token{
				T:      MULT,
				Raw:    "x",
				Weight: 2,
			}
			lexer.Token_Buffer = append(lexer.Token_Buffer, new_token)

		case "/":
			new_token := Token{
				T:      DIV,
				Raw:    "/",
				Weight: 2,
			}
			lexer.Token_Buffer = append(lexer.Token_Buffer, new_token)
		default:
			fmt.Printf("Invaild Operation: %s is not a known operation in the lexing stage", lexer.Current_Token)
			os.Exit(1)
		}
	}
	if lexer.Current_Pos < len(*lexer.input) {
		lexer.Current_Token = (*lexer.input)[lexer.Current_Pos]
		lexer.Current_Pos++
		return Lex(lexer)
	}
	return true
}

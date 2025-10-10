package parse

import (
	"fmt"
	"geocalc/lexer"
	"log"
	"os"
	"strconv"
)

type Parser struct {
	input       []lexer.Token
	current_pos int
}

func Init_Parser(input []lexer.Token) (ret Parser) {
	ret.input = input
	return
}
func (self *Parser) Parse() (result int) {
	output := parse_tokens(self, 0) // first input to accumulator should always be 0
	return output
}

// this needs to be cleaned up
func parse_tokens(parser *Parser, accumulator_in int) int {
	current_token := grab_token(parser)
	current_weight := 0
	local_accumulator := 0
	accumulator := accumulator_in
	if parser.current_pos == 1 {
		if current_token.T == lexer.NUMBER {
			accumulator = convtoint(current_token.Raw)
			current_token = parser.input[parser.current_pos]
		} else {
			fmt.Println("Cannot start expression with operator" + current_token.T)
		}
		current_token = grab_token(parser)
	}
	for parser.current_pos < len(parser.input) {
		switch current_token.T {
		case lexer.ADD:
			current_weight = 1
			number_token := grab_token(parser)
			local_accumulator = convtoint(number_token.Raw)
			next_operator_weight, isEOF := peek_token_weight(parser)
			if current_weight < next_operator_weight {
				accumulator += parse_tokens(parser, local_accumulator)
			} else if next_operator_weight == 0 && !isEOF {
				fmt.Println("Syntax Error: No operator given to number")
				os.Exit(1)
			} else {
				accumulator += local_accumulator
			}
		case lexer.SUB:
			current_weight = 1
			number_token := grab_token(parser)
			local_accumulator = convtoint(number_token.Raw)
			next_operator_weight, isEOF := peek_token_weight(parser)
			if current_weight < next_operator_weight {
				accumulator -= parse_tokens(parser, local_accumulator)
			} else if next_operator_weight == 0 && !isEOF {
				fmt.Println("Syntax Error: No operator given to number")
				os.Exit(1)
			} else {
				accumulator -= local_accumulator
			}

		case lexer.MULT:
			current_weight = 2
			number_token := grab_token(parser)
			local_accumulator = convtoint(number_token.Raw)
			next_operator_weight, _ := peek_token_weight(parser)
			if next_operator_weight > current_weight {
				next_operator := grab_token(parser)
				switch next_operator.T {
				case lexer.EXPONENT:
					local_accumulator = handl_expo(parser, local_accumulator)
				default:
					log.Fatalln("Unknown higher weight operator")
				}
			}
			accumulator = (accumulator * local_accumulator)
		case lexer.DIV:
			current_weight = 2
			number_token := grab_token(parser)
			local_accumulator := convtoint(number_token.Raw)
			next_operator_weight, _ := peek_token_weight(parser)
			if next_operator_weight > current_weight {
				local_accumulator = handl_expo(parser, local_accumulator)
			}
			accumulator = (accumulator / local_accumulator)

		case lexer.EXPONENT:
			accumulator = handl_expo(parser, accumulator)
		}
		if parser.current_pos < len(parser.input) {
			current_token = grab_token(parser)
		}

	}

	return accumulator
}
func handl_expo(parser *Parser, accumulator_in int) int {
	accumulator := accumulator_in
	number_token := grab_token(parser)
	local_accumulator := convtoint(number_token.Raw)
	base := accumulator
	for range local_accumulator - 1 {
		accumulator *= base
	}
	return accumulator
}
func grab_token(parser *Parser) (ret lexer.Token) {
	if parser.current_pos < len(parser.input) {
		ret = parser.input[parser.current_pos]
		parser.current_pos++
	}
	return ret
}
func peek_token_weight(parser *Parser) (ret int, EOF bool) {
	if parser.current_pos < len(parser.input) {
		ret = parser.input[parser.current_pos].Weight
	} else {
		EOF = true
	}
	return
}
func convtoint(input string) int {
	output, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	return output
}

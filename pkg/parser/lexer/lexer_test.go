package lexer_test

import (
	"Penguin/pkg/parser/lexer"
	"fmt"
	"testing"
)

var code string = `let x = 10;
function () {}
while {

}
if {

}

function add using (number x,number y) returns number{
	return x + y;
}
`

func TestGetNextToken(t *testing.T) {
	lexer := lexer.NewLexer(code)
	for {
		token := *lexer.GetNextToken()
		if token.Content == "eof" {
			break
		}
		println(fmt.Sprintf("%d:%d %s \"%s\"", lexer.Line, lexer.Column, string(token.Type), token.Content))
	}
}

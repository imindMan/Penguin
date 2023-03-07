package lexer_test

import (
	"Penguin/pkg/parser/lexer"
	"fmt"
	"log"
	"testing"
)

var code string = `
let x = 10;
function () {}
while {

}
if {

}
<=
>=
==
!=
++
--
+=
-=
/=
*=
function add using (number x,number y) returns number{
	return x + y;
}`

func TestGetNextToken(t *testing.T) {
	lxr := lexer.NewLexer(code)
	for {
		token := *lxr.GetNextToken()
		if token.Content == "eof" {
			break
		}
		if token.Content == "error" {
			log.Fatalf("ERROR tokenizing %d:%d", lxr.Line, lxr.Column)
		}
		println(fmt.Sprintf("%d:%d %s \"%s\"", lxr.Line, lxr.Column, string(token.Type), token.Content))
	}
}

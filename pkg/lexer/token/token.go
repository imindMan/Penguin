package token

import "fmt"

type Token struct {
	TokenType TokenType
	Value     string
	Location  Location
}

type Location struct {
	StartLine   int
	StartColumn int
	EndLine     int
	EndColumn   int
}

type TokenType string

// hold keywords
var keywords map[string]TokenType = map[string]TokenType{
	"character": Character,
	"string":    String,
	"number":    Number,
	"and":       And,
	"else":      Else,
	"fixed":     Fixed,
	"function":  Function,
	"if":        If,
	"let":       Let,
	"not":       Not,
	"nothing":   Nothing,
	"or":        Or,
	"return":    Return,
	"returns":   Returns,
	"using":     Using,
	"while":     While,
}

// List of TokenTypes
const (
	Illegal = "ILLEGAL"
	EOF     = "EOF"
	Comment = "COMMENT"

	Identifier       = "IDENTIFIER"
	CharacterLiteral = "CHARACTER LITERAL"
	StringLiteral    = "STRING LITERAL"
	NumberLiteral    = "NUMBER LITERAL"

	Add      = "ADD"      // +
	Subtract = "SUBTRACT" // -
	Multiply = "MULTIPLY" // *
	Divide   = "DIVIDE"   // /
	Mod      = "MOD"      // %

	AddAssign      = "ADD ASSIGN"      // +=
	SubtractAssign = "SUBTRACT ASSIGN" // -=
	MultiplyAssign = "MULTIPLY ASSIGN" // *=
	DivideAssign   = "DIVIDE ASSIGN"   // /=
	ModAssign      = "MOD ASSIGN"      // %=

	Increment = "INCREMENT" // ++
	Decrement = "DECREMENT" // --

	EqualsTo            = "EQUALS TO"              // ==
	LessThan            = "LESS THAN"              // <
	GreaterThan         = "GREATER THAN"           // >
	Assign              = "ASSIGN"                 // =
	NotEqualsTo         = "NOT EQUALS TO"          // !=
	LessThanEqualsTo    = "LESS THAN EQUALS TO"    // <=
	GreaterThanEqualsTo = "GREATER THAN EQUALS TO" // >=

	OpenParenthesis     = "OPEN PARENTHESIS"      // (
	OpenSquareBracket   = "OPEN SQUARE BRACKET"   // [
	OpenCurlyBracket    = "OPEN CURLY BRACKET"    // {
	Comma               = "COMMA"                 // ,
	Period              = "PERIOD"                // .
	ClosedParenthesis   = "CLOSED PARENTHESIS"    // )
	ClosedSquareBracket = "CLOSED SQUARE BRACKET" // ]
	ClosedCurlyBracket  = "CLOSED CURLY BRACKET"  // }
	Semicolon           = "SEMICOLON"             // ;

	Character = "CHARACTER"
	String    = "STRING"
	Number    = "NUMBER"
	And       = "AND"
	Else      = "ELSE"
	Fixed     = "FIXED"
	Function  = "FUNCTION"
	If        = "IF"
	Let       = "LET"
	Not       = "NOT"
	Nothing   = "NOTHING"
	Or        = "OR"
	Return    = "RETURN"
	Returns   = "RETURNS"
	Using     = "USING"
	While     = "WHILE"
)

func (tkn Token) String() string {
	return fmt.Sprintf("{ TokenType: %s, Value: %s }", string(tkn.TokenType), tkn.Value)
}

// LowestPrecedence represents lowest operator precedence.
const LowestPrecedence = 0

// Precedence returns the precedence for the operator token.
func (tkn Token) Precedence() int {
	switch tkn.TokenType {
	case Or:
		return LowestPrecedence + 1
	case And:
		return LowestPrecedence + 2
	case EqualsTo, NotEqualsTo, LessThan, LessThanEqualsTo, GreaterThan, GreaterThanEqualsTo:
		return LowestPrecedence + 3
	case Add, Subtract:
		return LowestPrecedence + 4
	case Multiply, Divide, Mod:
		return LowestPrecedence + 5
	}
	return LowestPrecedence
}

func GetKeywordOrIdentifier(s string) Token {
	if KeywordType, ok := keywords[s]; ok {
		return Token{TokenType: KeywordType, Value: s}
	}
	return Token{TokenType: Identifier, Value: s}
}

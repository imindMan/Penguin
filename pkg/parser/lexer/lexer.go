package lexer

import (
	"unicode"
)

type Lexer struct {
	Data    []rune //all src as array of rune
	AtIndex int    //current index
	Char    *rune  //pointer to rune at `atIndex` index in data array
	Line    int    //line number of ch
	Column  int    //column number of ch
}

/*
* Type: what type of token it is
* Content: actual content of token
 */
type Token struct {
	Type    TokenType
	Content string
}

type TokenType string

func NewLexer(data string) *Lexer {
	r := []rune(data)
	l := Lexer{
		Data:    r,
		AtIndex: 0,
		Line:    1,
		Column:  0,
	}
	l.nextChar()
	return &l
}

func (l *Lexer) GetNextToken() *Token {
	token := Token{Type: ERROR, Content: string(ERROR)}
	// Handling EOF
	if l.Char == nil {
		token = Token{Type: EOF, Content: string(EOF)}
	}
	l.skipWhitespace()
	switch {
	case l.isLetter():
		{
			word := l.readWhile(l.isLetter)
			token = getIdentifierOrKeyword(word)
		}
	case l.isDigit():
		{
			token = Token{Type: NUMBER_LITERAL, Content: l.readNumber()}
		}
	default:
		{
			token = *l.getOtherToken()
			if token.Type == ERROR {
				return &token
			}
		}
	}

	return &token
}

func (l *Lexer) nextChar() {
	if l.AtIndex >= len(l.Data) {
		l.Char = nil
		return
	}
	ch := l.Data[l.AtIndex]
	l.Char = &ch
	l.AtIndex++
	if ch == '\n' {
		l.Line++
		l.Column = 0
	} else {
		l.Column++
	}
}

func (l *Lexer) peakNextChar() *rune {
	idx := l.AtIndex + 1
	if idx >= len(l.Data) {
		return nil
	}
	return &l.Data[idx]
}
func (l *Lexer) skipWhitespace() {
	for l.Char != nil && unicode.IsSpace(*l.Char) {
		l.nextChar()
	}
}

func (l *Lexer) readWhile(predicate func() bool) string {
	result := ""
	for l.Char != nil && predicate() {
		result += string(*l.Char)
		l.nextChar()
	}
	return result
}

func (l *Lexer) isDigit() bool {
	return l.Char != nil && unicode.IsDigit(*l.Char)
}

func (l *Lexer) isLetter() bool {
	return l.Char != nil && unicode.IsLetter(*l.Char)
}

func (l *Lexer) readNumber() string {
	result := l.readWhile(l.isDigit)
	if *l.Char == '.' {
		result += string(*l.Char)
		l.nextChar()
		result += l.readWhile(l.isDigit)
	}
	return result
}

func (l *Lexer) readStringOrCharacter() string {
	word := l.readWhile(func() bool {
		return l.isLetter() || *l.Char == '\\' || *l.Char == '"' || *l.Char == '\''
	})
	return word
}

func (l *Lexer) getOtherToken() *Token {
	ch := l.Char
	token := Token{Type: ERROR, Content: string(ERROR)}
	if ch == nil {
		return &Token{Type: EOF, Content: string(EOF)}

	}
	switch *ch {
	case '"':
		{
			word := l.readStringOrCharacter()
			token = Token{Type: STRING_LITERAL, Content: word}
		}

	case '\'':
		{
			word := l.readStringOrCharacter()
			token = Token{Type: CHARACTER_LITERAL, Content: word}
		}

	case '+':
		{
			switch *l.peakNextChar() {
			case '+':
				{
					l.GetNextToken()
					token = Token{Type: INCREMENT, Content: string(INCREMENT)}
				}
			case '=':
				{
					l.GetNextToken()
					token = Token{Type: ADD_ASSIGN, Content: string(ADD_ASSIGN)}
				}
			default:
				{
					token = Token{Type: ADD, Content: string(ADD)}
				}
			}
		}

	case '-':
		{
			switch *l.peakNextChar() {
			case '-':
				{
					l.GetNextToken()
					token = Token{Type: DECREMENT, Content: string(DECREMENT)}
				}
			case '=':
				{
					l.GetNextToken()
					token = Token{Type: SUBTRACT_ASSIGN, Content: string(SUBTRACT_ASSIGN)}
				}
			default:
				{
					token = Token{Type: SUBTRACT, Content: string(SUBTRACT)}
				}
			}
		}

	case '*':
		{
			switch *l.peakNextChar() {
			case '=':
				{
					l.GetNextToken()
					token = Token{Type: MULTIPLY_ASSIGN, Content: string(MULTIPLY_ASSIGN)}
				}
			default:
				{
					token = Token{Type: MULTIPLY, Content: string(MULTIPLY)}
				}
			}
		}

	case '/':
		{
			switch *l.peakNextChar() {
			case '=':
				{
					l.GetNextToken()
					token = Token{Type: DIVIDE_ASSIGN, Content: string(DIVIDE_ASSIGN)}
				}
			default:
				{
					token = Token{Type: DIVIDE, Content: string(DIVIDE)}
				}
			}
		}

	case '%':
		{
			token = Token{Type: MODULUS, Content: string(MODULUS)}
		}

	case '^':
		{
			token = Token{Type: POWER, Content: string(POWER)}
		}
	case '<':
		{
			if *l.peakNextChar() == '=' {
				l.GetNextToken()
				token = Token{Type: LESS_THAN_EQUALS_TO, Content: string(LESS_THAN_EQUALS_TO)}
			} else {
				token = Token{Type: LESS_THAN, Content: string(LESS_THAN)}
			}
		}
	case '>':
		{
			if *l.peakNextChar() == '=' {
				l.GetNextToken()
				token = Token{Type: GREATER_THAN_EQUALS_TO, Content: string(GREATER_THAN_EQUALS_TO)}
			} else {
				token = Token{Type: GREATER_THAN, Content: string(GREATER_THAN)}
			}
		}
	case '=':
		{
			if *l.peakNextChar() == '=' {
				l.GetNextToken()
				token = Token{Type: EQUALS_TO, Content: string(EQUALS_TO)}
			} else {
				token = Token{Type: ASSIGN, Content: string(ASSIGN)}
			}
		}
	case '!':
		{
			if *l.peakNextChar() == '=' {
				l.GetNextToken()
				token = Token{Type: NOT_EQUALS_TO, Content: string(NOT_EQUALS_TO)}
			} else {
				token = Token{Type: LOGICAL_NOT, Content: string(LOGICAL_NOT)}
			}
		}
	case ',':
		{
			token = Token{Type: COMMA, Content: string(COMMA)}
		}
	case '.':
		{
			token = Token{Type: DOT, Content: string(DOT)}
		}
	case ';':
		{
			token = Token{Type: SEMICOLON, Content: string(SEMICOLON)}
		}
	case '(':
		{
			token = Token{Type: OPEN_PARENTHESES_BRACKET, Content: string(OPEN_PARENTHESES_BRACKET)}
		}
	case ')':
		{
			token = Token{Type: CLOSE_PARENTHESES_BRACKET, Content: string(CLOSE_PARENTHESES_BRACKET)}
		}
	case '{':
		{
			token = Token{Type: OPEN_CURLY_BRACE, Content: string(OPEN_CURLY_BRACE)}
		}
	case '}':
		{
			token = Token{Type: CLOSE_CURLY_BRACE, Content: string(CLOSE_CURLY_BRACE)}
		}
	case '[':
		{
			token = Token{Type: OPEN_SQUARE_BRACKET, Content: string(OPEN_SQUARE_BRACKET)}
		}
	case ']':
		{
			token = Token{Type: CLOSE_SQUARE_BRACKET, Content: string(CLOSE_SQUARE_BRACKET)}
		}
	}
	l.nextChar()
	return &token
}

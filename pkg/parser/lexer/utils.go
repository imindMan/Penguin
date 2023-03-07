package lexer

var keywords = map[string]TokenType{
	string(BOOLEAN_FALSE_LITERAL): BOOLEAN_FALSE_LITERAL,
	string(BOOLEAN_KEYWORD):       BOOLEAN_KEYWORD,
	string(BOOLEAN_TRUE_LITERAL):  BOOLEAN_TRUE_LITERAL,
	string(CHARACTER_KEYWORD):     CHARACTER_KEYWORD,
	string(ELSE_KEYWORD):          ELSE_KEYWORD,
	string(FIXED_KEYWORD):         FIXED_KEYWORD,
	string(FUNCTION_KEYWORD):      FUNCTION_KEYWORD,
	string(IF_KEYWORD):            IF_KEYWORD,
	string(LENGTH_OF_KEYWORD):     LENGTH_OF_KEYWORD,
	string(LET_KEYWORD):           LET_KEYWORD,
	string(LOGICAL_AND):           LOGICAL_AND,
	string(LOGICAL_OR):            LOGICAL_OR,
	string(NOTHING_KEYWORD):       NOTHING_KEYWORD,
	string(NUMBER_KEYWORD):        NUMBER_KEYWORD,
	string(RETURN_KEYWORD):        RETURN_KEYWORD,
	string(RETURNS_KEYWORD):       RETURNS_KEYWORD,
	string(STRING_KEYWORD):        STRING_KEYWORD,
	string(USING_KEYWORD):         USING_KEYWORD,
	string(WHILE_KEYWORD):         WHILE_KEYWORD,
}

func getIdentifierOrKeyword(s string) Token {
	if t, ok := keywords[s]; ok {
		return Token{Content: s, Type: t}
	}
	return Token{Type: IDENTIFIER, Content: s}
}

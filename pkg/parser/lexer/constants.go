package lexer

const (
	// Keywords
	BOOLEAN_KEYWORD   TokenType = `boolean`
	CHARACTER_KEYWORD TokenType = `character`
	ELSE_KEYWORD      TokenType = `else`
	FIXED_KEYWORD     TokenType = `fixed`
	FUNCTION_KEYWORD  TokenType = `function`
	IF_KEYWORD        TokenType = `if`
	LENGTH_OF_KEYWORD TokenType = `lengthOf`
	LET_KEYWORD       TokenType = `let`
	NOTHING_KEYWORD   TokenType = `nothing`
	NUMBER_KEYWORD    TokenType = `number`
	RETURN_KEYWORD    TokenType = `return`
	RETURNS_KEYWORD   TokenType = `returns`
	STRING_KEYWORD    TokenType = `string`
	USING_KEYWORD     TokenType = `using`
	WHILE_KEYWORD     TokenType = `while`
	// Operators
	// Unary Operators
	INCREMENT   TokenType = `++`
	DECREMENT   TokenType = `--`
	LOGICAL_NOT TokenType = `!`
	// Arithmetic Operators
	ADD      TokenType = `+`
	SUBTRACT TokenType = `-`
	MULTIPLY TokenType = `*`
	DIVIDE   TokenType = `/`
	MODULUS  TokenType = `%`
	POWER    TokenType = `^`
	// Relational Operators
	LESS_THAN              TokenType = `<`
	GREATER_THAN           TokenType = `>`
	EQUALS_TO              TokenType = `==`
	NOT_EQUALS_TO          TokenType = `!=`
	LESS_THAN_EQUALS_TO    TokenType = `<=`
	GREATER_THAN_EQUALS_TO TokenType = `>=`
	// Logical Operators
	LOGICAL_OR  TokenType = `or`
	LOGICAL_AND TokenType = `and`
	// Assignment Operators
	ASSIGN          TokenType = `=`
	ADD_ASSIGN      TokenType = `+=`
	SUBTRACT_ASSIGN TokenType = `-=`
	MULTIPLY_ASSIGN TokenType = `*=`
	DIVIDE_ASSIGN   TokenType = `/=`
	// Delimiters
	COMMA     TokenType = `,`
	SEMICOLON TokenType = `;`
	DOT       TokenType = `.`
	// Grouping Symbols
	OPEN_PARENTHESES_BRACKET  TokenType = `(`
	CLOSE_PARENTHESES_BRACKET TokenType = `)`
	OPEN_SQUARE_BRACKET       TokenType = `[`
	CLOSE_SQUARE_BRACKET      TokenType = `]`
	OPEN_CURLY_BRACE          TokenType = `{`
	CLOSE_CURLY_BRACE         TokenType = `}`
	// Literals
	NUMBER_LITERAL        TokenType = `number_literal`    // can be int or a float
	STRING_LITERAL        TokenType = `string_literal`    // anything enclosed in ``
	CHARACTER_LITERAL     TokenType = `character_literal` // anything enclosed in ''
	BOOLEAN_TRUE_LITERAL  TokenType = `true`
	BOOLEAN_FALSE_LITERAL TokenType = `false`
	//Identifier
	IDENTIFIER TokenType = `identifier`
	// Miscellaneous
	ERROR TokenType = `error`
	EOF   TokenType = `eof`
)

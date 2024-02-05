package token

// possible token types
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers
	IDENTIFIER = "IDENTIFIER"
	INT        = "INT"

	// Operators
	ASSIGNMENT = "="
	EQ         = "=="
	STRICT_EQ  = "==="
	NOT        = "!"
	NOT_EQ     = "!="
	LT         = "<"
	LT_EQ      = "<="
	GT         = ">"
	GT_EQ      = ">="
	MODULO     = "%"
	PLUS       = "+"
	MINUS      = "-"
	DIVIDE     = "/"
	MULTIPLY   = "*"

	PLUS_EQ     = "+="
	MINUS_EQ    = "-="
	DIVIDE_EQ   = "/="
	MULTIPLY_EQ = "*="

	INCREMENT = "+"
	DECREMENT = "--"

	//  Delimeter
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"
	QUESTION  = "?"

	// BRACKETS
	LEFT_PARENTHESIS  = "("
	RIGHT_PARENTHESIS = ")"
	LEFT_BRACE        = "{"
	RIGHT_BRACE       = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	CONST    = "CONST"
	IF       = "IF"
	ELSE     = "ELSE"
	FOR      = "FOR"
	BREAK    = "BREAK"
	CONTINUE = "CONTINUE"
	RETURN   = "RETURN"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
)

var keywords = map[string]TokenType{
	"function": FUNCTION,
	"let":      LET,
	"const":    CONST,
	"if":       IF,
	"else":     ELSE,
	"for":      FOR,
	"break":    BREAK,
	"continue": CONTINUE,
	"return":   RETURN,
	"true":     TRUE,
	"false":    FALSE,
}

// checks identifiers against the keywords table
func LookupKeywords(identifier string) TokenType {
	if tok, ok := keywords[identifier]; ok {
		return tok
	}
	//  If not seen in keywords table return IDENTIFIER token
	return IDENTIFIER
}

// create new Token (token type and literal)
func NewToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

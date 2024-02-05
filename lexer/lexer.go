package lexer

// maps inputed source code or string to Lexer struct
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

//	check if input char under examination is an identifier
//
// e.g variables can be a letter, mixed with numbers and an underscore but never starting with a number
func isLetter(ch byte) bool {
	//  check within a...z or A..Z or _
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// checks if input char under examination is a number
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

package lexer

import "github.com/miriam-samuels/interpreter/token"

type Lexer struct {
	input       string
	line        int  // current line number
	position    int  // current pos in input
	readPostion int  // after current pos/ char
	ch          byte //current char under examination
}

//  reads char under examination and sets the next char
func (l *Lexer) readChar() {
	// check if read position is greater than length of input
	if l.readPostion >= len(l.input) {
		l.ch = 0 // end of file EOF
	} else {
		l.ch = l.input[l.readPostion] // current character under examination
	}

	l.position = l.readPostion // set position of cursor to current readPosition

	l.readPostion += 1 // set read position to next character
}

//  peeks at char under examination
func (l *Lexer) peekChar() byte {
	// check if read position is greater than length of input
	if l.readPostion >= len(l.input) {
		return 0 // end of file EOF
	} else {
		return l.input[l.readPostion] // current character under examination
	}
}

func (l *Lexer) readIdentifier() string {
	// save the starting position
	pos := l.position

	// keep looping throught input charaters till isLetter returns false
	for isLetter(l.ch) {
		l.readChar() // read current .. then moves to next
	}
	return l.input[pos:l.position] // retuns char from starting pos till last pos of identifier
}

func (l *Lexer) readNumber() string {
	// save the starting position
	position := l.position
	// keep looping throught input charaters till isDigit returns false
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position] // retuns char from starting pos till last pos of identifier
}

// skip all whitespaces including tabs , space, newline e.t.c.
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// looks at current char under examination against existing token types and return token depending on the character
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case ';':
		tok = token.NewToken(token.SEMICOLON, l.ch)
	case ',':
		tok = token.NewToken(token.COMMA, l.ch)
	case '(':
		tok = token.NewToken(token.LEFT_PARENTHESIS, l.ch)
	case ')':
		tok = token.NewToken(token.RIGHT_PARENTHESIS, l.ch)
	case '{':
		tok = token.NewToken(token.LEFT_BRACE, l.ch)
	case '}':
		tok = token.NewToken(token.RIGHT_BRACE, l.ch)
	case '%':
		tok = token.NewToken(token.MODULO, l.ch)
	case '=':
		// check if it's == instead of =
		if l.peekChar() == '=' {
			ch := l.ch   //store current char
			l.readChar() // go to next char
			// check if it's === instead of ==
			if l.peekChar() == '=' {
				ch := l.ch   // store current char
				l.readChar() // go to next char
				tok = token.Token{Type: token.STRICT_EQ, Literal: string(ch) + string(l.ch)}
			} else {
				tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
			}
		} else {
			tok = token.NewToken(token.ASSIGNMENT, l.ch)
		}
	case '+':
		// check if it's += instead of +
		if l.peekChar() == '=' {
			ch := l.ch   // store current char
			l.readChar() // go to next char
			tok = token.Token{Type: token.PLUS_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = token.NewToken(token.PLUS, l.ch)
		}
	case '-':
		// check if it's -= instead of -
		if l.peekChar() == '=' {
			ch := l.ch   // store current char
			l.readChar() // go to next char
			tok = token.Token{Type: token.MINUS_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = token.NewToken(token.MINUS, l.ch)
		}
	case '/':
		// check if it's /= instead of /
		if l.peekChar() == '=' {
			ch := l.ch   // store current char
			l.readChar() // go to next char
			tok = token.Token{Type: token.DIVIDE_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = token.NewToken(token.DIVIDE, l.ch)
		}
	case '*':
		// check if it's *= instead of *
		if l.peekChar() == '=' {
			ch := l.ch   // store current char
			l.readChar() // go to next char
			tok = token.Token{Type: token.MULTIPLY_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = token.NewToken(token.MULTIPLY, l.ch)
		}
	case '!':
		// check if it's != instead of !
		if l.peekChar() == '=' {
			ch := l.ch   // store current char
			l.readChar() // go to next char
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = token.NewToken(token.NOT, l.ch)
		}
	case '>':
		// check if it's >= instead of >
		if l.peekChar() == '=' {
			ch := l.ch   // store current char
			l.readChar() // go to next char
			tok = token.Token{Type: token.GT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = token.NewToken(token.GT, l.ch)
		}
	case '<':
		// check if it's <= instead of <
		if l.peekChar() == '=' {
			ch := l.ch   // store current char
			l.readChar() // go to next char
			tok = token.Token{Type: token.LT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = token.NewToken(token.LT, l.ch)
		}
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupKeywords(tok.Literal)
			return tok

		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok

		} else {
			tok = token.NewToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar() // change char under examination and update positions

	return tok
}

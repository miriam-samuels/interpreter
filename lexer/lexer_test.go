package lexer

import (
	"testing"

	"github.com/miriam-samuels/interpreter/token"
)

func TestToken(t *testing.T) {
	input := `
		let five = 5;
		let ten = 10;
		let add = fn(x, y) {
		  x + y;
		};
		let result = add(five, ten);
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		// test cases for each char
		{token.LET, "let"},
		{token.IDENTIFIER, "five"},
		{token.ASSIGNMENT, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "ten"},
		{token.ASSIGNMENT, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "add"},
		{token.ASSIGNMENT, "="},
		{token.FUNCTION, "fn"},
		{token.LEFT_PARENTHESIS, "("},
		{token.IDENTIFIER, "x"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "y"},
		{token.RIGHT_PARENTHESIS, ")"},
		{token.LEFT_BRACE, "{"},
		{token.IDENTIFIER, "x"},
		{token.PLUS, "+"},
		{token.IDENTIFIER, "y"},
		{token.SEMICOLON, ";"},
		{token.RIGHT_BRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "result"},
		{token.ASSIGNMENT, "="},
		{token.IDENTIFIER, "add"},
		{token.LEFT_PARENTHESIS, "("},
		{token.IDENTIFIER, "five"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "ten"},
		{token.RIGHT_PARENTHESIS, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	// create lexer struct (input,ch,pos,readPos...) from input
	l := New(input)

	// loop through test cases
	for i, tt := range tests {
		//  gets token that reps current char under examination
		tok := l.NextToken()

		//  compare tokenized char (type,literal) to expexted char from test case
		if tok.Type != tt.expectedType {
			t.Fatalf("test[%d]: token type wrong. expected %q got=%q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("test[%d]: literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

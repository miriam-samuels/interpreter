package repl // READ EVAL PRINT LOOP

import (
	"bufio"
	"fmt"
	"io"

	"github.com/miriam-samuels/interpreter/lexer"
	"github.com/miriam-samuels/interpreter/token"
)

const PROMPT = "- %%"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	lineNumber := 1

	//  start infinte loop to keep reading till no more
	for {
		fmt.Printf(PROMPT)

		scanned := scanner.Scan() //  scan line

		// if there is nothing scanned means we are at end of input
		if !scanned {
			fmt.Println("Error reading input.")
			return // break out of ingoop
		}

		line := scanner.Text() // retrieve the text of scanned line

		l := lexer.New(line) // creates lexical analyzer for line

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}

		lineNumber++ // increment line NUmber
	}

}

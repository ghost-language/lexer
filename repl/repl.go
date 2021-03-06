package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/ghost-language/lexer/lexer"
	"github.com/ghost-language/lexer/token"
)

const prompt = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(prompt)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		lexer := lexer.New(line)

		for tok := lexer.NextToken(); tok.Type != token.EOF; tok = lexer.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}

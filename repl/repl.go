package repl

import (
	"bufio"
	"fmt"
	"io"

	"rolang/lexer"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		fmt.Println(line)

		fmt.Println(l.NextToken())
	}
}

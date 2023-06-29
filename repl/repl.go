package repl

import (
	"bufio"
	"fmt"
	"monkey/lexer"
	"monkey/token"
	"os"
	"os/user"
)

func Start() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")

	in, out := os.Stdin, os.Stdout
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, ">> ")
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.Eof; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}

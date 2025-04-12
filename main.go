package main

import (
	"bufio"
	"fmt"
	"monkey/scanner"
	"os"
	"strings"
)

func repl() {
	for {
		fmt.Printf(">")

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("read input error: %v\n", err.Error())
			return
		}
		if strings.TrimSpace(input) == "exit" {
			return
		}

		s := scanner.NewScanner(input)
		tokens := s.ScanTokens()
		for _, t := range tokens {
			fmt.Printf("%s", t.String())
		}
		fmt.Println("")
	}
}

func main() {
	if len(os.Args) == 1 {
		repl()

		return
	}
}

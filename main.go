package main

import (
	"fmt"
	"os"
	"yikes/parser"

	"github.com/containerd/console"
	"golang.org/x/term"
)

func main() {
	fmt.Printf(">>> yikes %v %v \n", parser.Version, parser.Github)

	if len(os.Args) < 2 || os.Args[1] == "" {
		current := console.Current()
		defer current.Reset()

		if err := current.SetRaw(); err != nil {
			panic(err)
		}
		term := term.NewTerminal(current, "")
		term.AutoCompleteCallback = func(line string, pos int, key rune) (newLine string, newPos int, ok bool) {
			return "", 0, false
		}

		println(">>> Entered interactive mode:")
		print(">>> ")
		for {
			for {
				line, err := term.ReadLine()
				if err != nil {
					println("\nbye")
					return
				}
				if line == "exit" {
					println("bye")
					return
				}
				parser.Run(line)
				print(">>> ")
			}
		}
	}

	fmt.Printf(">>> Reading file %v\n", os.Args[1])
	f, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	parser.Run(string(f))
}

package main

import (
	"flag"
	"fmt"
	"os"
	"yikes/parser"

	"github.com/containerd/console"
	"golang.org/x/term"
)

var (
	debug = flag.Bool("debug", false, "Enable debug mode")
	file  = flag.String("file", "", "Script file")
)

func main() {
	flag.Parse()
	if *debug {
		parser.EnableDebug()
	}

	fmt.Printf("\n Welcome to yikes %v\n Github: %v \n", parser.Version, parser.Github)

	if *file == "" {
		current := console.Current()
		defer current.Reset()

		if err := current.SetRaw(); err != nil {
			panic(err)
		}
		term := term.NewTerminal(current, "")
		term.AutoCompleteCallback = func(line string, pos int, key rune) (newLine string, newPos int, ok bool) {
			return "", 0, false
		}

		println("\n>>> Entering interactive mode:")
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

	fmt.Printf(">>> Reading file %v\n", *file)
	f, err := os.ReadFile(*file)
	if err != nil {
		panic(err)
	}
	parser.Run(string(f))
}

package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jumballaya/servo/repl"
)

func run(fromFile bool, config *repl.Config) {
	if !fromFile {
		fmt.Printf("Hello! This is the Servo programming language!\n")
		fmt.Printf("Feel free to type commands\n")
		repl.Start(os.Stdin, os.Stdout, config)
	} else {
		file, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			fmt.Printf("Error has occured: %q", err)
			return
		}
		repl.Run(string(file), os.Stdout, config)
	}
}

func main() {
	config := &repl.Config{Verbose: true}
	run(len(os.Args) > 1, config)
}

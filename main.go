package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jumballaya/servo/repl"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Hello! This is the Servo programming language!\n")
		fmt.Printf("Feel free to type commands\n")
		repl.Start(os.Stdin, os.Stdout)
	} else {
		file, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			fmt.Printf("Error has occured: %q", err)
			return
		}

		repl.Run(string(file), os.Stdout, false)
	}

}

package main

import (
	"os"
	"os/user"

	"github.com/abhi-abhiram/interpreter/repl"
)

func main() {
	_, err := user.Current()
	if err != nil {
		print(err)
	}

	repl.Start(os.Stdin, os.Stdout)

}

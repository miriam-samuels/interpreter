package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/miriam-samuels/interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Mimi programming language!\n", user.Username)
	fmt.Println("What are you waiting for, type in your commands")

	repl.Start(os.Stdin, os.Stdout)
}

package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	curUser, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", curUser.Username)
	fmt.Println("Feel free to type a command")
	repl.Start(os.Stdin, os.Stdout)
}

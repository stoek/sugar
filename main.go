package main

import (
	"fmt"
	"os"
	"os/user"
	"sugar/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Sugar programming language!\n",
		user.Username)
	fmt.Printf("This is a REPL-ish, feel free to type some stuff to see it being converted to tokens! \n")
	repl.Start(os.Stdin, os.Stdout)
}

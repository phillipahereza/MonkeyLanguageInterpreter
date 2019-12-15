package main

import (
	"fmt"
	"github.com/phillipahereza/MonkeyLanguageInterpreter/repl"
	"os"
	osUser "os/user"
)

func main() {
	user, err := osUser.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey Programming Language\n", user.Username)
	fmt.Printf("Feel free to type in commands \n")

	repl.Start(os.Stdin, os.Stdout)

}

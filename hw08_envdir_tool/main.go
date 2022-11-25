package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 3 {
		fmt.Println("at least 2 argument required: envDir and command")
		os.Exit(1)
	}

	env, err := ReadDir(args[1])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	os.Exit(RunCmd(args[2:], env))
}

package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	if len(cmd) == 0 {
		fmt.Println("cmd required")
		return 1
	}

	for key, val := range env {
		if val.NeedRemove {
			err := os.Unsetenv(key)
			if err != nil {
				fmt.Println(err.Error())
				return 1
			}
		} else {
			err := os.Setenv(key, val.Value)
			if err != nil {
				fmt.Println(err.Error())
				return 1
			}
		}
	}

	execCommand := cmd[0]
	command := exec.Command(execCommand, cmd[1:]...)
	command.Stdout = os.Stdout

	if err := command.Run(); err != nil {
		var exitError *exec.ExitError

		if errors.As(err, &exitError) {
			return exitError.ExitCode()
		}
	}

	return
}

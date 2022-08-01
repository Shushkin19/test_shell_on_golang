package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)
var ErrNoPath = errors.New("path required")

func str_exec(input string) error {
	input = strings.TrimSuffix(input, "\n")
	args := strings.Split(input, " ")
	switch args[0] {
    case "cd":
        // 'cd' to home with empty path not yet supported.
        if len(args) < 2 {
            return ErrNoPath
        }
        // Change the directory and return the error.
        return os.Chdir(args[1])
    case "exit":
        os.Exit(0)
    }
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func main() {
	reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Print("> ")
        // Read the keyboad input.
        input, err := reader.ReadString('\n')
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
        }

        // Handle the execution of the input.
        if err = str_exec(input); err != nil {
            fmt.Fprintln(os.Stderr, err)
        }
    }
}

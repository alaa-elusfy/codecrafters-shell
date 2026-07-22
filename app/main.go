package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	var command string
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("$ ")
		if scanner.Scan() {
			command = scanner.Text()
		}

		command = strings.TrimSpace(command)
		if command == "exit" {
			break
		} else if strings.HasPrefix(command, "echo") {
			_, args, _ := strings.Cut(command, " ")
			fmt.Println(args)
		} else if strings.HasPrefix(command, "type") {
			checkType(command)
		} else {
			fmt.Printf("%s: command not found\n", command)
		}
	}
}

func checkType(command string) {
	_, cmd, _ := strings.Cut(command, " ")
	builtins := []string{"echo", "exit", "type"}

	if slices.Contains(builtins, cmd) {
		fmt.Printf("%s is a shell builtin\n", cmd)
	} else {
		envPath := os.Getenv("PATH")
		paths := strings.Split(envPath, string(os.PathListSeparator))

		cmdExist := false
		for _, path := range paths {
			cmdPath := path + string(os.PathSeparator) + cmd

			fileInfo, err := os.Stat(cmdPath)
			if err != nil {
				continue
			}

			if fileInfo.Mode().Perm()&0111 != 0 {
				cmdExist = true
				fmt.Printf("%s is %s\n", cmd, cmdPath)
				break
			}
		}

		if !cmdExist {
			fmt.Printf("%s: not found\n", cmd)
		}
	}
}

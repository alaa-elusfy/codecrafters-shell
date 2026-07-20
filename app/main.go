package main

import (
	"bufio"
	"fmt"
	"os"
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
		}

		if strings.HasPrefix(command, "echo") {
			_, args, _ := strings.Cut(command, " ")
			fmt.Println(args)
			continue
		}

		fmt.Printf("%s: command not found\n", command)
	}
}

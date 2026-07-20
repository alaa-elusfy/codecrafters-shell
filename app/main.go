package main

import (
	"fmt"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func exit() {
	os.Exit(0)
}

func main() {
	var command string

	for {
		fmt.Print("$ ")
		fmt.Scanf("%s", &command)
		if command == "exit" {
			exit()
		}
		fmt.Printf("%s: command not found\n", command)
	}
}

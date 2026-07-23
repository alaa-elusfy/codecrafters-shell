package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
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
		} else if command == "pwd" {
			pwd, err := os.Getwd()
			if err != nil {
				log.Print("Can't Get Working Directory", err)
			}
			fmt.Println(pwd)

		} else if strings.HasPrefix(command, "cd") {
			_, targetDir, _ := strings.Cut(command, " ")

			if targetDir == "~" {
				targetDir = os.Getenv("HOME")
			}

			if err := os.Chdir(targetDir); err != nil {
				fmt.Printf("cd: %s: No such file or directory\n", targetDir)
			}
		} else if strings.HasPrefix(command, "echo") {
			_, args, _ := strings.Cut(command, " ")
			fmt.Println(args)
		} else if strings.HasPrefix(command, "type") {
			checkType(command)
		} else {
			args := strings.Fields(command)
			if exist, _ := checkCmdInPath(args[0]); exist {
				execCmd := exec.Command(args[0], args[1:]...)
				out, err := execCmd.CombinedOutput()
				if err != nil {
					log.Print("Error executing command:", err)
				}

				fmt.Print(string(out))

			} else {
				fmt.Printf("%s: command not found\n", command)
			}

		}
	}
}

func checkType(command string) {
	_, cmd, _ := strings.Cut(command, " ")
	builtins := []string{"echo", "exit", "type", "pwd"}

	if slices.Contains(builtins, cmd) {
		fmt.Printf("%s is a shell builtin\n", cmd)
	} else {

		if exist, path := checkCmdInPath(cmd); !exist {
			fmt.Printf("%s: not found\n", cmd)
		} else {
			fmt.Print(path)
		}
	}
}

func checkCmdInPath(cmd string) (bool, string) {
	envPath := os.Getenv("PATH")
	paths := strings.Split(envPath, string(os.PathListSeparator))

	for _, path := range paths {
		cmdPath := path + string(os.PathSeparator) + cmd

		fileInfo, err := os.Stat(cmdPath)
		if err != nil {
			continue
		}

		if fileInfo.Mode().Perm()&0111 != 0 {
			return true, fmt.Sprintf("%s is %s\n", cmd, cmdPath)
		}
	}

	return false, ""
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var _ = fmt.Print

func handleEcho(s []string) {
	for i := 0; i < len(s); i++ {
		if i != len(s)-1 {
			fmt.Printf("%s ", s[i])
		} else {
			fmt.Printf("%s\n", s[i])
		}
	}
}

func handleType(s string) {
	typesList := []string{"echo", "exit", "type"}

	for i := 0; i < len(typesList); i++ {
		if s == typesList[i] {
			fmt.Printf("%s is a shell builtin\n", s)
			return
		}
	}

	fmt.Printf("%s: not found\n", s)
}

func getCommandsList(cmd string) []string {
	return strings.Split(cmd, " ")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	status := 1
	for status == 1 {
		fmt.Print("$ ")

		cmd, err := reader.ReadString('\n')

		if err != nil {
			panic(err)
		}

		cmd = cmd[0 : len(cmd)-1]
		cmdList := getCommandsList(cmd)

		if cmdList[0] == "type" {
			handleType(cmdList[1])
			continue
		}

		if cmdList[0] == "exit" {
			status = 0
			break
		}

		if cmdList[0] == "echo" {
			handleEcho(cmdList[1:])
			continue
		}

		fmt.Printf("%s: command not found\n", cmd)
	}

}

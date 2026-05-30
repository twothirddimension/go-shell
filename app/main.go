package main

import (
	"bufio"
	"fmt"
	"os"
)

var _ = fmt.Print

func handleEcho(s string) {
	fmt.Println(s)
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

		if cmd == "exit" {
			status = 0
			break
		}

		// echo -
		if cmd[0:5] == "echo " {
			handleEcho(cmd[5:])
			continue
		}

		fmt.Printf("%s: command not found\n", cmd)
	}

}

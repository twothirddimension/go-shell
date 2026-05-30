package main

import (
	"fmt"
)

var _ = fmt.Print

func main() {
	status := 1
	for status == 1 {
		var cmd string
		fmt.Print("$ ")
		fmt.Scanln(&cmd)

		if cmd == "exit" {
			status = 0
			break
		}

		fmt.Printf("%s: command not found\n", cmd)
	}

}

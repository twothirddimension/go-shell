package main

import (
	"fmt"
)

var _ = fmt.Print

func main() {
	// TODO: Uncomment the code below to pass the first stage
	var cmd string
	fmt.Print("$ ")
	fmt.Scanln(&cmd)

	fmt.Printf("%s: command not found\n", cmd)

}

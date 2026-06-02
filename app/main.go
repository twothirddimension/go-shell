package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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

func handlePWD() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println(wd)
}

func handleType(s string) {
	typesList := []string{"echo", "exit", "type", "pwd", "cd"}

	for i := 0; i < len(typesList); i++ {
		if s == typesList[i] {
			fmt.Printf("%s is a shell builtin\n", s)
			return
		}
	}

	if fullPath, err := exec.LookPath(s); err == nil {
		fmt.Printf("%s is %s\n", s, fullPath)
		return
	}

	fmt.Printf("%s: not found\n", s)
}

func getCommandsList(cmd string) []string {
	return strings.Split(cmd, " ")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	pathValues := ""
	status := 1
	for status == 1 {
		fmt.Print("$ ")

		cmd, err := reader.ReadString('\n')

		if err != nil {
			panic(err)
		}

		cmd = cmd[0 : len(cmd)-1]
		cmd = strings.TrimSpace(cmd)
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

		if cmdList[0] == "pwd" {
			if len(cmdList) == 1 {
				handlePWD()
				continue
			}
		}

		if cmdList[0] == "cd" {
			target := os.Getenv("HOME")
			if len(cmdList) > 1 {
				target = cmdList[1]
			}

			if strings.HasPrefix(target, "~") {
				home, err := os.UserHomeDir()
				if err == nil {
					target = filepath.Join(home, strings.TrimPrefix(target, "~"))
				} else {
					target = filepath.Join(os.Getenv("HOME"), strings.TrimPrefix(target, "~"))
				}
			}

			absPath, err := filepath.Abs(target)
			if err != nil {
				fmt.Printf("Error resolving path %s\n", err.Error())
				continue
			}

			realPath, err := filepath.EvalSymlinks(absPath)
			if err != nil {
				realPath = absPath
			}

			info, err := os.Stat(realPath)
			if err != nil {
				fmt.Printf("cd: %s: No such file or directory\n", target)
				continue
			}

			if !info.IsDir() {
				fmt.Printf("cd: %s: Not a directory\n", target)
				continue
			}

			err = os.Chdir(realPath)
			if err != nil {
				fmt.Printf("cd: %s: No such file or directory\n", target)
			}
			continue
		}

		if strings.Contains(cmd, "PATH=") {
			parts := strings.SplitN(cmd, "=", 2)
			if len(parts) == 2 && parts[0] == "PATH" {
				pathValues = os.ExpandEnv(parts[1])
				os.Setenv("PATH", pathValues)
				continue
			}
		}

		// Run external program
		if fullPath, err := exec.LookPath(cmdList[0]); err == nil {
			cmdObj := exec.Command(fullPath, cmdList[1:]...)
			cmdObj.Args[0] = cmdList[0]
			cmdObj.Stdout = os.Stdout
			cmdObj.Stderr = os.Stderr
			cmdObj.Stdin = os.Stdin
			err := cmdObj.Run()
			if err != nil {
				fmt.Print(err)
			}
			continue
		}

		fmt.Printf("%s: command not found\n", cmd)
	}

}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	// Uncomment this block to pass the first stage

	// Wait for user input
	for {
		fmt.Fprint(os.Stdout, "$ ")
		command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		split := strings.Split(command, " ")
		if len(split) >= 2 && split[0] == "exit" && strings.TrimSpace(split[1]) == "0" {
			//fmt.Println("exit with code: ", split[1])
			os.Exit(0)
		}
		fmt.Fprintf(os.Stderr, "%s: command not found\n", strings.TrimSpace(command))
	}
}

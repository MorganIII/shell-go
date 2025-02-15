package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	for {
		fmt.Fprint(os.Stdout, "$ ")
		command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		if len(command) < 1 {
			log.Fatal("invalid command")
		}

		split := strings.Split(command, " ")
		command = split[0]
		args := split[1:]
		c := Command{Name: command, Args: args}
		c.trimArgSpaces()
		c.handleCommand()
	}

}

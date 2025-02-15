package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	Name string
	args []string
}

func (c Command) handleCommand() {
	switch c.Name {
	case "exit":
		if len(c.args) != 1 {
			log.Fatal("Usage: exit <exit code>")
		}
		code, err := strconv.Atoi(c.args[0])
		if err != nil || code < 0 || code > 255 {
			log.Fatal(err)
		}
		os.Exit(code)
	case "echo":
		echoed := strings.Join(c.args, " ")
		fmt.Println(echoed)
	default:
		fmt.Fprintf(os.Stderr, "%s: command not found\n", strings.TrimSpace(c.Name))
	}
}

func (c *Command) trimArgSpaces() {
	for i := range c.args {
		c.args[i] = strings.TrimSpace(c.args[i])
	}
}

func main() {
	// Uncomment this block to pass the first stage

	// Wait for user input
	for {
		fmt.Fprint(os.Stdout, "$ ")
		command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		if len(command) < 1 {
			log.Fatal("invalid command")
		}

		split := strings.Split(command, " ")
		command = split[0]
		args := split[1:]
		c := Command{Name: command, args: args}
		c.trimArgSpaces()
		c.handleCommand()
	}
}

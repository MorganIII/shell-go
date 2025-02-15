package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	Name string
	Args []string
}

func (c Command) handleCommand() {
	switch c.Name {
	case "exit":
		if len(c.Args) != 1 {
			log.Fatal("Usage: exit <exit code>")
		}
		code, err := strconv.Atoi(c.Args[0])
		if err != nil || code < 0 || code > 255 {
			log.Fatal(err)
		}
		os.Exit(code)
	case "echo":
		echoed := strings.Join(c.Args, " ")
		fmt.Println(echoed)
	case "type":
		command := c.Args[0]
		if _, ok := shellCommands[command]; !ok {
			fmt.Fprintf(os.Stderr, "%s: not found\n", command)
		} else {
			fmt.Printf("%s is a shell builtin\n", command)
		}
	default:
		fmt.Fprintf(os.Stderr, "%s: command not found\n", strings.TrimSpace(c.Name))
	}
}

func (c *Command) trimArgSpaces() {
	for i := range c.Args {
		c.Args[i] = strings.TrimSpace(c.Args[i])
	}
}

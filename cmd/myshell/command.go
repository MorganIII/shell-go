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
		if _, ok := shellCommands[command]; ok {
			fmt.Printf("%s is a shell builtin\n", command)
		} else if path := c.lookupCommand(); path != "" {
			fmt.Printf("%s is %s\n", c.Args[0], path)
		} else {
			fmt.Fprintf(os.Stderr, "%s: not found\n", command)
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

func (c Command) lookupCommand() string {
	path := os.Getenv("PATH")
	existPaths := strings.Split(path, string(os.PathListSeparator))
	for i := len(existPaths) - 1; i >= 0; i-- {
		entries, err := os.ReadDir(existPaths[i])
		if err != nil {
			continue
		}
		for _, e := range entries {
			if !e.Type().IsDir() && e.Name() == c.Args[0] {
				return existPaths[i] + string(os.PathSeparator) + e.Name()
			}
		}
	}
	return ""
}

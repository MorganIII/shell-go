package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
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
		var buildCommand Command
		buildCommand.Name = command
		if _, ok := shellCommands[command]; ok {
			fmt.Printf("%s is a shell builtin\n", command)
		} else if path := buildCommand.lookupCommand(); path != "" {
			fmt.Printf("%s is %s\n", c.Args[0], path)
		} else {
			fmt.Fprintf(os.Stderr, "%s: not found\n", command)
		}
	case "pwd":
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(dir)
	default:
		if path := c.lookupCommand(); path != "" {
			output, err := exec.Command(c.Name, c.Args[:]...).Output()
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Printf("%s", string(output))
				return
			}
		}
		fmt.Fprintf(os.Stderr, "%s: command not found\n", strings.TrimSpace(c.Name))

	}
}

func (c *Command) trimArgSpaces() {
	for i := range c.Args {
		c.Args[i] = strings.TrimSpace(c.Args[i])
	}
	c.Name = strings.TrimSpace(c.Name)
}

func (c Command) lookupCommand() string {
	path := os.Getenv("PATH")
	existPaths := strings.Split(path, string(os.PathListSeparator))
	for i := len(existPaths) - 1; i >= 0; i-- {
		existPath := existPaths[i]
		entries, err := os.ReadDir(existPath)
		if err != nil {
			continue
		}
		for _, e := range entries {
			filePath := filepath.Join(existPath, e.Name())
			if !e.Type().IsDir() && e.Name() == c.Name && checkFileExec(filePath) {
				return filePath
			}
		}
	}
	return ""
}

func checkFileExec(filePath string) bool {
	isExecutable := false
	if runtime.GOOS == "windows" {
		isExecutable, _ = filepath.Match("*.exe", filepath.Base(filePath))
	} else {
		info, err := os.Stat(filePath)
		if err == nil {
			isExecutable = info.Mode()&0111 != 0
		}
	}
	return isExecutable
}

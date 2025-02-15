package main

const (
	EXIT = "exit"
	ECHO = "echo"
	TYPE = "type"
)

var shellCommands = map[string]struct{}{
	EXIT: {},
	ECHO: {},
	TYPE: {},
}

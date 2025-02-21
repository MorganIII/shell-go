package main

const (
	EXIT = "exit"
	ECHO = "echo"
	TYPE = "type"
	PWD  = "pwd"
)

var shellCommands = map[string]struct{}{
	EXIT: {},
	ECHO: {},
	TYPE: {},
	PWD:  {},
}

type Run interface {
	run()
}

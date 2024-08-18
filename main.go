package main

import (
	"fmt"
	"os"
	cm "gamch1k.org/ggit/commands"
)

const (
    cmdInit   = "init"
    cmdAdd    = "add"
    cmdCommit = "commit"
	cmdHelp = "-h"
)

func main() {
	root_command := os.Args[1]

	if root_command == cmdInit {
		cm.InitRepo()
	} else if root_command == cmdAdd {
		cm.Add()
	} else if root_command == cmdCommit {

	} else if root_command == cmdHelp {
		cm.Help()
	} else {
		fmt.Println("Unknown command, try to use ggit -h")
	}
}

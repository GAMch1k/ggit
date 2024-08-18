package commands

import (
	"fmt"
)


func Help() {
	fmt.Println("Available commands:")
	fmt.Println("init - Create repository")
	fmt.Println("add - Add files to repository")
	fmt.Println("commit - Commit files")
	fmt.Println("-h - Show help")
}

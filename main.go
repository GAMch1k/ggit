package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"gamch1k.org/ggit/utils"
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
		initRepo()
	} else if root_command == cmdAdd {
		addCode()
	} else if root_command == cmdCommit {

	} else if root_command == cmdHelp {
		help()
	} else {
		fmt.Println("Unknown command, try to use ggit -h")
	}
}


func initRepo() {
	rootDir := utils.GetEnvVariable("ROOT_PATH")

	if utils.IsRepoExists() {
		fmt.Println("Repository already exists")
		return
	}
	
	os.Mkdir(rootDir, 0700)
	os.Mkdir(rootDir + "/commits", 0700)

	err := os.WriteFile(rootDir + "/.ggitignore", []byte(".ggit"), 0700)
	if utils.ErrorHandler(err) {
		return
	}

	fmt.Println("Created repository")
}


func addCode() {
	if !utils.IsRepoExists() {
		fmt.Println("Repository isn't created")
		return
	}

	files := []string{}
	err := filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
			if utils.ErrorHandler(err) {
				return err
			}

			if !d.IsDir() {
				files = append(files, path)
			}
			return nil
		})

	if utils.ErrorHandler(err) {
		return
	}

	for _, file := range files {
		if utils.IsFileInGgitignore(file) {
			continue
		}
		
		fmt.Println("Processing file ->", file)
	}
}



func help() {
	fmt.Println("Available commands:")
	fmt.Println("init - Create repository")
	fmt.Println("add - Add files to repository")
	fmt.Println("commit - Commit files")
	fmt.Println("-h - Show help")
}

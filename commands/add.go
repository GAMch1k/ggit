package commands

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"gamch1k.org/ggit/utils"
)

// This function is responsible for adding files to the staging area
// It first checks if the repository exists
// Then it walks through all files in the current directory and its subdirectories
// For each file, it checks if the file is in the .ggitignore file
// If the file is not in the .ggitignore, it processes the file (e.g., adds it to the staging area)

func Add() {
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

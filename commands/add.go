package commands

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"gamch1k.org/ggit/utils"
)

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

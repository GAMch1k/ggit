package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

func IsFolderExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func IsRepoExists() bool {
	rootDir := GetEnvVariable("ROOT_PATH")

	return IsFolderExists(rootDir)
}

func loadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func GetEnvVariable(key string) string {
	loadEnvVariables()
	
	value := os.Getenv(key)
	if value == "" {
		panic("Environment variable " + key + " is not set")
	}
	return value
}

func IsFileInGgitignore(file string) bool {
	ggitignorePath := GetEnvVariable("ROOT_PATH") + "/.ggitignore"

	if _, err := os.Stat(ggitignorePath); os.IsNotExist(err) {
		fmt.Println("Error: .ggitignore file not found")
		panic(err)
	}

	content, err := os.ReadFile(ggitignorePath)
	if err != nil {
		fmt.Println("Error: Unable to read .ggitignore file")
		panic(err)
	}

	
	ignored := strings.Split(strings.ReplaceAll(string(content), "\r\n", "\n"), "\n")

	for _, f := range ignored {
		if string(f) == "" || strings.HasPrefix(string(f), "#") {
			continue
		}

		if strings.Contains(string(f), "*") {
			f = strings.ReplaceAll(string(f), "*", "")
			if strings.Contains(file, string(f)) {
				return true
			}
		} else if strings.Contains(string(f), "/") {
			folderName := filepath.ToSlash(string(f))
			folderName = strings.ReplaceAll(folderName, "/", "")
			
			filePath := filepath.ToSlash(file)
			

			dirs := strings.Split(filePath, "/")

			dirs = dirs[:len(dirs)-1]

			if len(dirs) == 0 {
				continue
			}

			for _, dir := range dirs {
				if dir == folderName {
					return true
				}
			}
		} else {
			if file == string(f) {
				return true
			}
		}
	}

	return false
}

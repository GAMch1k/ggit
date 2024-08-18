package commands

import (
	"fmt"
	"os"

	"gamch1k.org/ggit/utils"
)


// InitRepo initializes a new repository in the current directory
// It creates a .ggitingore file and the necessary directories
// for storing commits and data


func InitRepo() {
	rootDir := utils.GetEnvVariable("ROOT_PATH")

	if utils.IsRepoExists() {
		fmt.Println("Repository already exists")
		return
	}
	
	os.Mkdir(rootDir, 0700)
	os.Mkdir(rootDir + "/commits", 0700)
	os.Mkdir(rootDir + "/database", 0700)

	utils.DownloadFile(rootDir + "/.ggitignore", "https://raw.githubusercontent.com/GAMch1k/ggit/master/.ggitignore.example")

	fmt.Println("Created repository")
}

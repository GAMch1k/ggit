package utils

import (
	"fmt"
	"os"
	"io"
	"net/http"
)

// DownloadFile downloads a file from the given URL and saves it to the specified filepath.

func DownloadFile(filepath string, url string) {
	fmt.Printf("Downloading %s -> %s\n", url, filepath)

	resp, err := http.Get(url)
	if ErrorHandler(err) {
		panic(err)
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if ErrorHandler(err) {
		panic(err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if ErrorHandler(err) {
		panic(err)
	}

	fmt.Println("Download complete!")
}

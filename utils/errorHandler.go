package utils


import (
	"fmt"
)


func ErrorHandler(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
		return true
	}
	return false
}
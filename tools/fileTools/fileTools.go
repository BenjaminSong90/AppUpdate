package fileTools

import (
	"os"
)

func PathExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

func MakeDir(path string) bool {
	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		return false
	} else {
		return true
	}
}
package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
)

func getSubjectConfigsFile() ([]fs.DirEntry, error) {
	wd, err := os.Getwd()
	if err != nil {

		return []fs.DirEntry{}, err
	}
	configPath := fmt.Sprintf("%s/configs", wd)

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return []fs.DirEntry{}, errors.New("configs folder doesnt exist in working dir, create one")
	}

	files, err := os.ReadDir(configPath)
	if err != nil {
		log.Println("Error reading directory:", err)
		return []fs.DirEntry{}, err
	}

	return files, nil
}

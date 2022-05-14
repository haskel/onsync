package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func fileFunc(path string, info os.FileInfo, fileErr error) error {
	if fileErr != nil {
		fmt.Println(fileErr)
		return fileErr
	}
	fmt.Printf("dir: %v: name: %s\n", info.IsDir(), path)
	return nil
}

func ListSubdirectories(path string) []string {
	directories := []string{path}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			continue
		}

		absoluteFilepath := filepath.Join(path, file.Name())
		directories = append(directories, ListSubdirectories(absoluteFilepath)...)
	}

	return directories
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "  ")
	return string(s)
}

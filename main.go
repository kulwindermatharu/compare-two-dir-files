package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var files []string
	var files1 []string
	//full dir
	root := "First Dir Path"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	//missing dir
	rootCheck := "Second Dir Path"
	err = filepath.Walk(rootCheck, func(path string, info os.FileInfo, err error) error {
		files1 = append(files1, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	justString := strings.Join(files1, ",")
	justString = justString + ","
	for _, file := range files {
		file = strings.Replace(file, "First Dir Path", "", -1)
		if !strings.Contains(justString, file+",") {
			fmt.Println(file)
		}

	}
}

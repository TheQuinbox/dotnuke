package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <path>\n", os.Args[0])
		os.Exit(1)
	}
	fileList := make([]string, 0)
	e := filepath.Walk(os.Args[1], func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return err
	})
	if e != nil {
		panic(e)
	}
	for _, file := range fileList {
		if strings.HasPrefix(filepath.Base(file), "._") {
			fmt.Printf("Found a file to delete: %s\n", filepath.Base(file))
		}
	}
}

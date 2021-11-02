package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <path>\n", os.Args[0])
		os.Exit(1)
	}
	fileList := make([]string, 0)
	e := filepath.Walk(os.args[1], func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return err
	})
	if e != nil {
		panic(e)
	}
	for _, file := range fileList {
		fmt.Println(file)
	}
}

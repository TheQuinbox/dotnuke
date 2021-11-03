package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var total, deleted, ignored int

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <path>\n", os.Args[0])
		os.Exit(1)
	}
	fileList := make([]string, 0)
	err := filepath.Walk(os.Args[1], func(path string, file os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return err
	})
	if err != nil {
		panic(err)
	}
	for _, file := range fileList {
		total += 1
		if strings.HasPrefix(filepath.Base(file), "._") {
			deleted += 1
			err := os.Remove(file)
			if err != nil {
				panic(err)
			}
		} else {
			ignored += 1
		}
	}
	fmt.Printf("Done! Out of %d %s, %d %s deleted, and %d %s kept.", total, plural(total, "file", "files"), deleted, plural(deleted, "was", "were"), ignored, plural(ignored, "was", "were"))
	os.Exit(0)
}

// Simple hacky little function for pluralization.
func plural(number int, first, second string) string {
	if number == 1 || number == -1 {
		return first
	} else {
		return second
	}
}

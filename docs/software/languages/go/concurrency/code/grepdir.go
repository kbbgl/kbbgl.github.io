package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func printFileContainsPattern(filename string, pattern string) {

	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	if strings.Contains(string(file), pattern) {
		fmt.Printf("File %s contains pattern '%s'\n", filename, pattern)
	}
}

//func visit(path string, f fs.FileInfo, err error) error {
//	if !f.IsDir() {
//
//	}
//}

func main() {

	args := os.Args[1:]

	stringToSearch := args[0]

	if len(stringToSearch) == 0 {
		log.Fatal("No search string")
	}

	dirToSearch := args[1]

	fmt.Printf("Searching directory %s for pattern '%s'...\n", dirToSearch, stringToSearch)

	err := filepath.WalkDir(dirToSearch, func(path string, d fs.DirEntry, err error) error {

		if err != nil {
			fmt.Printf("Error accessing path %s: %v\n", path, err)
			return err
		}

		if !d.IsDir() {
			absPath, err := filepath.Abs(path)
			if err != nil {
				return err
			}

			go printFileContainsPattern(absPath, stringToSearch)

		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error accessing directory %s: %v\n", dirToSearch, err)
	}

	time.Sleep(10 * time.Second)
}

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func printFileContainsPattern(filename string, pattern string) {

	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	if strings.Contains(string(file), pattern) {
		fmt.Printf("File %s contains pattern %s\n", filename, pattern)
	}
}

func main() {

	args := os.Args[1:]

	patternToMatch := args[0]
	filesToMatch := args[1:]

	for _, f := range filesToMatch {
		go printFileContainsPattern(f, patternToMatch)
	}

	time.Sleep(10 * time.Second)

}

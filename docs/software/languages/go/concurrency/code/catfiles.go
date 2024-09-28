package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func printFile(filename string) {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(string(file))
}

func main() {

	args := os.Args[1:]

	log.Printf("Number of arguments: %d", len(args))

	for _, arg := range args {
		go printFile(arg)
	}

	time.Sleep(10 * time.Second)

}

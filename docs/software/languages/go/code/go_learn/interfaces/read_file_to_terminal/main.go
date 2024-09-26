package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args[1]

	data, err := os.Open(args)

	if err != nil {
		fmt.Println("Error opening file: ", err)
		os.Exit(1)
	}

	fmt.Println(data)

	io.Copy(os.Stdout, data)
}

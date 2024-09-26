package main

import (
	"fmt"
	"io"
	"net/http"
)

type logWriter struct{}

func main() {

	resp, err := http.Get("https://httpbin.org/anything")

	if err != nil {
		print(err)
	}

	lw := logWriter{}

	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {

	fmt.Println(string(bs))

	return len(bs), nil

}

package main

import (
	"io"
	"net/http"
	"os"
)

func main() {

	resp, err := http.Get("https://httpbin.org/anything")

	if err != nil {
		print(err)
	}

	io.Copy(os.Stdout, resp.Body)

}

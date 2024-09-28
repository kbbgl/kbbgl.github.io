package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func countWords(url string, frequency map[string]int) {
	response, err := http.Get(url)
	defer response.Body.Close()

	if response.StatusCode != 200 || err != nil{
		panic("Server returning error status code: " + response.Status)
	}

	body, err := io.ReadAll(response.Body)

	words := strings.Fields(string(body))

	for _, word := range words {
		word = strings.ToLower(strings.Trim(word, ".,"))
		frequency[word]++
	}

	fmt.Println("Completed:", url)
}

func main()  {

	var frequency = make(map[string]int)

	for i := 1000; i < 1030; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		countWords(url, frequency)
	}

	for word, f := range frequency {
		fmt.Printf("%s -> %d\n", word, f)
	}

}

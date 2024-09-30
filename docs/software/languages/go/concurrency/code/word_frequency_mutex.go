package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"
)

func countWordsConcurrently(url string, frequency map[string]int, mutex *sync.Mutex) {
	response, err := http.Get(url)
	defer response.Body.Close()

	if response.StatusCode != 200 || err != nil{
		panic("Server returning error status code: " + response.Status)
	}

	body, err := io.ReadAll(response.Body)

	wordRegex := regexp.MustCompile(`[a-zA-Z]+`)

	mutex.Lock()
	for _, word := range wordRegex.FindAllString(string(body), -1) {
		wordLower := strings.ToLower(word)
		frequency[wordLower]++
	}
	mutex.Unlock()
	fmt.Println("Completed:", url)
}

func main()  {

	var frequency = make(map[string]int)
	mutex := sync.Mutex{}

	for i := 1000; i < 1030; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		go countWordsConcurrently(url, frequency, &mutex)
	}

	time.Sleep(10 * time.Second)
	mutex.Lock()
	for word, f := range frequency {
		fmt.Printf("%s -> %d\n", word, f)
	}
	mutex.Unlock()
}

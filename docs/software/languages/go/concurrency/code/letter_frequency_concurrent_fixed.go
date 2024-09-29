package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

func getAlphabet() (string) {
	alphabet := ""
	for i := 'a';i <= 'z';i++ {
		alphabet += string(i)
	}
	return alphabet
}


func countLetters(url string, frequency []int, mutex *sync.Mutex) {
	response, err := http.Get(url)
	defer response.Body.Close()

	if response.StatusCode != 200 || err != nil{
		panic("Server returning error status code: " + response.Status)
	}

	body, err := io.ReadAll(response.Body)

	alphabet := getAlphabet()
	
	mutex.Lock()
	for _, b := range body {
		c := strings.ToLower(string(b))
		cIndex := strings.Index(alphabet, c)
		if cIndex >= 0 {
			frequency[cIndex] += 1
		}
	}
	mutex.Unlock()

	fmt.Println("Completed:", url)
}

func main()  {

	var frequency = make([]int, 26)
	mutex := sync.Mutex{}

	for i := 1000; i < 1030; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		go countLetters(url, frequency, &mutex)
	}

	time.Sleep(10 * time.Second)

	mutex.Lock()
	for i, c := range getAlphabet() {
		fmt.Printf("%c-%d\n", c, frequency[i])
	}
	mutex.Unlock()
}

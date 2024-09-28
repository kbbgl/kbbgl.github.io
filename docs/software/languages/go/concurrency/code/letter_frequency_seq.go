package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getAlphabet() (string) {
	alphabet := ""
	for i := 'a';i <= 'z';i++ {
		alphabet += string(i)
	}
	return alphabet
}


func countLetters(url string, frequency []int) {
	response, err := http.Get(url)
	defer response.Body.Close()

	if response.StatusCode != 200 || err != nil{
		panic("Server returning error status code: " + response.Status)
	}

	body, err := io.ReadAll(response.Body)

	alphabet := getAlphabet()
	
	for _, b := range body {
		c := strings.ToLower(string(b))
		cIndex := strings.Index(alphabet, c)
		if cIndex >= 0 {
			frequency[cIndex] += 1
		}
	}

	fmt.Println("Completed:", url)
}

func main()  {

	var frequency = make([]int, 26)

	for i := 1000; i < 1030; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		countLetters(url, frequency)
	}

	for i, c := range getAlphabet() {
		fmt.Printf("%c-%d\n", c, frequency[i])
	}

}

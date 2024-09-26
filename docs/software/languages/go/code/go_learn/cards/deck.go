package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"♤", "♧", "♢", "♡"}
	cardValues := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	// if we don't need the index i/j,
	// we replace them with an underscore
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(cards deck, handSize int) (deck, deck) {
	return cards[:handSize], cards[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	byteslice, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
		// return newDeck()
		os.Exit(1)
	}

	sliceOfStrings := strings.Split(string(byteslice), ",")

	return deck(sliceOfStrings)
}

func (d deck) shuffle() {
	for index := range d {
		rand.Seed(time.Now().UnixNano())
		max := len(d) - 1
		newRandomIndex := rand.Intn(max + 1)

		// Swap
		d[index], d[newRandomIndex] = d[newRandomIndex], d[index]

	}
}

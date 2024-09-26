package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length is 52, but got %d", len(d))
	}
}

func TestFirstDeckCard(t *testing.T) {

	d := newDeck()
	expected := "A♤"

	actual := d[0]

	if actual != expected {
		t.Errorf("Expected first card to be %v, got %v", expected, actual)
	}
}

func TestLastDeckCard(t *testing.T) {

	d := newDeck()
	expected := "K♡"

	actual := d[len(d)-1]

	if actual != expected {
		t.Errorf("Expected first card to be %v, got %v", expected, actual)
	}

}

func TestSaveAndLoadDeckFromFilesystem(t *testing.T) {

	testingFilenames := "_decktesting"

	os.Remove(testingFilenames)

	deck := newDeck()

	deck.saveToFile(testingFilenames)

	loadedDeck := newDeckFromFile(testingFilenames)

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length is 52, but got %d", len(loadedDeck))
	}

	os.Remove(testingFilenames)

}

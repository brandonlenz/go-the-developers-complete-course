package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	verifyDefaultDeck(d, t)
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// Before
	filename := "_decktesting"
	os.Remove(filename)

	// Given a saved deck
	d := newDeck()
	d.saveToFile(filename)

	// When loading from said file
	loadedDeck := newDeckFromFile(filename)

	// Then the loaded deck should pass basic data verifications
	verifyDefaultDeck(loadedDeck, t)

	// Cleanup
	os.Remove(filename)
}

func verifyDefaultDeck(d deck, t *testing.T) {
	deckLength := len(d)
	expectedLength := 52
	if deckLength != expectedLength {
		t.Errorf("Expected deck to contain %d cards, but got %d", expectedLength, deckLength)
	}

	firstCard := d[0]
	expectedFirstCard := "Ace of Spades"
	if firstCard != expectedFirstCard {
		t.Errorf("Expected first card to be %s, but got %s", expectedFirstCard, firstCard)
	}

	lastCard := d[len(d)-1]
	expectedLastCard := "King of Clubs"
	if lastCard != expectedLastCard {
		t.Errorf("Expected last card to be %s, but got %s", expectedLastCard, lastCard)
	}
}

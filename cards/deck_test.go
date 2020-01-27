package main

import (
	"os"
	"testing"
)

var (
	d = newDeck()
	hand, remainingDeck = deal(d, 5)
)

func TestNewDeck(t *testing.T) {
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Diamonds" {
		t.Errorf("Expected first card in deck to be 'Ace of Diamonds', but got '%v'", d[0])
	}

	if d[len(d) -1] != "King of Clubs" {
		t.Errorf("Expected last card in deck to be 'King of Clubs', but got '%v'", d[len(d) -1])
	}
}

func TestDeal(t *testing.T) {
	if len(hand) != 5 {
		t.Errorf("Expected hand size to be 5, but got %v", len(hand))
	}

	if len(remainingDeck) != 47 {
		t.Errorf("Expected remaining deck size to be 47, but got %v", len(remainingDeck))
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected loaded deck length of 52, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
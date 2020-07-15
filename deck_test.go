package main

import (
	"fmt"
	"os"
	"testing"
)

// Tests tbd
func TestNewDeck(t *testing.T) {
	cards := newDeck()
	if len(cards) != 40 {
		t.Errorf("Expected deck length of 40, but got %v", len(cards))
	}

	if cards[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades but got %v", cards[0])
	}

	if cards[len(cards)-1] != "Ten of Clubs" {
		t.Errorf("Expected last card of Ten of Clubs but got %v", cards[len(cards)-1])
	}
}

func TestNewDeckFromFile(t *testing.T) {

	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 40 {
		t.Errorf("Expected 40 cards in deck, got %v", len(loadedDeck))
	}

	err := os.Remove("_decktesting")
	if err != nil {
		fmt.Println("Error removing deck")
	}
}

package main

import (
	"fmt"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected the deck size to be 52, but received %v", len(d))
	}

	if d[0] != "Ace of Hearts" {
		t.Errorf("Expected the first card to be Ace of Hearts, but got %v", d[0])
	}
	if d[len(d)-1] != "King of Spades" {
		t.Errorf("Expected the last card to be King of Spades, but got %v", d[len(d)-1])
	}
}

func TestReadFileSaveToFile(t *testing.T) {
	fp := "_deck.csv"
	os.Remove(fp)

	d := newDeck()
	err := d.saveToFile(fp, ",", 0644)
	if err != nil {
		fmt.Println("Error:", err)
	}

	d = readFile(fp, ",")

	if len(d) != 52 {
		t.Errorf("Expected the deck size to be 52, but received %v", len(d))
	}

	os.Remove(fp)
}

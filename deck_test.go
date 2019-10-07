// Testing with go does not involve using an external testing framework
// go has its own native testing mechanism
// to make a new test, just make a new file with naming convention <package>_test.go
package main

import (
	"os"
	"testing"
)

// deciding what to test
func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52 but got %v", len(d))
	}
	if d[0].toString() != "2 of Spades" {
		t.Errorf("Expected first card of 2 of Spades got %v", d[0])
	}
	if d[len(d)-1].toString() != "A of Clubs" {
		t.Errorf("Expected first card of A of Clubs got %v", d[len(d)-1])
	}
}

func TestSaveDeckToFileAndNewDeckFromFile(t *testing.T) {
	// must setup and tear down filesystem state
	filename := "_desk_test"
	os.Remove(filename)
	d := newDeck()
	d.saveToFile(filename)

	loadedDeck := newDeckFromFile(filename)
	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52 but got %v", len(loadedDeck))
	}
	// tear down
	os.Remove(filename)
}

func TestCardToString(t *testing.T) {
	c := card{"A", "Spades"}
	cStr := c.toString()
	if cStr != "A of Spades" {
		t.Errorf("Expected card.toString() to be 'A of Spades' got %v", cStr)
	}
}

package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of deck which is a slice of strings
type deck []string

// golang equivalent to a method
// Func with a declared receiver type
// any variable of type deck has access to the print method
// d is the type instance (convention is a 1-2 letter abbreviation of the type)
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	// suitChars := []string{"U+2660", "U+2665", "U+2663", "U+2666"}
	fmt.Println("\n...Creating new deck")
	cards := deck{}
	suits := suits()
	values := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	for _, suit := range suits {
		for _, value := range values {
			card := value + " of " + suit
			cards = append(cards, card)
		}
	}
	return cards

}

func (d deck) shuffle() deck {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	// rand.Seed(42)
	for i := range d {
		randIdx := r.Intn(len(d) - 1)
		// Fancy! no temporary variable necessary!
		d[i], d[randIdx] = d[randIdx], d[i]
	}
	return d
}

func suits() []string {
	return []string{"Spades", "Hearts", "Diamonds", "Clubs"}
}

func deal(d deck, handSize int) (deck, deck, error) {
	if handSize > len(d) {
		return nil, nil, fmt.Errorf("Error: deal() - deck has less than %v cards", handSize)
	}
	return d[:handSize], d[handSize:], nil
}

func (d deck) toString() string {
	return strings.Join([]string(d), "\n")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// option 1: log error and fall back to a call to newDeck()
		// option 2: log error and quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	str := string(bs)
	strs := strings.Split(str, "\n")
	return deck(strs)
}

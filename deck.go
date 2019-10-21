package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var cardValues = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

var cardSuits = []string{"Spades", "Hearts", "Diamonds", "Clubs"}

// Create a new struct for cards
type card struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
}

func (c card) toString() string {
	// Check if value can be cast to a number
	return c.Value + " of " + c.Suit
}

func (c card) isFace() bool {
	// Check if value can be cast to a number
	if _, err := strconv.Atoi(c.Value); err == nil {
		return false
	}
	return true
}

func (c card) trumps(otherCard card) bool {
	// Check if value can be cast to a number
	trumps := false
	idx := Index(cardValues, c.Value)
	otherIdx := Index(cardValues, otherCard.Value)
	if idx > otherIdx {
		trumps = true
	}
	return trumps
}

//Deck : Create a new type of Deck which is a slice of strings
type Deck []card

// golang equivalent to a method
// Func with a declared receiver type
// any variable of type Deck has access to the print method
// d is the type instance (convention is a 1-2 letter abbreviation of the type)
func (d Deck) print() {
	for i, card := range d {
		fmt.Println(i, card.toString())
	}
}

func newDeck() Deck {
	// suitChars := []string{"U+2660", "U+2665", "U+2663", "U+2666"}
	fmt.Println("\n...Creating new Deck")
	cards := Deck{}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			thisCard := card{value, suit}
			cards = append(cards, thisCard)
		}
	}
	return cards

}

func (d Deck) shuffle() Deck {
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

func deal(d Deck, handSize int) (Deck, Deck, error) {
	if handSize > len(d) {
		return nil, nil, fmt.Errorf("Error: deal() - Deck has less than %v cards", handSize)
	}
	return d[:handSize], d[handSize:], nil
}

func (d Deck) toString() string {
	cardStrings := []string{}
	for _, card := range d {
		cardStrings = append(cardStrings, card.toString())
	}
	return strings.Join(cardStrings, "\n")
}

func (d Deck) getSuit(s string) Deck {
	results := Deck{}
	for _, card := range d {
		if strings.ToLower(card.Suit) == strings.ToLower(s) {
			results = append(results, card)
		}
	}
	return results
}

func (d Deck) saveToFile(filename string) error {
	jsonArr, err := json.MarshalIndent(d, "", "    ")
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}
	return ioutil.WriteFile(filename, jsonArr, 0666)
}

func newDeckFromFile(filename string) Deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// option 1: log error and fall back to a call to newDeck()
		// option 2: log error and quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	var cards Deck
	json.Unmarshal(bs, &cards)
	return cards
}

// Hand : hand of cards
type Hand struct {
	Cards     Deck      `json:"cards,omitempty"`
	TimeStamp time.Time `json:"timeStamp"`
	Size      int       `json:"size"`
}

type hands []Hand

// ByValue : custom collection sort for slice of cards
type ByValue []card

func (v ByValue) Len() int           { return len(v) }
func (v ByValue) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }
func (v ByValue) Less(i, j int) bool { return v[i].Value < v[j].Value }

func (d Deck) sortByValue() {
	sort.Stable(ByValue(d))
}

// BySuit : custom collection sort for slice of cards
type BySuit []card

func (s BySuit) Len() int           { return len(s) }
func (s BySuit) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s BySuit) Less(i, j int) bool { return s[i].Suit < s[j].Suit }

func (d Deck) sortBySuit() {
	sort.Stable(BySuit(d))
}

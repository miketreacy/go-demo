package main

import (
	"fmt"
	"os"
)

// variable declarations outside of funcs must be explicitly typed
// var deckSize int = 20

func main() {
	// var card string = "Ace of Spades"
	// := is only for initialization inside of funcs
	// compiler infers the type
	// deckSize = 52
	// fmt.Println(card)

	// a new slice of type string
	// cards := deck{newCard(), "Ace of Diamonds"}
	// fmt.Println(cards)
	// append does not modify existing slice, returns new slice
	// cards = append(cards, "Six of Spades")

	// print command-line arguments
	var args = os.Args[1:]
	fmt.Println("Command line arguments:")
	fmt.Println(args)

	// make http request
	url := "https://motivic.io"
	statusCode, status, content := fetch(url)
	fmt.Println("response:", statusCode, status)
	fmt.Println(content)

	// do card stuff
	fmt.Println("\n...getting deck")
	cards := newDeck()

	fmt.Println("...shuffling deck")
	cards.shuffle()

	fmt.Println("...saving deck to file")
	cards.saveToFile("my_cards.txt")

	fmt.Println("...loading deck from file")
	cards = newDeckFromFile("my_cards.txt")

	// run the server
	fmt.Println("...spinning up the server")
	serve(cards)

}

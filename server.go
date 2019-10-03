package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

var cardDeck deck
var port string = "8080"

func serve(cards deck) {
	cardDeck = cards
	http.HandleFunc("/", helloWeb)
	http.HandleFunc("/deal/", dealHand)
	fmt.Println("...listening at localhost:" + port)
	http.ListenAndServe(":8080", nil)
}

func helloWeb(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

type cardlooper func(card string) (ok bool)

func suitCheck(suit string) cardlooper {
	return func(card string) bool {
		cardSuit := strings.Split(card, " ")[2]
		return cardSuit == suit
	}
}

func dealHand(w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path, "/")
	handSizeStr := paths[2]
	handSize, _ := strconv.Atoi(handSizeStr)
	hand, cards := deal(cardDeck, handSize)
	cardDeck = cards

	fmt.Fprintf(w, "You have been dealt a hand of %v cards!\n", strconv.Itoa(len(hand)))
	fmt.Fprintf(w, "\n%v\n", hand.toString())
	fmt.Fprintf(w, "\n%v cards left in deck:\n", strconv.Itoa(len(cardDeck)))
	for _, s := range suits() {
		cards := filter([]string(cardDeck), suitCheck(s))
		fmt.Fprintf(w, "%v %v\n", strconv.Itoa(len(cards)), s)

	}
}

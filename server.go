package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var cardDeck deck
var port string = "8080"

func serve(cards deck) {
	cardDeck = cards
	http.HandleFunc("/", helloWeb)
	http.HandleFunc("/deal/", dealHand)
	http.HandleFunc("/api/", apiHandler)
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
	hand, cards, err := deal(cardDeck, handSize)
	if err != nil {
		fmt.Fprintf(os.Stderr, "deal: %v\n", err)
		cardDeck = newDeck()
		hand, cards, err = deal(cardDeck, handSize)
	}
	cardDeck = cards

	fmt.Fprintf(w, "You have been dealt a hand of %v cards!\n", strconv.Itoa(len(hand)))
	fmt.Fprintf(w, "\n%v\n", hand.toString())
	fmt.Fprintf(w, "\n%v cards left in deck:\n", strconv.Itoa(len(cardDeck)))
	for _, s := range suits() {
		cards := filter([]string(cardDeck), suitCheck(s))
		fmt.Fprintf(w, "%v %v\n", strconv.Itoa(len(cards)), s)

	}
}

type hand struct {
	Cards     deck      `json:"cards"`
	TimeStamp time.Time `json:"timeStamp"`
	Size      int       `json:"size"`
}

type hands []hand

// basic JSON API
func apiHandler(w http.ResponseWriter, r *http.Request) {
	ts := time.Now()
	h := hand{Cards: cardDeck, TimeStamp: ts, Size: len(cardDeck)}
	data, err := json.MarshalIndent(h, "", "    ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "apiHandler: json.MarshallIndent %v: %v\n", h, err)
		os.Exit(1)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)

}

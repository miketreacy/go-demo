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

// CardDeck : application state variable that holds a deck instance persisted in memory
var CardDeck Deck
var port string = "8080"

// TODO: declare a routeMap config for server that can be exported to callers

func serve(cards Deck) {
	CardDeck = cards
	http.HandleFunc("/", helloWeb)
	http.HandleFunc("/deal/", dealHand)
	http.HandleFunc("/api/", apiHandler)
	http.HandleFunc("/scrape/", scraper)
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
	hand, cards, err := deal(CardDeck, handSize)
	if err != nil {
		fmt.Fprintf(os.Stderr, "deal: %v\n", err)
		CardDeck = newDeck()
		hand, cards, err = deal(CardDeck, handSize)
	}
	CardDeck = cards

	fmt.Fprintf(w, "You have been dealt a hand of %v cards!\n", strconv.Itoa(len(hand)))
	fmt.Fprintf(w, "\n%v\n", hand.toString())
	fmt.Fprintf(w, "\n%v cards left in deck:\n", strconv.Itoa(len(CardDeck)))
	for _, s := range cardSuits {
		cards := CardDeck.getSuit(s)
		fmt.Fprintf(w, "%v %v\n", strconv.Itoa(len(cards)), s)

	}
}

// basic JSON API
func apiHandler(w http.ResponseWriter, r *http.Request) {
	ts := time.Now()
	h := Hand{Cards: CardDeck, TimeStamp: ts, Size: len(CardDeck)}
	var jsonData []byte
	jsonData, err := json.MarshalIndent(h, "", "    ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "apiHandler: json.MarshallIndent %v: %v\n", h, err)
		os.Exit(1)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func scraper(w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path, "/")
	url := paths[2]
	_, _, html, err := fetch("https://"+url, 9999999)
	if err != nil {
		fmt.Fprintf(w, "Can't fetch %v", err)
	}
	fmt.Fprintf(w, html)
}

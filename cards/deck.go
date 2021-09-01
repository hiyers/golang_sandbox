package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

// Create a new type 'deck'
// of type slice of strings
type deck []string

// d deck is reciever. Attaches method to type deck. cards variable gets passed to this function
// usually first letter of deck, i.e. d is used as name of reciever as per the convention
func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}
}

// you don't yet have a deck, so reciever not required. you are creating a new deck
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards

}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {

	return strings.Join([]string(d), ",") // convert back as deck itself is slice of string
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDecFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option #1  - log the error and return the call to new Deck
		// Option #2 - log the error and entirely quit the program
		fmt.Println("Error", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",") // Ace of Spades, Two of Spades, ...
	return deck(s)
}

func (d deck) shuffle() {
	for i := range d {
		newPosition := rand.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i] // one line swap
	}
}

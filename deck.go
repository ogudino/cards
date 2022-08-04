package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// create a new type of 'deck'
// which is a slice of strings
type deck []string

// create and return a list of playing cards
// Essentially an array of strings
func newDeck() deck { //return value deck
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit) //created new card in cards slice
		}
	}

	return cards
}

func (d deck) print() { //reciever
	for i, card := range d { //for index of this element in the array, current d we're iterationg over, take the slice of 'cards' and loop over it
		fmt.Println(i, card) //run this one time for each card in the slice
	}
}

func deal(d deck, handSize int) (deck, deck) { //arguement deck and arguement handsize type int. Return two values of type deck
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string { //reciever of type deck
	return strings.Join([]string(d), ",")

}

func (d deck) saveToFile(filename string) error { //function to save string into .txt file
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)

}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

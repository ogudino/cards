package main

import "fmt"

//create a new type of 'deck'
//which is a slice of strings
type deck []string

//create and return a list of playing cards
//Essentially an array of strings
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

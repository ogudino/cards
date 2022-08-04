package main

import "fmt"

//create a new type of 'deck'
//which is a slice of strings
type deck []string

func (d deck) print() {
	for i, card := range d { //for index of this element in the array, current card we're iterationg over, take the slice of 'cards' and loop over it
		fmt.Println(i, card) //run this one time for each card in the slice
	}
}

package main

import "fmt"

func main() {
	// cards := newDeck()

	// hand, remainingCards := deal(cards, 5) //return two values

	// hand.print()
	// remainingCards.print()
	cards := newDeck()
	fmt.Println(cards.toString())
}

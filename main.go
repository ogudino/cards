package main

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 5) //return two values

	hand.print()
	remainingCards.print()
}

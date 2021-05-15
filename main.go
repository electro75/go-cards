package main

import "fmt"

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	fmt.Println("----------------")
	remainingCards.print()
	fmt.Println(cards.toString())

	cards.saveToFile("myCards")
}

func newCard() string {
	return "Five of Diamonds"
}

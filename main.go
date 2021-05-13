package main

import "fmt"

func main() {
	// card := "Ace of Spades" // := is used to a declare and initialise a new variable.
	card := newCard()

	fmt.Println(card)
}

func newCard() string { // since we return a string value from this function.
	return "Five of Diamonds"
}

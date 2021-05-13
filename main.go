package main

import "fmt"

func main() {
	cards := []string{newCard(), "Ace of Spades"}
	cards = append(cards, "King of Hearts")

	for i, card := range cards {
		fmt.Println(i, card)
	}

	// fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}

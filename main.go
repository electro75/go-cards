package main

func main() {
	cards := deck{}
	cards = newDeck()

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}

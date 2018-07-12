package main

func main() {
	cards := deck{"Ace of Spades", "Ace of Diamonds"}
	// append returns a new slice of cards with the new value
	cards = append(cards, "Two of Spades")
	cards.print()
}

func newCard() string {
	return "Ace of Spades"
}

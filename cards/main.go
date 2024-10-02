package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	cards := newDeck()
	cards.saveToFile("my_cards")
	cards.shuffle()
	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()
}

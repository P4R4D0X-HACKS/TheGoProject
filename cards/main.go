package main




func main() {
	// var card string = "Ace of spade"
	// card := "Ace of spade" // Only when we are defining a new variable not reassigning
	// card := newCard()
	// fmt.Println(card)
	// for i, card := range cards{
	// 	fmt.Println(i, card)
	// }

	// cards := newDeck()

	// hand , remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// cards := newDeck()
	// cards.saveToFile("my_cards")

	// card := newDeckFromFile("my_cards")
	// card.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}

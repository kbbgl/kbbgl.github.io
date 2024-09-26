package main

func main() {

	cards := newDeck()

	// cards.print()

	// hand, remainingDeck := deal(cards, 5)

	// fmt.Println("Hand")
	// hand.print()
	// fmt.Println(hand.toString())

	// // fmt.Println(("Remaining Deck"))
	// remainingDeck.print()

	println("After shuffle")

	cards.shuffle()

	cards.print()

}

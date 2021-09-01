package main

func main() {
	//var card string = "Ace of Spade"
	// card := "Ace of Spade"
	// card = "Five of Diamonds"
	// card = newCard()
	// cards := []string{"Ace of Diammonds", newCard()}
	// cards := deck{"Ace of Diammonds", newCard()}
	// cards = append(cards, "Ace of Spades")
	// for _, card := range cards {
	// 	fmt.Println(card)
	// }
	// cards := newDeck()
	// fmt.Println(cards.toString())
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// fmt.Println("-----------")
	// remainingCards.print()
	// cards.saveToFile("my_cards")
	cards := newDecFromFile("my_cards")
	// cards.shuffle()
	cards.print()
}

func newCard() string {
	return "Six of Diamonds"
}

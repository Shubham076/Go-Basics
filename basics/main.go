package main
func main() {
	cards := newDeck()		
	// cards.print();
	// cards.saveToFile("deck.txt")
	// newcards := loadFile("deck.txt")
	// newcards.print()
	cards.shuffle()
	cards.print()
	// hand, rem := deal(cards, 10);
}

package main

func main() {

	cards := deck{"Five of Diamonds", "Ace of Spades"}
	cards = append(cards, "Six of Spades")

	cards.print()
}

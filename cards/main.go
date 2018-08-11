package main

//[]string    -a slice of type string
func main() {
	cards := deck{"salam", newCard()}
	cards = append(cards, "Nou element adaugat")

	cards.numeFunctie()
}
func newCard() string {
	return "Five of Diamonds"
}

package main

import "fmt"

func main() {
	deck := newDeck()
	deck.saveToFile("data.csv", ",", 0644)
	deck = readFile("data.csv", ",")
	fmt.Println("Before:", deck)
	deck = deck.shuffle()
	fmt.Println("After:", deck)
}

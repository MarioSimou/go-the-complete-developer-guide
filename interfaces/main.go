package main

import "fmt"

func main() {
	eb := englishBot{language: "english"}
	sb := spanishBot{language: "spanish"}

	printGreeting(eb)
	printGreeting(sb)
}

type bot interface {
	getGreeting() string
}

type englishBot struct{ language string }
type spanishBot struct{ language string }

func (eb englishBot) getGreeting() string {
	return "hello there"
}

func (sb spanishBot) getGreeting() string {
	return "holla"
}
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

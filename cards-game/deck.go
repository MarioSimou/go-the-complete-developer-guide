package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

/*
	Description: routine that generates a deck of cards

	input: -
	output: deck
*/
func newDeck() deck {
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	values := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	d := deck{}

	for _, s := range suits {
		for _, v := range values {
			d = append(d, strings.Join([]string{v, "of", s}, " "))
		}
	}
	return d
}

/*
	Description: routine that read a file from a specified location and
	returns a deck of cards

	input:
		filename: the file location
		sep: separator value used to write the file
	output:
		deck: deck of cards
*/
func readFile(filename string, sep string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
	}

	return deck(strings.Split(string(bs), sep))
}

/*
	Description: Receiver routine that print each card from the deck
	input: -
	output: -
*/
func (d deck) print() {
	for i, c := range d {
		fmt.Println(i, c)
	}
}

/*
	Description: Receiver routine that return n cards from the deck,
	as well as the modified deck (although in reality is a new instance of the deck)

	input:
		n = number of cards that should be returned
	output:
		hand = n returned cards
		dekc = the deck without the n returned cards
*/
func (d deck) deal(n int) (deck, deck) {
	return d[:n], d[n:]
}

/*
	Description: Function that returns a bytes representation of the deck

	input:
		sep = the separator value
	output:
		bs = a slice of bytes with the converted deck
*/
func (d deck) toBytes(sep string) []byte {
	return []byte(strings.Join(d, sep))
}

/*
	Description: Function that stores the deck in the filesystem

	input:
		filename = the name of the file to be stored
		sep = the separator value for the deck
		perm = default permissions of the file
	outpu:
		err = an error message if the WriteFile routine fails
*/
func (d deck) saveToFile(filename string, sep string, perm os.FileMode) error {
	bs := d.toBytes(sep)
	return ioutil.WriteFile(filename, bs, perm)
}

/*
	Description: Receiver function that shuffles the deck of cards

	input: -
	output: deck = shuffled deck of cards
*/
func (d deck) shuffle() deck {
	rand.Seed(time.Now().UnixNano())
	p := rand.Perm(len(d))
	nd := deck{}

	for i := 0; i < len(d); i++ {
		nd = append(nd, d[p[i]])
	}

	return nd
}

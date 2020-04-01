package main

import (
	"database/sql"
	"fmt"
)

// Card is a struct used to store card information
type Card struct {
	id   int
	name string
	set  string
}

var database *sql.DB
var cardMap map[int]Card

func main() {
	cardMap = make(map[int]Card)
	var err error

	database, err = sql.Open("sqlite3", "./cards.db")
	if err != nil {
		fmt.Println("error opening file")
		fmt.Println(err)
		return
	}

	connect()
	GetCards()
	printCards()

	card := Card{
		id:   2,
		name: "Handheld Blaster",
		set:  "Wave 2",
	}

	AddCard(&card)

	GetCards()
	printCards()
}

func printCards() {
	for _, card := range cardMap {
		fmt.Printf("%d %s %s\n", card.id, card.name, card.set)
	}
}

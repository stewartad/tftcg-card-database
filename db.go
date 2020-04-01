package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// AddCard adds a card to the database
func AddCard(card *Card) {
	statement, _ := database.Prepare("INSERT INTO cards (cardname, cardset) VALUES (?, ?)")
	statement.Exec(card.name, card.set)
}

// GetCards populates the given map with all cards from the database
func GetCards() {
	rows, err := database.Query("SELECT id, cardname, cardset FROM cards")
	if err != nil {
		fmt.Println(err)
		return
	}
	var id int
	var name string
	var set string
	for rows.Next() {
		rows.Scan(&id, &name, &set)
		cardMap[id] = Card{id: id, name: name, set: set}
	}
}

func connect() {
	var err error
	var statement *sql.Stmt

	statement, err = database.Prepare("CREATE TABLE IF NOT EXISTS cards (id INTEGER PRIMARY KEY, cardname TEXT, cardset TEXT)")
	if err != nil {
		fmt.Println("error preparing statement")
		fmt.Println(err)
		return
	}

	statement.Exec()
	statement.Close()
	// statement, _ = database.Prepare("INSERT INTO cards (cardname, cardset) VALUES (?, ?)")
	// statement.Exec("Improvised Shield", "Wave 1")

	// rows, _ := database.Query("SELECT id, cardname, cardset FROM cards")
	// var id int
	// var name string
	// var set string
	// for rows.Next() {
	// 	rows.Scan(&id, &name, &set)
	// 	fmt.Println(strconv.Itoa(id) + ": " + name + " " + set)
	// }
}

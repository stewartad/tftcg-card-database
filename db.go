package main

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// AddCard adds a card to the database
func (db databaseInfo) AddCard(card *Card) {
	db.stmts[addCard].Exec(card.name, card.set)
}

func (db databaseInfo) GetCard(id int) {

}

// GetAllCards populates the global card map with all cards from the database
func (db databaseInfo) GetAllCards() {
	rows, err := db.stmts[getAllCards].Query()
	if err != nil {
		log.Fatal(err)
		return
	}

	var id int
	var name string
	var set string

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name, &set)
		if err != nil {
			log.Fatal(err)
		}
		cardMap[id] = Card{id: id, name: name, set: set}
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

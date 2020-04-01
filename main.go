package main

import (
	"database/sql"
	"fmt"
	"log"
)

const cardStmtAmount = 2
const (
	addCard     = iota
	getAllCards = iota
)

// Card is a struct used to store card information
type Card struct {
	id   int
	name string
	set  string
}

// a struct to hold the database object and prepared statements
type databaseInfo struct {
	db    *sql.DB
	stmts [cardStmtAmount]*sql.Stmt
}

var cardMap map[int]Card

// this will probably stop being in main and will instead be a subpackage of the main application
func main() {
	// probably will move this to a struct so we can have a map for each card type
	cardMap = make(map[int]Card)
	var err error
	var dbInfo databaseInfo

	// open the database
	dbInfo.db, err = sql.Open("sqlite3", "./cards.db")
	if err != nil {
		log.Fatal(err)
		return
	}

	// prepare statements
	addCardStmt, err := dbInfo.db.Prepare("INSERT INTO cards (cardname, cardset) VALUES (?, ?)")
	getAllStmt, err := dbInfo.db.Prepare("SELECT id, cardname, cardset FROM cards")

	// add statements to map
	dbInfo.stmts[addCard] = addCardStmt
	dbInfo.stmts[getAllCards] = getAllStmt

	dbInfo.GetAllCards()
	printCards()

	card := Card{
		id:   2,
		name: "Handheld Blaster",
		set:  "Wave 2",
	}

	dbInfo.AddCard(&card)

	dbInfo.GetAllCards()
	printCards()
}

func printCards() {
	for _, card := range cardMap {
		fmt.Printf("%d %s %s\n", card.id, card.name, card.set)
	}
}

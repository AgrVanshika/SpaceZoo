package lib

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Method struct {
	Topic       string
	Name        string
	Description string
	Source      string
}

type Technique struct {
	Name        string
	Description string
	Source      string
}

func GetRandomFromCategory(category string) {
	// if windows
	// PATH=$PATH:$(where python3)
	// if linux
	// PATH=$PATH:$(which python3)
	// Connect to the "bank" database
	// Attempt to connect
	// config, err := pgx.ParseConfig(CONNSTRING)
	// if err != nil {
	// 	log.Fatal("error configuring the database: ", err)
	// }
	db, err := sql.Open("postgres", CONNSTRING)
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}
	defer db.Close()

}

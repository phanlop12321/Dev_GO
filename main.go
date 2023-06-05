package main

import (
	"log"

	"github.com/phanlop12321/Dev_GO/db"
)

func main() {
	db, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Reset(); err != nil {
		log.Fatal(err)
	}
	if err := db.AutoMigrate(); err != nil {
		log.Fatal(err)
	}

}

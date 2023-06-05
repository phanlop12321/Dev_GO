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
}

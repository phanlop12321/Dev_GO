package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/phanlop12321/Dev_GO/db"
	"github.com/phanlop12321/Dev_GO/handler"
)

func main() {
	db, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	// if err := db.Reset(); err != nil {
	// 	log.Fatal(err)
	// }
	if err := db.AutoMigrate(); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.GET("/income", handler.Getincome(db))
	r.POST("/register", handler.Register(db))
	r.POST("/login", handler.Login(db))

	r.Run(":8080")

}

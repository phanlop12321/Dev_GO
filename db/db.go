package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	db *gorm.DB
}

func NewDB() (*DB, error) {
	url := "host=localhost user=peagolang password=supersecret dbname=peagolang port=64329 sslmode=disable"
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	return &DB{db: db}, nil
}
func (db *DB) Reset() error {
	err := db.db.Migrator().DropTable(
		&Savings{},
		&Income{},
		&Expenses{},
		&User{},
	)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
func (db *DB) AutoMigrate() error {
	err := db.db.Migrator().AutoMigrate(
		&Savings{},
		&Income{},
		&Expenses{},
		&User{},
	)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

type Savings struct {
	ID            uint `gorm:"primaryKey"`
	IDUser        string
	Total         uint
	IncomeTotal   uint
	ExpensesTotal uint
}

type Income struct {
	ID          uint `gorm:"primaryKey"`
	IDUser      string
	Description string
	Money       uint
}
type Expenses struct {
	ID          uint `gorm:"primaryKey"`
	IDUser      string
	Description string
	Money       uint
}
type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"uniqueIndex"`
	Password string
}

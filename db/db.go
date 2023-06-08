package db

import (
	"log"

	"github.com/phanlop12321/Dev_GO/model"
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
		&User{},
	)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

type Savings struct {
	ID          uint `gorm:"primaryKey"`
	IDUser      string
	Description string
	Money       uint
	Status      bool
}
type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"uniqueIndex"`
	Password string
}

func (db *DB) CreateIncome(i *model.Income) error {
	income := Savings{
		IDUser:      i.IDUser,
		Description: i.Description,
		Money:       i.Money,
		Status:      true,
	}
	if err := db.db.Create(&income).Error; err != nil {
		return err
	}
	i.ID = income.ID
	return nil
}

func (db *DB) CreateExpenses(i *model.Income) error {
	income := Savings{
		IDUser:      i.IDUser,
		Description: i.Description,
		Money:       i.Money,
		Status:      false,
	}
	if err := db.db.Create(&income).Error; err != nil {
		return err
	}
	i.ID = income.ID
	return nil
}

func (db *DB) CreateUser(u *model.User) error {
	user := User{
		Username: u.Username,
		Password: u.Password,
	}
	if err := db.db.Create(&user).Error; err != nil {
		return err
	}
	u.ID = user.ID
	return nil
}
func (db *DB) GetUserByUsername(username string) (*model.User, error) {
	var user User
	if err := db.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &model.User{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
	}, nil
}
func (db *DB) GetIncome() ([]model.Income, error) {
	var income []Savings
	if err := db.db.Find(&income, "Status = ?", true).Error; err != nil {
		return nil, err
	}
	result := []model.Income{}
	for _, inc := range income {
		result = append(result, model.Income{
			ID:          inc.ID,
			IDUser:      inc.IDUser,
			Description: inc.Description,
			Money:       inc.Money,
			Status:      true,
		})
	}
	return result, nil

}
func (db *DB) GetExpenses() ([]model.Income, error) {
	var expenses []Savings
	if err := db.db.Find(&expenses, "Status = ?", false).Error; err != nil {
		return nil, err
	}
	result := []model.Income{}
	for _, inc := range expenses {
		result = append(result, model.Income{
			ID:          inc.ID,
			IDUser:      inc.IDUser,
			Description: inc.Description,
			Money:       inc.Money,
			Status:      false,
		})
	}
	return result, nil

}

func (db *DB) DeletSaveByID(id uint) error {
	var save Savings
	if err := db.db.Where("id = ?", id).Delete(&save).Error; err != nil {
		return err
	}
	return nil
}

func (db *DB) GetSaveByID(id uint) (*model.Income, error) {
	var save Savings
	if err := db.db.Where("id = ?", id).First(&save).Error; err != nil {
		return nil, err
	}
	return &model.Income{
		ID: save.ID,
	}, nil
}

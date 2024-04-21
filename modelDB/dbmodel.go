package modeldb

import (
	modelapp "goFinalTask/modelAPP"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(dsn string) error {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db
	log.Println("DB Connected")
	db.AutoMigrate(&modelapp.Users{})
	log.Println("Create if not exists Users table in DB")
	db.AutoMigrate(&modelapp.Transactions{})
	log.Println("Create if not exists Transactions table in DB")

	return nil
}

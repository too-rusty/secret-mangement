package db

import (
	"fmt"
	"secretemail/models"
	"sync"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var db *gorm.DB

var initialize sync.Once

func Initialize() (*gorm.DB, error) {
	var err error
	host := viper.GetString("database.host")
	port := viper.GetInt("database.port")
	user := viper.GetString("database.user")
	password := viper.GetString("database.password")
	dbname := viper.GetString("database.dbname")

	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	initialize.Do(func() {

		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return
		}

	})

	if err != nil {
		fmt.Println("could not be connected!")
		return nil, err
	}

	fmt.Println("Successfully connected!")

	return db, nil

}

// Run migrations
func Migrate() error {
	_, err := Initialize()
	if err != nil {
		return err
	}
	return db.AutoMigrate(
		&models.Key{},
		&models.Message{},
	)
}

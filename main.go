package main

import (
	"fmt"
	"log"
	"secretemail/db"
	"secretemail/server"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	if err := envSetup(); err != nil {
		log.Fatalln(err)
	}
	server, err := server.New()
	if err != nil {
		log.Fatalln(err)
	}

	go server.StartBgTask()

	if err := db.Migrate(); err != nil {
		log.Fatalln(err)
	}

	if err := server.AttachEndpoints().Run(); err != nil {
		log.Fatalln(err)
	}
}

func envSetup() error {

	// Load the .env file in the current directory
	godotenv.Load()
	// or

	// godotenv.Load(".env")

	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(".")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		return fmt.Errorf("fatal error config file: %w", err)
	}
	return nil
}

package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

// Connect opens a connection to the database.
// The function returns a *gorm.DB object.
func Connect() *gorm.DB {

	// Load the environment variables from the .env file.
	err := godotenv.Load(".env")
	if err != nil {
		// If there was an error loading the environment variables, panic.
		panic(err)
	}

	// Create a connection string.
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"))

	// Open a connection to the database using the connection string.
	db, err := gorm.Open(mysql.Open(dns))
	if err != nil {
		// If there was an error opening the connection, panic.
		panic(err)
	}

	// Return the database connection.
	return db
}

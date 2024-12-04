// databases/postgres.go
package databases

import (
	"log"
	"shtlpg/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Import PostgreSQL dialect
	"github.com/spf13/viper"
)

var db *gorm.DB

// ConnectPostgres initializes the PostgreSQL connection with GORM
func ConnectPostgres() error {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
		return err
	}

	dbURL := viper.GetString("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set in .env")
		return nil
	}

	var err error
	db, err = gorm.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Could not connect to PostgreSQL: %v", err)
		return err
	}

	// Auto migrate your models to ensure the tables are created
	db.AutoMigrate(&models.User{})

	log.Println("Successfully connected to PostgreSQL")
	return nil
}

// GetDB returns the GORM DB instance
func GetDB() *gorm.DB {
	return db
}

// CloseDB closes the database connection
func CloseDB() {
	db.Close()
}

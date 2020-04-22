package yahooScraper

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
)

func saveForex(forex ForexPair) ForexPair {
	s := fmt.Sprintf("dbname=%s user=%s password=%s port=5432", os.Getenv("DATABASE_NAME"), os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"))
	db, err := gorm.Open("postgres", s)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.NewRecord(forex)
	db.Create(&forex)

	return forex
}

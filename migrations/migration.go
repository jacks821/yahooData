package migrations

import (
	scripts "../scripts"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
)

func runMigrations() {
	s := fmt.Sprintf("dbname=%s user=%s password=%s port=5432", os.Getenv("DATABASE_NAME"), os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"))
	db, err := gorm.Open("postgres", s)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.AutoMigrate(&scripts.ForexPair{}, &scripts.CryptoPair{})
	fmt.Println("Successful Migration!")
}

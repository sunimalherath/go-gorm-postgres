package main

import (
	"log"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	Name string
	Age  int
}

func main() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=testgormdb password=root sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()
	db.AutoMigrate(&User{})
}

package crud

import (
	_ "fmt"
	"log"
	"mta9896/restapi/internal/database"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type Item struct {
	Id int `json:"id" validate:"required,min=1" db:"id"`
	Title string `json:"title" validate:"required" db:"title"`
	Description string `json:"description" db:"description"`
}

func List() []Item {
	item := Item{}
	rows := database.FetchAllItems()


	var items []Item

	for rows.Next() {
		if err := rows.StructScan(&item); err != nil {
			log.Fatal(err)
		}

		log.Println(item)

		items = append(items, item)
	}

	err := rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return items
}

func Create(item Item) {

}
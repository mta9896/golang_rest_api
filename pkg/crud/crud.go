package crud

import ( 
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type Item struct {
	Id int `json:"id" validate:"required,min=1"`
	Title string `json:"title" validate:"required"`
	Description string `json:"description"`
}

var items = []Item {
	{
		Id: 1,
		Title: "First Title",
		Description: "First description",
	},

	{
		Id: 2,
		Title: "Second Title",
		Description: "Second description",
	},

	{
		Id: 3,
		Title: "Third Title",
		Description: "Third description",
	},
}

func List() []Item {
	return items
}

func Create(item Item) {
	items = append(items, item)

	return
}
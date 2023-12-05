package database

import (
	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"mta9896/restapi/internal/entity"
)

var db *sqlx.DB

func Initialize() (err error) {
	db, err = sqlx.Open("mysql", "user:password@tcp(localhost:3306)/mydatabase")
    if err != nil {
        return err
    }

	//defer db.Close()

	err = db.Ping()
	if err != nil {
		return err
	}

	return nil
}

func FetchAllItems() ([]entity.Item, error) {
	rows, err := db.Queryx("SELECT id, title from items")
	if err != nil {
		return nil, err
	}

	var item entity.Item
	var items []entity.Item

	for rows.Next() {
		if err := rows.StructScan(&item); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return items, nil
}

func InsertItem(item entity.Item) (err error) {
	query := "insert into items (id, title, description) values (:id, :title, :description)"
	
	_, err = db.NamedExec(query, item)
	if err != nil {
		return err
	}

	return nil
}
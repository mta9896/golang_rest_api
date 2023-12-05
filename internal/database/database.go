package database

import (
	"log"
	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func FetchAllItems() *sqlx.Rows {
	db, err := sqlx.Open("mysql", "user:password@tcp(localhost:3306)/mydatabase")
    if err != nil {
        log.Fatal(err)
    }

    defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Queryx("SELECT id, title from items")
	if err != nil {
		log.Fatal(err)
	}
	//defer rows.Close()

	// for rows.Next() {
	// 	if err := rows.StructScan(&item); err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	log.Println(item)

	// 	//items = append(items, item)
	// }

	// err = rows.Err()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	return rows
}
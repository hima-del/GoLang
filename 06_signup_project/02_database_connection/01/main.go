package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type User struct {
	username string
	password string
}

func main() {
	db, err := sql.Open("postgres", "postgres://admin:password@localhost/project?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")

	rows, err := db.Query("SELECT * FROM users;")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	bks := make([]User, 0)
	for rows.Next() {
		bk := User{}
		err := rows.Scan(&bk.username, &bk.password) // order matters
		if err != nil {
			panic(err)
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}

	for _, bk := range bks {
		// fmt.Println(bk.isbn, bk.title, bk.author, bk.price)
		fmt.Printf("%s, %s\n", bk.username, bk.password)
	}
}

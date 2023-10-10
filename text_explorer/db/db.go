package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "hfedoriv_texts"
)

var db_connection *sql.DB

func DB_connect() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	fmt.Println("Successfully connected!")
	db_connection = db
	return nil
}

func Article_add(title string, text string, user_id int8) error {
	sqlStatement := "INSERT INTO articles (title, text, user_id) VALUES ($1, $2, $3)"
	_, err := db_connection.Exec(sqlStatement, title, text, user_id)
	return err
}

func Article_update(title string, text string, id int8) error {
	sqlStatement := "UPDATE articles SET title = $1, text = $2 WHERE id = $3;"
	_, err := db_connection.Exec(sqlStatement, title, text, id)
	return err
}

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

func Article_get() error {
	sqlStatement := "SELECT * FROM articles;"
	_, err := db_connection.Exec(sqlStatement)
	return err
}

func Article_get_all(id uint8) error {
	sqlStatement := "SELECT * FROM articles WHERE id = $1;"
	_, err := db_connection.Exec(sqlStatement, id)
	return err
}

func Article_add(title string, text string, user_id uint8) error {
	sqlStatement := "INSERT INTO articles (title, text, user_id) VALUES ($1, $2, $3)"
	_, err := db_connection.Exec(sqlStatement, title, text, user_id)
	return err
}

func Article_update(title string, text string, id uint8) error {
	sqlStatement := "UPDATE articles SET title = $1, text = $2 WHERE id = $3;"
	_, err := db_connection.Exec(sqlStatement, title, text, id)
	return err
}

func Article_delete(id uint8) error {
	sqlStatement := "DELETE FROM articles WHERE id = $1;"
	_, err := db_connection.Exec(sqlStatement, id)
	return err
}

package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5434 // PostgreSQL 預設端口
	user     = "postgres"
	password = "ar0708"
	dbname   = "postgres"
)

var DB *sql.DB

func ConnectToDB() *sql.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	//寫死 之後研究怎麼跑migration
	migration, err := os.ReadFile("database/migrations/20240116_create_todoList.sql")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(string(migration))
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to database")
	DB = db // 設置全局變量 DB
	return db
}

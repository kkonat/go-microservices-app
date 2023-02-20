package main

import (
	"authentication/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "80"

var counts int64

type Config struct {
	Repo data.Repository
}

func main() {
	log.Println("Starting auth service")

	conn := connectToDB()
	if conn == nil {
		log.Panic("Unable to connect to database")
	}

	app := Config{}
	app.setupRepo(conn)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")

	for {
		log.Printf("connecting: %s", dsn)
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Postgress not yet ready")
			counts++
		} else {
			log.Println("Connected to PostgreSQL database")
			return connection
		}
		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Sleeping... ")
		time.Sleep(2 * time.Millisecond)
	}
}

func (app *Config) setupRepo(conn *sql.DB) {
	db := data.NewPostgresRepository(conn)
	app.Repo = db
}

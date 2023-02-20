package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const webPort = "80"

type Config struct {
	Rabbit *amqp.Connection
}

func main() {

	rabbitConn, err := connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	app := Config{
		Rabbit: rabbitConn,
	}

	log.Printf("Starting broker service on port... %s\n", webPort)

	// define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// start the server
	log.Printf("Listening...")
	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}

func connect() (*amqp.Connection, error) {

	var counts int64
	var waitTime = int64(2)
	var connection *amqp.Connection

	log.Println("Attempting to connect to  Rabbit...")

	for {
		c, err := amqp.Dial("amqp://guest:guest@rabbit")
		if err != nil {
			log.Println("Error dialing", err)
			log.Println("Waiting for reconnect")
			counts++
		} else {
			log.Println("Connected to RabbitMQ")
			connection = c
			break
		}

		if counts > 5 {
			fmt.Println(err)
			return nil, err
		}

		waitTime *= waitTime
		log.Printf("Warting for %d seconds...", waitTime)
		time.Sleep(time.Second * time.Duration(waitTime))
	}
	return connection, nil
}

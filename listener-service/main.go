package main

import (
	"fmt"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	rabbitConn, err := connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	defer rabbitConn.Close()
}

func connect() (*amqp.Connection, error) {

	var counts int64
	var waitTime = int64(2)
	var connection *amqp.Connection

	log.Println("Attempting to connect to  Rabbit...")

	for {
		c, err := amqp.Dial("amqp://guest:guest@localhost")
		if err != nil {
			log.Println("Error dialing", err)
			log.Println("Waiting for reconnect")
			counts++
		} else {
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

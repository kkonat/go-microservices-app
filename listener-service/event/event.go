package event

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func declareExchange(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(
		"logs_topic", // exch queueName
		"topic",      // type
		true,         // is durable?
		false,        // is to be auto-deleted?
		false,        // internal exch ?
		false,        // no-wait?
		nil,          // arguments
	)
}

func declareRandomQueue( ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"", // name
		false, // durable
		false, // delete autho when unused
		true, // exlusive?
		false, // no-wait?
		nil, // attribs
	) 
}
package mq

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type MQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func New(url string) (*MQ, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}
	return &MQ{conn: conn, channel: ch}, nil
}

func (m *MQ) Close() {
	if m.channel != nil {
		_ = m.channel.Close()
	}
	if m.conn != nil {
		_ = m.conn.Close()
	}
}

func (m *MQ) Publish(exchange, routingKey string, body []byte) error {
	// ensure queue exists for the example (simple direct publish)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := m.channel.PublishWithContext(ctx,
		exchange,   // exchange
		routingKey, // routing key (queue name)
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
			Timestamp:   time.Now(),
		})
	return err
}

func (m *MQ) StartConsumer(queue string, handler func([]byte)) error {
	// declare queue
	_, err := m.channel.QueueDeclare(
		queue,
		true,  // durable
		false, // autoDelete
		false, // exclusive
		false, // noWait
		nil,
	)
	if err != nil {
		return err
	}

	msgs, err := m.channel.Consume(
		queue,
		"",    // consumer
		true,  // autoAck - consider manual ack for production
		false, // exclusive
		false, // noLocal
		false, // noWait
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			// simple handler
			handler(d.Body)
		}
		log.Println("consumer stopped")
	}()

	return nil
}
package rabbitmq

import (
	"context"

	"github.com/streadway/amqp"
)

type AuditService interface {
	Insert(ctx context.Context, msg []byte) error
}

type Client struct {
	conn    *amqp.Connection
	ch      *amqp.Channel
	msgs    <-chan amqp.Delivery
	service AuditService
}

func NewClient(service AuditService) *Client {
	return &Client{
		conn:    new(amqp.Connection),
		ch:      new(amqp.Channel),
		msgs:    make(<-chan amqp.Delivery),
		service: service,
	}
}

func (s *Client) closeConnection() error {
	return s.conn.Close()
}

func (s *Client) closeChannel() error {
	return s.ch.Close()
}

func (s *Client) Close() error {
	err := s.closeChannel()
	if err != nil {
		return err
	}

	err = s.closeConnection()
	return err
}

func (s *Client) Listening(uri string) error {
	conn, err := amqp.Dial(uri)
	if err != nil {
		return err
	}

	s.conn = conn

	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	s.ch = ch

	q, err := s.ch.QueueDeclare(
		"messageQueue", // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		return err
	}

	s.msgs, err = s.ch.Consume(
		q.Name, //queue
		"",     //consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *Client) Serve(ctx context.Context) error {
	for d := range s.msgs {
		if err := s.service.Insert(context.Background(), d.Body); err != nil {
			return err
		}
	}
	return nil
}

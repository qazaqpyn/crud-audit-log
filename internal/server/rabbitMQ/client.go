package rabbitmq

import "github.com/streadway/amqp"

type Client struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewClient(uri string) (*Client, error) {
	conn, err := amqp.Dial(uri)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &Client{
		conn: conn,
		ch:   ch,
	}, nil
}

func (s *Client) closeConnection() error {
	return s.conn.Close()
}

func (s *Client) closeChannel() error {
	return s.ch.Close()
}

func (s *Client) CloseServer() error {
	err := s.closeChannel()
	if err != nil {
		return err
	}

	err = s.closeConnection()
	return err
}

func (s *Client) StartListening() (<-chan amqp.Delivery, error) {
	q, err := s.ch.QueueDeclare(
		"messageQueue", // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		return nil, err
	}

	msgs, err := s.ch.Consume(
		q.Name, //queue
		"",     //consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}

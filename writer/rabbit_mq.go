package writer

import (
	"encoding/json"
	"fmt"
	"github.com/hyperyuri/logg"
	"time"

	"github.com/streadway/amqp"
)

type rabbitW struct {
	user string
	pass string
	host string
	port int
	queue string
	conn *amqp.Connection
}

//WithRabbitMode send the logs to a rabbitmq queue with the core.Message ToJSON() struct
func WithRabbitMode(queue, user, pass, host string, port int) *rabbitW {
	return &rabbitW{
		user: user,
		pass: pass,
		queue: queue,
		host: host,
		port: port,
	}
}

func  (s *rabbitW) Write(m *logg.Message) {
	for i := 0; i < 5; i++ {
		if s.conn == nil {
			err := s.startConn()
			if err != nil {
				time.Sleep(time.Minute * 1)
				continue
			}
		}

		ch, err := s.conn.Channel()
		if err != nil {
			time.Sleep(time.Minute * 1)
			continue
		}


		q, err := ch.QueueDeclare(
			"hello", // name
			false,   // durable
			false,   // delete when unused
			false,   // exclusive
			false,   // no-wait
			nil,     // arguments
		)
		if err != nil {
			time.Sleep(time.Minute * 1)
			continue
		}

		message, err := json.Marshal(m.ToJSON())
		if err != nil {
			time.Sleep(time.Minute * 1)
			continue
		}

		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing {
				ContentType: "application/json",
				Body:        message,
			})
		if err != nil {
			time.Sleep(time.Minute * 1)
			continue
		}

		ch.Close()
		break;
	}
}

func  (s *rabbitW) startConn() error{
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", s.user, s.pass, s.host, s.port))
	if err != nil {
		return err
	}

	s.conn = conn
	return nil
}

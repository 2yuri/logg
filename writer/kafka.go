package writer

import (
	"context"
	"encoding/json"
	"github.com/hyperyuri/logg"
	"github.com/segmentio/kafka-go"
	"time"
)

type kafkaW struct {
	brokers []string
	topic string
	writer *kafka.Writer
}

//WithKafkaMode send the logs to a kafka topic with the core.Message ToJSON() struct
func WithKafkaMode(broker []string, topic string) *kafkaW {
	return &kafkaW{
		brokers: broker,
		topic: topic,
	}
}

func (s *kafkaW) Write(m *logg.Message) {
	if s.writer == nil {
		s.setWriter()
	}

	for i := 0; i < 5; i++ {
		message, err := json.Marshal(m.ToJSON())
		if err != nil {
			time.Sleep(time.Minute * 1)
			continue
		}
		err = s.writer.WriteMessages(context.Background(), kafka.Message{
			Value: message,
		})
		if err != nil {
			time.Sleep(time.Minute * 1)
			continue
		}
		break
	}
}

func (s *kafkaW) setWriter() {
	s.writer = &kafka.Writer{
		Addr:  kafka.TCP(s.brokers...),
		Topic: s.topic,
	}
}


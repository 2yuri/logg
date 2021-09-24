package writer

import "github.com/segmentio/kafka-go"

type kafkaW struct {
	brokers []string
	topic string
}

var reader *kafka.Writer

func (s *kafkaW) write() error {
	if reader == nil {
		reader = s.newReader()
	}


}

func (s *kafkaW) newReader() error {

}


package core

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

type Message struct {
	text string
	datetime string
	path string
	level int
	levelLabel string
}

func (s *Message) SetDatetime(datetime time.Time) {
	s.datetime = datetime.Format("2006/01/02 15:04:05")
}

func (s *Message) SetPath() {
	_, fn, line, _ := runtime.Caller(3)
	file := strings.Split(fn, "/")

	s.path = fmt.Sprintf("%v:%v", file[len(file)-1], line)
}

func NewMessage(level int, text string) *Message {
	m := &Message{text: text, level: level}
	m.SetLevelLabel(level)
	m.SetPath()
	m.SetDatetime(time.Now())

	return m
}

func (s *Message) SetLevelLabel(level int) {
	switch level {
	case PanicLevel:
		s.levelLabel = "PANIC"
		break
	case InfoLevel:
		s.levelLabel = "INFO"
		break
	case ErrorLevel:
		s.levelLabel = "ERROR"
		break
	case WarningLevel:
		s.levelLabel = "WARNING"
		break
	}
}

func (s *Message) ToBytes() {

}

func (s *Message) String() string {
	return fmt.Sprintf("%v %s %s - %s", s.datetime, s.path, s.levelLabel, s.text)
}
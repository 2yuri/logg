package core

type IWriter interface {
	Write(m *Message)
}
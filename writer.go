package logg

type Writer interface {
	Write(m *Message)
}
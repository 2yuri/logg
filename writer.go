package logg

type Writer interface {
	//Write should send your message (kafka, rest, rabbit or others)
	Write(m *Message)
}
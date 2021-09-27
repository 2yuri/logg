package writer

import (
	"fmt"
	"github.com/hyperyuri/logg"
)

type debug struct {
}

func (s *debug) Write(m *logg.Message) {
	fmt.Println(m.String())
}

//WithDebugMode set the logs to show in terminal
func WithDebugMode() *debug {
	return &debug{}
}
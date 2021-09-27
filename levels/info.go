package levels

import (
	"fmt"
	"github.com/hyperyuri/logg"
)

type Info struct {
	*logg.Logger
}

func (s *Info) Info(patter string, args ...interface{}) {
	m := logg.NewMessage(logg.InfoType, fmt.Sprintf(patter, args...), s.Stack(), s.App())

	go func() {
		s.Writer.Write(m)
	}()
}

func (s *Info) Infof(patter string, args ...interface{}) {
	m := logg.NewMessage(logg.InfoType, fmt.Sprintf(patter, args...), s.Stack(), s.App())

	go func() {
		s.Writer.Write(m)
	}()
}

package levels

import (
	"fmt"
	"github.com/hyperyuri/logg"
)

type Warning struct {
	*logg.Logger
}

func (s *Warning) Warning(patter string, args ...interface{}) {
	m := logg.NewMessage(logg.WarningType, fmt.Sprintf(patter, args...), s.Stack(), s.App())

	go func() {
		s.Writer.Write(m)
	}()
}

func (s *Warning) Warningf(patter string, args ...interface{}) {
	m := logg.NewMessage(logg.WarningType, fmt.Sprintf(patter, args...), s.Stack(), s.App())

	go func() {
		s.Writer.Write(m)
	}()
}

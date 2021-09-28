package levels

import (
	"fmt"
	"github.com/hyperyuri/logg"
)

type Error struct {
	*logg.Logger
}

func (s *Error) Error(patter string, args ...interface{}) {
	m := logg.NewMessage(logg.ErrorType, fmt.Sprintf(patter, args...), s.Stack(), s.App())

	s.Writer.Write(m)
}

func (s *Error) SyncError(patter string, args ...interface{}) {
	m := logg.NewMessage(logg.ErrorType, fmt.Sprintf(patter, args...), s.Stack(), s.App())

	go func() {
		s.Writer.Write(m)
	}()
}

package levels

import (
	"fmt"
	"github.com/hyperyuri/logg"
	"os"
)

type Panic struct {
	*logg.Logger
}

func (s *Panic) Panic(patter string, args ...interface{}) {
	m := logg.NewMessage(logg.PanicType, fmt.Sprintf(patter, args...), s.Stack(), s.App())
	s.Writer.Write(m)

	os.Exit(0)
}
package v1

import (
	"fmt"
	"github.com/ecocentauro/logg/core"
	"log"
)

type Error struct {
}

func (s *Error) Error(patter string, args ...interface{}) {
	m := core.NewMessage(core.ErrorLevel, fmt.Sprintf(patter, args...))
	core.Default.Write(m)
	return
}

func (s *Error) Errorf(patter string, args ...interface{}) {
	m := core.NewMessage(core.ErrorLevel, fmt.Sprintf(patter, args...))
	log.Print(m)
	return
}

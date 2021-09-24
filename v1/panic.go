package v1

import (
	"github.com/ecocentauro/logg/core"
	"fmt"
	"log"
)

type Panic struct {
}

func (s *Panic) Panicf(patter string, args ...interface{}) {
	m := core.NewMessage(core.PanicLevel, fmt.Sprintf(patter, args...))
	log.Print(m)
	return
}

func (s *Panic) Panic(patter string, args ...interface{}) {
	m := core.NewMessage(core.PanicLevel, fmt.Sprintf(patter, args...))
	log.Print(m)
	return
}

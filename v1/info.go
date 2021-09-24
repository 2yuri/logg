package v1

import (
	"github.com/ecocentauro/logg/core"
	"fmt"
	"log"
)

type Info struct {
}

func (s *Info) Info(patter string, args ...interface{}) {
	m := core.NewMessage(core.InfoLevel, fmt.Sprintf(patter, args...))
	log.Print(m)
	return
}

func (s *Info) Infof(patter string, args ...interface{}) {
	m := core.NewMessage(core.InfoLevel, fmt.Sprintf(patter, args...))
	log.Print(m)
	return
}

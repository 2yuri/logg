package v1

import (
	"fmt"
	"github.com/ecocentauro/logg/core"
)

type Warning struct {
	core.IWriter
}

func (s *Warning) Warning(patter string, args ...interface{}) {
	m := core.NewMessage(core.WarningLevel, fmt.Sprintf(patter, args...))
	fmt.Println(m.String())
	return
}

func (s *Warning) Warningf(patter string, args ...interface{}) {
	m := core.NewMessage(core.WarningLevel, fmt.Sprintf(patter, args...))
	fmt.Println(m.String())
	return
}

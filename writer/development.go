package writer

import (
	"fmt"
	"github.com/ecocentauro/logg/core"
)

type debug struct {
}

func (s *debug) Write(m core.Message) error {
	fmt.Println(m.String())
	return nil
}

func WithDebugMode() *debug {
	return &debug{}
}
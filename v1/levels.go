package v1

import "github.com/ecocentauro/logg/core"

type Levels struct {
	core.ILevels
}

func WithLevel() *Levels {
	return &Levels{
		struct {
			core.ErrorInterface
			core.WarningInterface
			core.InfoInterface
			core.PanicInterface
		}{
			&Error{},
			&Warning{},
			&Info{},
			&Panic{},
		},
	}
}



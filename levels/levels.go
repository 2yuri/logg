package levels

import "github.com/hyperyuri/logg"

type Levels struct {
	logg.Levels
}


func NewLevels(logger *logg.Logger) *Levels {
	return &Levels{struct {
		logg.ErrorLevel
		logg.WarningLevel
		logg.InfoLevel
		logg.PanicLevel
	}{
		&Error{
			logger,
		},
		&Warning{
			logger,
		},
		&Info{
			logger,
		},
		&Panic{
			logger,
		},
	},
	}
}

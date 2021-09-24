package main

import (
	"github.com/ecocentauro/logg/core"
	v1 "github.com/ecocentauro/logg/v1"
	"github.com/ecocentauro/logg/writer"
)

func main() {
	logger := core.NewCore(v1.WithLevel(), writer.WithDebugMode())
	logger.Error("Err on create text basic brother")
}

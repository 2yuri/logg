package main

import (
	"github.com/hyperyuri/logg"
	"github.com/hyperyuri/logg/levels"
	"github.com/hyperyuri/logg/writer"
	"time"
)

func main() {
	logger := logg.NewLogger().
		WithWriter(writer.WithDebugMode()).
		WithDefaultConfig(logg.NewDefaultConfig("team-x", "project-x"))
	logger.SetLevels(levels.NewLevels(logger))

	logger.Error("This is an example of an error")
	//logger.Panic("system down")
	logger.Info("This not will show up!")

	time.Sleep(time.Second * 10)
}

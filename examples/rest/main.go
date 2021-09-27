package main

import (
"github.com/hyperyuri/logg"
"github.com/hyperyuri/logg/levels"
"github.com/hyperyuri/logg/writer"
)

func main() {
	logger := logg.NewLogger().
		WithWriter(writer.WithRabbitMode("test", "admin", "12345", "localhost", 9000)).
		WithDefaultConfig(logg.NewDefaultConfig("team-x", "project-x"))
	logger.SetLevels(levels.NewLevels(logger))

	logger.Error("This is an erros message")
}
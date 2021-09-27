package main

import (
	"github.com/hyperyuri/logg"
	"github.com/hyperyuri/logg/levels"
	"github.com/hyperyuri/logg/writer"
)

func main() {
	logg1 := logg.NewLogger().
		WithWriter(writer.WithKafkaMode([]string{"localhost:9092"}, "test")).
		WithDefaultConfig(logg.NewDefaultConfig("team-x", "project-x"))
	logg1.SetLevels(levels.NewLevels(logg1))
	logg1.Error("This is an erros message")

	logg2 := logg.NewLogger().
		WithWriter(writer.WithKafkaMode([]string{"localhost:9092"}, "test")).
		WithDefaultConfig(logg.NewDefaultConfig("team-y", "project-y"))
	logg2.SetLevels(levels.NewLevels(logg2))
	logg2.Error("This is an erros message")
}

# 🔥 logg 

Open Source Cloud logging engine in Go.

## Instalation
``go get -u github.com/hyperyuri/logg``

## Quick Start

```
//Setup logger
logger := logg.NewLogger().
WithWriter(writer.WithDebugMode()).
WithDefaultConfig(logg.NewDefaultConfig("team-x", "project-x"))
logger.SetLevels(levels.NewLevels(logger))

//To use the log
logger.Error("This is an example of an error")
```
- Check the examples folder to see more options

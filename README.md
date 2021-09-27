# 🔥 logg 

## Open Source Cloud logging library in Go. 

## About the project
- Send your logs to kafka, rabbitMQ or REST! (Feel free to implement more writers if you want!)

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

## 💪 How to contribute to the project?
- Do a fork of the project.
- Create a new branch with your changes: git checkout -b my-feature
- Save your changes and create a commit with a message: git commit -m "feature: My new feature"
- Send your changes and send a PR

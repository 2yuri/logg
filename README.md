# 🔥 logg

## Open Source Cloud logging library in Go.

## About the project

- Connect your golang microservices logs with this engine!
- Send your logs to kafka, rabbitMQ or REST! And then manage them!
- Feel free to implement more writers

## Instalation

`go get -u github.com/hyperyuri/logg`

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
- Submitting issue (please, follow [template](https://github.com/hyperyuri/logg/tree/master/.github/ISSUE_TEMPLATE/ISSUE_FORM.md))
- Create a new branch with your changes: git checkout -b my-feature
- Save your changes and create a commit with a message: git commit -m "feature: My new feature"
- Send your changes and send a PR

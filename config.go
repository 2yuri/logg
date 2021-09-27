package logg

type DefaultConfig struct {
	stack string
	app string
}

func NewDefaultConfig(stack string, app string) *DefaultConfig {
	return &DefaultConfig{stack: stack, app: app}
}

package logg

type DefaultConfig struct {
	stack string
	app string
}

//NewDefaultConfig create a config that will set your project and the maintainer
func NewDefaultConfig(stack string, app string) *DefaultConfig {
	return &DefaultConfig{stack: stack, app: app}
}

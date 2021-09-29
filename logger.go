package logg

type Logger struct {
	Levels
	Writer

	config *DefaultConfig
}

func (l *Logger) Stack() string {
	if l.config == nil {
		return ""
	}

	return l.config.stack
}

func (l *Logger) App() string {
	if l.config == nil {
		return ""
	}

	return l.config.app
}

func (l *Logger) SetLevels(lev Levels) {
	l.Levels = lev
}

func NewLogger() *Logger {
	core := &Logger{}
	return core
}

func (l *Logger) WithWriter(writer Writer) *Logger{
	l.Writer = writer

	return l
}

func (l *Logger) WithDefaultConfig(config *DefaultConfig) *Logger{
	l.config = config

	return l
}

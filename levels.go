package logg

type Levels interface {
	ErrorLevel
	WarningLevel
	InfoLevel
	PanicLevel
}

type ErrorLevel interface {
	Errorf(patter string, args ...interface{})
	Error(patter string, args ...interface{})
}

type WarningLevel interface {
	Warningf(patter string, args ...interface{})
	Warning(patter string, args ...interface{})
}

type PanicLevel interface {
	Panicf(patter string, args ...interface{})
	//Panic will break up your code after print the error, os.Exit(0)
	Panic(patter string, args ...interface{})
}

type InfoLevel interface {
	Infof(patter string, args ...interface{})
	Info(patter string, args ...interface{})
}
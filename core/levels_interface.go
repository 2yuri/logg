package core

type ILevels interface {
	ErrorInterface
	WarningInterface
	InfoInterface
	PanicInterface
}

type ErrorInterface interface {
	Errorf(patter string, args ...interface{})
	Error(patter string, args ...interface{})
}

type WarningInterface interface {
	Warningf(patter string, args ...interface{})
	Warning(patter string, args ...interface{})
}

type PanicInterface interface {
	Panicf(patter string, args ...interface{})
	Panic(patter string, args ...interface{})
}

type InfoInterface interface {
	Infof(patter string, args ...interface{})
	Info(patter string, args ...interface{})
}
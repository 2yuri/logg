package logg

type Levels interface {
	ErrorLevel
	WarningLevel
	InfoLevel
	PanicLevel
}

type ErrorLevel interface {
	//Error format to an error log and send to some writer
	Error(patter string, args ...interface{})

	//SyncError run in concurrency, format to an error log and send to some writer
	SyncError(patter string, args ...interface{})
}

type WarningLevel interface {
	//Warning format to a warning log and send to some writer
	Warning(patter string, args ...interface{})

	//SyncWarning run in concurrency, format to a warning log and send to some writer
	SyncWarning(patter string, args ...interface{})
}

type PanicLevel interface {
	//Panic format to a PANIC log and send to some writer, after it runs os.Exit
	Panic(patter string, args ...interface{})
}

type InfoLevel interface {
	//Info format to a info log and send to some writer
	Info(patter string, args ...interface{})

	//SyncInfo run in concurrency, format to a info log and send to some writer
	SyncInfo(patter string, args ...interface{})
}
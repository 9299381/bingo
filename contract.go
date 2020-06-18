package bingo

//------------------//
type ILogger interface {
	Info(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Trace(args ...interface{})
	Debug(args ...interface{})
	Panic(args ...interface{})

	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Tracef(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
}

//--------------------//
type IProvider interface {
	Boot()
	Register()
}

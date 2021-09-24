package core

type Core struct {
	ILevels
	IWriter
}

var Default *Core

func NewCore(pkg ILevels, writer IWriter) *Core {
	core := &Core{
		pkg,
		writer,
	}

	Default = core

	return core
}

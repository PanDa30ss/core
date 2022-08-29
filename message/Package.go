package message

type DumpFunc func(msg *Message)
type LoadFunc func(msg *Message)
type IPackage interface {
	GetDumpFunc() DumpFunc
	GetLoadFunc() LoadFunc
}

type Package struct {
	Dump DumpFunc
	Load LoadFunc
}

func (this *Package) GetDumpFunc() DumpFunc {
	return this.Dump
}
func (this *Package) GetLoadFunc() LoadFunc {
	return this.Load
}

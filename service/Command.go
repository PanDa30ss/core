package service

type ICommand interface {
	Execute()
}
type ExecuteFunc func(params ...interface{})

type command struct {
	executeFunc ExecuteFunc
	params      []interface{}
}

func (this *command) Execute() {
	this.executeFunc(this.params...)
}

func MakeCommand(executeFunc ExecuteFunc, params ...interface{}) ICommand {
	command := &command{executeFunc: executeFunc, params: params}
	return command
}

package http

type httpCommand struct {
	context *HttpContext
	server  *HttpServer
}

func (this *httpCommand) Execute() {
	this.server.CallFunc(this.context)
}

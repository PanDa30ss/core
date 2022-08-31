package http

type httpCommand struct {
	context *HttpContext
	server  *HttpServer
}

func (this *httpCommand) Execute() {
	this.server.CallFunc(this.context)
}

type httpCallBackFunc func(result *HttpResult, params ...interface{})
type httpCallBack struct {
	callback httpCallBackFunc
	result   *HttpResult
	params   []interface{}
}

func (this *httpCallBack) Execute() {
	if this.callback == nil {
		return
	}
	this.callback(this.result, this.params...)
}

func defaultFunc(result *HttpResult, params ...interface{}) {
}

func makeHttpCallback(callback httpCallBackFunc, params ...interface{}) *httpCallBack {
	ret := &httpCallBack{}
	ret.callback = callback
	ret.params = params
	return ret
}

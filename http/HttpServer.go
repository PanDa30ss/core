package http

import (
	log "core/logManager"
	"net/http"
	"time"
)

type HttpServer struct {
	httpParam         *httpParams
	httpsParam        *httpParams
	handleHandleFuncs map[string]HttpHandleFunc
}

type HttpHandleFunc func(context *HttpContext)

type httpParams struct {
	serveMux *http.ServeMux
	address  string
	certFile string
	keyFile  string
}

func (this *HttpServer) Init(address string, opts ...string) {
	var param *httpParams = &httpParams{}
	param.serveMux = http.NewServeMux()
	param.address = address
	for url, _ := range this.handleHandleFuncs {
		param.serveMux.Handle(url, &httpHandler{dt: time.Second * 10, server: this})
	}
	if len(opts) == 0 {
		this.httpParam = param
	} else {
		param.certFile = opts[0]
		param.keyFile = opts[1]
		this.httpsParam = param
	}

}

func (this *HttpServer) Start() bool {
	if this.httpParam != nil {
		go http.ListenAndServe(this.httpParam.address, this.httpParam.serveMux)
	}
	if this.httpsParam != nil {
		go http.ListenAndServeTLS(this.httpsParam.address, this.httpsParam.certFile, this.httpsParam.keyFile, this.httpsParam.serveMux)
	}
	return true
}

func (this *HttpServer) Register(url string, handleFunc HttpHandleFunc) bool {
	// this.handleHandleFuncs["/"+url] = handleFunc
	this.handleHandleFuncs[url] = handleFunc
	return true
}

func (this *HttpServer) CallFunc(context *HttpContext) {
	defer func() {
		if e := recover(); e != nil {
			log.Info(e)
			context.Finish()
		}
	}()
	url := context.r.RequestURI
	if _, ok := this.handleHandleFuncs[url]; ok {
		this.handleHandleFuncs[url](context)
	}
	// this.handleHandleFuncs[url](context)
}

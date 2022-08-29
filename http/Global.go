package http

func MakeHttpServer() *HttpServer {
	return &HttpServer{handleHandleFuncs: make(map[string]HttpHandleFunc)}
}

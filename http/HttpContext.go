package http

import (
	"net/http"
)

type HttpContext struct {
	done chan bool
	w    http.ResponseWriter
	r    *http.Request
}

func (this *HttpContext) Write(buff []byte) {
	this.w.Write(buff)
}

func (this *HttpContext) Done() chan bool {
	return this.done
}
func (this *HttpContext) Finish() {

	close(this.done)
}

func makeHttpContext(w http.ResponseWriter, r *http.Request) *HttpContext {
	return &HttpContext{done: make(chan bool), w: w, r: r}
}

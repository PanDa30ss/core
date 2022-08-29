package http

import (
	"core/service"
	"net/http"
	"time"
)

type httpHandler struct {
	dt     time.Duration
	server *HttpServer
}

func (this *httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context := makeHttpContext(w, r)
	var cmd = httpCommand{context, this.server}
	service.Post(&cmd)
	select {
	case <-context.Done():
		return

	case <-time.After(this.dt):
		context.Finish()
		return
	}
}

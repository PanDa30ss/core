package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PanDa30ss/core/service"
)

type httpHandler struct {
	dt     time.Duration
	server *HttpServer
}

func (this *httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(2222)
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

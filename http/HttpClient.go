package http

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/PanDa30ss/core/service"
)

var HTTP_CLIENT_TIMEOUT time.Duration = 5

func Get(url string, f func(result *HttpResult, params ...interface{}), params ...interface{}) {
	go doGet(url, f, params...)
}

func doGet(url string, f func(result *HttpResult, params ...interface{}), params ...interface{}) {
	client := &http.Client{Timeout: HTTP_CLIENT_TIMEOUT * time.Second}
	resp, err := client.Get(url)
	ret := &HttpResult{}
	callback := makeHttpCallback(f, params...)
	callback.result = ret
	if err != nil {
		ret.Err = err
		service.Post(callback)
		return
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			ret.Err = err
			service.Post(callback)
			return
		}
	}
	ret.Err = err
	ret.Result = result.String()
	service.Post(callback)
}

func Post(url string, data string, contentType string, f func(result *HttpResult, params ...interface{}), params ...interface{}) {
	go doPost(url, data, contentType, f, params...)
}

func doPost(url string, data string, contentType string, f func(result *HttpResult, params ...interface{}), params ...interface{}) {
	client := &http.Client{Timeout: HTTP_CLIENT_TIMEOUT * time.Second}
	resp, err := client.Post(url, contentType, bytes.NewBuffer([]byte(data)))
	ret := &HttpResult{}
	callback := makeHttpCallback(f, params...)
	callback.result = ret
	if err != nil {
		ret.Err = err
		service.Post(callback)
	}
	defer resp.Body.Close()

	result, readErr := ioutil.ReadAll(resp.Body)
	ret.Err = readErr
	ret.Result = string(result)
	service.Post(callback)

}

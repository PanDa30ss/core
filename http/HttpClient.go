package http

import (
	"bytes"
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
	defer resp.Body.Close()

	ret := &HttpResult{}
	callback := makeHttpCallback(f, params...)
	callback.result = ret
	if err != nil {
		ret.Err = err
		service.GoPost(callback)
		return
	}
	result, readErr := ioutil.ReadAll(resp.Body)
	ret.Err = readErr
	ret.Result = string(result)
	service.Post(callback)
}

func Post(url string, data string, contentType string, f func(result *HttpResult, params ...interface{}), params ...interface{}) {
	go doPost(url, data, contentType, f, params...)
}

func doPost(url string, data string, contentType string, f func(result *HttpResult, params ...interface{}), params ...interface{}) {
	client := &http.Client{Timeout: HTTP_CLIENT_TIMEOUT * time.Second}
	resp, err := client.Post(url, contentType, bytes.NewBuffer([]byte(data)))
	defer resp.Body.Close()
	ret := &HttpResult{}
	callback := makeHttpCallback(f, params...)
	callback.result = ret
	if err != nil {
		ret.Err = err
		service.Post(callback)
		return
	}

	result, readErr := ioutil.ReadAll(resp.Body)
	ret.Err = readErr
	ret.Result = string(result)
	service.Post(callback)

}

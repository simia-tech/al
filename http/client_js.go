// +build js

package http

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	gohttp "net/http"
	"strings"
	"syscall/js"

	"github.com/simia-tech/errx"
)

func (c *Client) Do(request *gohttp.Request) (*gohttp.Response, error) {
	done := make(chan struct{})

	requestElement := js.Global().Get("XMLHttpRequest").New()
	requestElement.Set("responseType", "arraybuffer")

	loadCb := js.NewEventCallback(0, func(event js.Value) {
		done <- struct{}{}
	})
	defer loadCb.Release()
	requestElement.Set("onload", loadCb)

	errorCb := js.NewEventCallback(0, func(event js.Value) {
		log.Printf("error")
		done <- struct{}{}
	})
	defer errorCb.Release()
	requestElement.Set("onerror", errorCb)

	requestElement.Call("open", request.Method, request.URL.String())
	for key, values := range request.Header {
		for _, value := range values {
			requestElement.Call("setRequestHeader", key, value)
		}
	}
	if m := strings.ToUpper(request.Method); (m == gohttp.MethodPost || m == gohttp.MethodPut) && request.Body != nil {
		data, err := ioutil.ReadAll(request.Body)
		if err != nil {
			return nil, errx.Annotatef(err, "read request body")
		}
		if err = request.Body.Close(); err != nil {
			return nil, errx.Annotatef(err, "close request body")
		}
		ta := js.TypedArrayOf([]uint8(data))
		defer ta.Release()
		requestElement.Call("send", ta)
	} else {
		requestElement.Call("send")
	}

	<-done

	responseStatus := fmt.Sprintf("HTTP/1.1 %d %s\n", requestElement.Get("status").Int(), requestElement.Get("statusText").String())
	responseHeader := requestElement.Call("getAllResponseHeaders").String() + "\n"

	responseElement := requestElement.Get("response")
	responseView := js.Global().Get("Uint8Array").New(responseElement)

	data := make([]byte, responseView.Length())
	for index := 0; index < responseView.Length(); index++ {
		data[index] = byte(responseView.Index(index).Int())
	}

	responseReader := bufio.NewReader(io.MultiReader(
		bytes.NewBufferString(responseStatus),
		bytes.NewBufferString(responseHeader),
		bytes.NewBuffer(data)))
	response, err := gohttp.ReadResponse(responseReader, request)
	if err != nil {
		return nil, errx.Annotatef(err, "read response")
	}

	return response, nil
}

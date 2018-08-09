// +build js

package http_test

import (
	"os"
	"syscall/js"
	"testing"
)

func TestMain(m *testing.M) {
	beforeUnloadChan := make(chan struct{})

	beforeUnloadCb := js.NewEventCallback(0, func(event js.Value) {
		beforeUnloadChan <- struct{}{}
	})
	defer beforeUnloadCb.Release()
	addEventListener := js.Global().Get("addEventListener")
	addEventListener.Invoke("beforeunload", beforeUnloadCb)

	result := m.Run()

	<-beforeUnloadChan

	os.Exit(result)
}

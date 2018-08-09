package http_test

import (
	gohttp "net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/simia-tech/al/http"
)

func TestClientGet(t *testing.T) {
	request, err := gohttp.NewRequest(gohttp.MethodGet, "https://httpbin.org/get", nil)
	require.NoError(t, err)

	response, err := http.DefaultClient.Do(request)
	require.NoError(t, err)
	defer response.Body.Close()

	assert.Equal(t, 200, response.StatusCode)
}

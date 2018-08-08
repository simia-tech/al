package local_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/simia-tech/al/storage/local"
)

func TestSetAndGet(t *testing.T) {
	s, err := local.NewStorage("test")
	require.NoError(t, err)

	require.NoError(t, s.Set("one", "value"))
	value, err := s.Get("one")
	require.NoError(t, err)
	assert.Equal(t, "value", value)
}

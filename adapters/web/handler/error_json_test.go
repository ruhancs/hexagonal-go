package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandler_JsonError(t *testing.T) {
	msg := "Hello json"
	result := JsonError(msg)
	require.Equal(t, []byte(`{"message":"Hello json"}`), result)
}
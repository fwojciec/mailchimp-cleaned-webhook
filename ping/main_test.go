package main

import (
	"context"
	"testing"

	"github.com/matryer/is"
)

func TestPingHandler(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	h := handler{}
	res, err := h.Handler(context.Background())
	is.NoErr(err)                 // expected no error
	is.Equal(res.StatusCode, 200) // expected 200 status code
}

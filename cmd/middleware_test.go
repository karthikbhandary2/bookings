package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/karthikbhandary2/bookings/cmd/web/middlewares"
)

func TestNoSurf(t *testing.T) {
	var myH MyHandler
	h := middlewares.NoSurf(&myH)

	switch v := h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Error(fmt.Sprintf("type is not http.Handler but is %T", v))
	}
}

func TestSessionLoad(t *testing.T) {
	var myH MyHandler
	h := middlewares.SessionLoad(&myH)

	switch v := h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Error(fmt.Sprintf("type is not http.Handler but is %T", v))
	}
}
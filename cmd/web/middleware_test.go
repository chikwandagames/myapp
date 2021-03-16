package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var myH myHandler
	h := NoSurf(&myH)

	// All we need to test if if myH is actually a handler

	// Assign v the type of h
	switch v := h.(type) {
	// if h is of type handler
	case http.Handler:
		// do nothing as the test has been passed

	default:
		t.Error(fmt.Sprintf("type is not http.Handler %T", v))
	}
}

func TestSessionLoad(t *testing.T) {
	var myH myHandler
	h := SessionLoad(&myH)

	// All we need to test if if myH is actually a handler

	// Assign v the type of h
	switch v := h.(type) {
	// if h is of type handler
	case http.Handler:
		// do nothing as the test has been passed

	default:
		t.Error(fmt.Sprintf("type is not http.Handler %T", v))
	}
}

func TestWriteToConsole(t *testing.T) {
	var myH myHandler
	h := WriteToConsole(&myH)

	// All we need to test if if myH is actually a handler

	// Assign v the type of h
	switch v := h.(type) {
	// if h is of type handler
	case http.Handler:
		// do nothing as the test has been passed

	default:
		t.Error(fmt.Sprintf("type is not http.Handler %T", v))
	}
}

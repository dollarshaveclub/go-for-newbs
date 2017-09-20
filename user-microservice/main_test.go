package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

//var benchreq *http.Request

func TestGetHandlerBadMethod(t *testing.T) {
	req, _ := http.NewRequest("PUT", "http://foo.bar.com/", bytes.NewBuffer(nil))
	w := httptest.NewRecorder()
	getHandler(w, req)
	if w.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected to fail! bad status: %v\n", w.Code)
	}
}

func BenchmarkGetHandler(b *testing.B) {
	req, _ := http.NewRequest("GET", "http://foo.bar.com/", bytes.NewBuffer(nil))
	w := httptest.NewRecorder()
	getHandler(w, req)
}

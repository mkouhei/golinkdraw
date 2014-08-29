package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeHome(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(serveHome))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if res.StatusCode != 200 {
		t.Error("Status code error")
		return
	}
}

package test

import (
	"net/http"
	"net/http/httptest"
)

func Setup() (
	httpclient *http.Client, mux *http.ServeMux, server *httptest.Server, teardown func(),
) {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	teardown = server.Close
	httpclient = server.Client()
	return
}

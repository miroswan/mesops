package v1

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/miroswan/mesops/test"
)

func MasterSetup() (
	master *Master, mux *http.ServeMux, teardown func(),
) {
	var httpclient *http.Client
	var server *httptest.Server
	httpclient, mux, server, teardown = test.Setup()
	master, _ = NewMasterBuilder(server.URL).SetHTTPClient(httpclient).Build()
	return
}

func AgentSetup() (
	agent *Agent, mux *http.ServeMux, teardown func(),
) {
	var httpclient *http.Client
	var server *httptest.Server
	httpclient, mux, server, teardown = test.Setup()
	agent, _ = NewAgentBuilder(server.URL).SetHTTPClient(httpclient).Build()
	return
}

func SetOutput(mux *http.ServeMux, output string) {
	mux.HandleFunc("/api/v1", func(rw http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodPost {
			fmt.Fprint(rw, output)
		}
	})
}

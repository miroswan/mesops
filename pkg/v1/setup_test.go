package v1

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/gogo/protobuf/proto"
	"github.com/mesos/go-proto/mesos/v1/agent"
	"github.com/mesos/go-proto/mesos/v1/master"
)

type ClientType int

const (
	MasterClient ClientType = iota
	AgentClient
)

type TestProtobufServer struct {
	mux             *http.ServeMux
	httpClient      *http.Client
	httpServer      *httptest.Server
	master          *Master
	mesos_v1_agent  *Agent
	clientType      ClientType
	input           []byte
	output          []byte
	ctx             context.Context
	cancelFunc      context.CancelFunc
	closeServerFunc func()
}

func NewTestProtobufServer(clientType ClientType) *TestProtobufServer {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	httpClient := server.Client()
	master, _ := NewMasterBuilder(server.URL).SetHTTPClient(httpClient).Build()
	mesos_v1_agent, _ := NewAgentBuilder(server.URL).SetHTTPClient(httpClient).Build()
	ctx, cancelFunc := context.WithCancel(context.Background())
	return &TestProtobufServer{
		mux:             mux,
		clientType:      clientType,
		httpServer:      server,
		httpClient:      httpClient,
		master:          master,
		mesos_v1_agent:  mesos_v1_agent,
		ctx:             ctx,
		cancelFunc:      cancelFunc,
		closeServerFunc: server.Close,
	}
}

// Teardown shuts down the test server then closes the context. Call this in a
// defer statement at the top of your tests
func (t *TestProtobufServer) Teardown() {
	t.closeServerFunc()
	t.cancelFunc()
}

func (t *TestProtobufServer) Master() *Master      { return t.master }
func (t *TestProtobufServer) Agent() *Agent        { return t.mesos_v1_agent }
func (t *TestProtobufServer) Ctx() context.Context { return t.ctx }

func (t *TestProtobufServer) SetOutput(b []byte) *TestProtobufServer {
	t.output = b
	return t
}

// Handle does not block. It attaches a generic HandleFunc to the http.ServeMux.
func (t *TestProtobufServer) Handle() {
	t.mux.HandleFunc("/api/v1", func(rw http.ResponseWriter, req *http.Request) {
		// Default error response
		serverErrorResponse := func(err error) {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(fmt.Sprintf("500 - %s", err)))
		}

		clientErrorResponse := func(err error) {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte(fmt.Sprintf("400 - %s", err)))
		}

		// Only do things in POST
		if req.Method == http.MethodPost {
			// Process intput
			if len(t.input) > 0 {
				// Read the request body
				b, err := ioutil.ReadAll(req.Body)
				if err != nil {
					clientErrorResponse(err)
				}
				// Umarshall the request into a Response
				switch t.clientType {
				case MasterClient:
					err = proto.Unmarshal(b, &mesos_v1_master.Call{})
				case AgentClient:
					err = proto.Unmarshal(b, &mesos_v1_agent.Call{})
				}
				// If there is an error, send error response
				if err != nil {
					clientErrorResponse(err)
				}
			}

			if len(t.output) > 0 {
				// Send response until there is nothing left to send
				rw.Header().Set("Content-Type", "application/x-protobuf")
				written, err := rw.Write(t.output)
				if err != nil {
					serverErrorResponse(err)
				}
				for written < len(t.output) {
					remainingSize := len(t.output) - written
					remaining := make([]byte, remainingSize)
					written, _ = rw.Write(remaining)
				}
			}
		}
	})
}

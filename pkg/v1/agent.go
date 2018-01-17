// MIT License
//
// Copyright (c) [2017-2018] [Demitri Swan]
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package v1

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/gogo/protobuf/proto"
	"github.com/mesos/go-proto/mesos/v1/agent"
)

// Agent is a struct that handles most interactions with the
// Mesos Operator Agent HTTP API. Build an Agent with an AgentBuilder.
type Agent struct {
	*client
}

// AgentBuilder is a builder that takes some manditory parameters and
// allows you to set optional parameters via its set methods. Call Build to
// return the final constructed struct. Create an AgentBuilder with
// NewAgentBuilder
type AgentBuilder struct {
	*clientBuilder
}

// NewAgentBuilder returns a pointer to an AgentBuilder. The serverURL is the
// base URL of the mesos_v1_agent, including the SCHEMA://FQDN_OR_IP:PORT
func NewAgentBuilder(serverURL string) *AgentBuilder {
	return &AgentBuilder{clientBuilder: newClientBuilder(serverURL)}
}

// SetHTTPClient sets the *http.Client for the Agent and returns a pointer
// to the AgentBuilder. If SetHTTPClient is not called, the Build method
// will use an http.Defaultclient.
//
// e.g.
//
// 	var b *AgentBuilder = NewAgentBuilder("https://127.0.0.1:5051").SetHTTPClient(myCustomclient)
func (b *AgentBuilder) SetHTTPClient(httpclient *http.Client) *AgentBuilder {
	b.clientBuilder.setHTTPclient(httpclient)
	return b
}

// SetMaxRetries sets maxRetries for the Agent and returns a pointer to an
// AgentBuilder. If SetMaxRetries is not called, it will be set to 15 seconds.
// Each HTTP request will retry up to the provided value upon failure.
// Binary exponential backoff is implemented as specified by RFC2616
//
// e.g.
//
// 	var b *AgentBuilder = NewAgentBuilder("https://127.0.0.1:5051").SetMaxRetries(5)
func (b *AgentBuilder) SetMaxRetries(maxRetries int) *AgentBuilder {
	b.clientBuilder.setMaxRetries(maxRetries)
	return b
}

// Build returns a pointer to a constructed Agent.
func (b *AgentBuilder) Build() (a *Agent, err error) {
	var client *client
	client, err = b.clientBuilder.build()
	if err != nil {
		return
	}
	a = &Agent{client: client}
	return
}

// sendSimpleCall configures a simple mesos_v1_agent.Call, marshalls it into binary format,
// and sends it over HTTP to the configured mesos_v1_agent. These calls don't need
// additional configuration other than the mesos_v1_agent.Call_Type
func (a *Agent) sendSimpleCall(ctx context.Context, callType mesos_v1_agent.Call_Type) (
	response *mesos_v1_agent.Response, httpResponse *http.Response, err error,
) {
	var callMsg proto.Message = &mesos_v1_agent.Call{Type: &callType}
	var b []byte
	b, err = proto.Marshal(callMsg)
	if err != nil {
		return
	}
	var buf io.Reader = bytes.NewBuffer(b)
	response = &mesos_v1_agent.Response{}
	httpResponse, err = a.client.doProtoWrapper(ctx, buf, response)
	return
}

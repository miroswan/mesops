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
	"github.com/mesos/go-proto/mesos/v1/master"
)

// MasterBuilder is a builder that takes some manditory parameters and
// allows you to set optional parameters via its set methods. Call Build to
// return the final constructed struct. Create an MasterBuilder with
// NewMasterBuilder
type MasterBuilder struct {
	*clientBuilder
}

// NewMasterBuilder returns a pointer to an MasterBuilder. The serverURL is the
// base URL of the mesos_v1_agent, including the SCHEMA://FQDN_OR_IP:PORT
func NewMasterBuilder(serverURL string) *MasterBuilder {
	return &MasterBuilder{clientBuilder: newClientBuilder(serverURL)}
}

// SetHTTPClient sets the *http.Client for the Master and returns a pointer
// to the MasterBuilder. If SetHTTPClient is not called, the Build method
// will use an http.Defaultclient.
//
// e.g.
//
// 	var b *MasterBuilder = NewMasterBuilder("https://127.0.0.1:5051").SetHTTPClient(myCustomclient)
func (b *MasterBuilder) SetHTTPClient(httpclient *http.Client) *MasterBuilder {
	b.clientBuilder.setHTTPclient(httpclient)
	return b
}

// SetMaxRetries sets maxRetries for the Agent and returns a pointer to an
// MasterBuilder. If SetMaxRetries is not called, it will be set to 15 seconds.
// Each HTTP request will retry up to the provided value upon failure.
// Binary exponential backoff is implemented as specified by RFC2616
//
// e.g.
//
// 	var b *MasterBuilder = NewMasterBuilder("https://127.0.0.1:5051").SetMaxRetries(5)
func (b *MasterBuilder) SetMaxRetries(maxRetries int) *MasterBuilder {
	b.clientBuilder.setMaxRetries(maxRetries)
	return b
}

// Build returns a pointer to a constructed Master.
func (b *MasterBuilder) Build() (m *Master, err error) {
	var client *client
	client, err = b.clientBuilder.build()
	if err != nil {
		return
	}
	m = &Master{client: client}
	return
}

// Master is a struct that handles most interactions with the
// Mesos Operator Agent HTTP API. Build an Master with an MasterBuilder.
type Master struct {
	*client
}

// sendSimpleCall configures a simple mesos_v1_master.Call, marshalls it into binary format,
// and sends it over HTTP to the configured mesos_v1_master. These calls don't need
// additional configuration other than the mesos_v1_master.Call_Type
func (m *Master) sendSimpleCall(ctx context.Context, callType mesos_v1_master.Call_Type) (
	response *mesos_v1_master.Response, httpResponse *http.Response, err error,
) {
	var callMsg proto.Message = &mesos_v1_master.Call{Type: &callType}
	var b []byte
	b, err = proto.Marshal(callMsg)
	if err != nil {
		return
	}
	var buf io.Reader = bytes.NewBuffer(b)
	response = &mesos_v1_master.Response{}
	httpResponse, err = m.client.doProtoWrapper(ctx, buf, response)
	return
}

// GetMaster retrieves information about the mesos_v1_master.
func (m *Master) GetMaster(ctx context.Context) (response *mesos_v1_master.Response, err error) {
	response, _, err = m.sendSimpleCall(ctx, mesos_v1_master.Call_GET_MASTER)
	return
}

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
	"context"
	"net/http"

	"github.com/gogo/protobuf/proto"
	"github.com/mesos/go-proto/mesos/v1/agent"
	"github.com/mesos/go-proto/mesos/v1/master"
)

// GetLoggingLevel retrieves the master’s logging level.
func (m *Master) GetLoggingLevel(ctx context.Context) (response *mesos_v1_master.Response, err error) {
	var httpResponse *http.Response
	response, httpResponse, err = m.sendSimpleCall(ctx, mesos_v1_master.Call_GET_LOGGING_LEVEL)
	defer httpResponse.Body.Close()
	return
}

// GetLoggingLevel retrieves the agent's logging level.
func (a *Agent) GetLoggingLevel(ctx context.Context) (response *mesos_v1_agent.Response, err error) {
	var httpResponse *http.Response
	response, httpResponse, err = a.sendSimpleCall(ctx, mesos_v1_agent.Call_GET_LOGGING_LEVEL)
	defer httpResponse.Body.Close()
	return
}

// SetLoggingLevel sets the logging verbosity level for a specified duration for
// mesos_v1_master. Mesos uses glog for logging. The library only uses verbose logging
// which means nothing will be output unless the verbosity level is set
// (by default it’s 0, libprocess uses levels 1, 2, and 3).
func (m *Master) SetLoggingLevel(ctx context.Context, call *mesos_v1_master.Call_SetLoggingLevel) (err error) {
	var callType mesos_v1_master.Call_Type = mesos_v1_master.Call_SET_LOGGING_LEVEL
	var message proto.Message = &mesos_v1_master.Call{Type: &callType, SetLoggingLevel: call}
	var httpResponse *http.Response
	httpResponse, err = m.client.makeCall(ctx, message, nil)
	defer httpResponse.Body.Close()
	return
}

// SetLoggingLevel sets the logging verbosity level for a specified duration for
// agent. Mesos uses glog for logging. The library only uses verbose logging
// which means nothing will be output unless the verbosity level is set
// (by default it’s 0, libprocess uses levels 1, 2, and 3).
func (a *Agent) SetLoggingLevel(ctx context.Context, call *mesos_v1_agent.Call_SetLoggingLevel) (err error) {
	var callType mesos_v1_agent.Call_Type = mesos_v1_agent.Call_SET_LOGGING_LEVEL
	var message proto.Message = &mesos_v1_agent.Call{Type: &callType, SetLoggingLevel: call}
	var httpResponse *http.Response
	httpResponse, err = a.client.makeCall(ctx, message, nil)
	defer httpResponse.Body.Close()
	return
}

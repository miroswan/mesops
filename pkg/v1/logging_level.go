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

	"github.com/gogo/protobuf/proto"
	"github.com/miroswan/mesops/pkg/v1/agent"
	"github.com/miroswan/mesops/pkg/v1/master"
)

// GetLoggingLevel returns a pointer to a GetLoggingLevel.
//
// References:
//
// 	* http://mesos.apache.org/documentation/latest/operator-http-api/#get_logging_level
func (m *Master) GetLoggingLevel(ctx context.Context) (response *master.Response, err error) {
	response, _, err = m.sendSimpleCall(ctx, master.Call_GET_LOGGING_LEVEL)
	return
}

// GetLoggingLevel returns a pointer to a GetLoggingLevel.
//
// References:
//
// 	* http://mesos.apache.org/documentation/latest/operator-http-api/#get_logging_level-1
func (a *Agent) GetLoggingLevel(ctx context.Context) (response *agent.Response, err error) {
	response, _, err = a.sendSimpleCall(ctx, agent.Call_GET_LOGGING_LEVEL)
	return
}

// SetLoggingLevel sets the logging level on the configured Master
//
// References:
// * http://mesos.apache.org/documentation/latest/operator-http-api/#set_logging_level
func (m *Master) SetLoggingLevel(ctx context.Context, call *master.Call_SetLoggingLevel) (err error) {
	var callType master.Call_Type = master.Call_SET_LOGGING_LEVEL
	var payload proto.Message = &master.Call{Type: &callType, SetLoggingLevel: call}
	var b []byte

	b, err = proto.Marshal(payload)
	if err != nil {
		return
	}
	var buf io.Reader = bytes.NewBuffer(b)
	_, err = m.client.doProtoWrapper(ctx, buf, nil)
	return
}

// SetLoggingLevel sets the logging level on the configured Agent
//
// References:
// * http://mesos.apache.org/documentation/latest/operator-http-api/#set_logging_level-1
func (a *Agent) SetLoggingLevel(ctx context.Context, call *agent.Call_SetLoggingLevel) (err error) {
	var callType agent.Call_Type = agent.Call_SET_LOGGING_LEVEL
	var payload proto.Message = &agent.Call{Type: &callType, SetLoggingLevel: call}
	var b []byte

	b, err = proto.Marshal(payload)
	if err != nil {
		return
	}
	var buf io.Reader = bytes.NewBuffer(b)
	_, err = a.client.doProtoWrapper(ctx, buf, nil)
	return
}

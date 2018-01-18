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

// GetContainers retrieves information about containers running on this agent.
// It contains ContainerStatus and ResourceStatistics along with some metadata
// of the containers.
func (a *Agent) GetContainers(ctx context.Context) (response *mesos_v1_agent.Response, err error) {
	var httpResponse *http.Response
	response, httpResponse, err = a.sendSimpleCall(ctx, mesos_v1_agent.Call_GET_CONTAINERS)
	defer httpResponse.Body.Close()
	return
}

// LaunchNestedContainer launches a nested container. Any authorized entity,
// including the executor itself, its tasks, or the operator can use this API to
// launch a nested container.
func (a *Agent) LaunchNestedContainer(ctx context.Context, call *mesos_v1_agent.Call_LaunchNestedContainer) (err error) {
	var callType mesos_v1_agent.Call_Type = mesos_v1_agent.Call_LAUNCH_NESTED_CONTAINER
	var payload proto.Message = &mesos_v1_agent.Call{Type: &callType, LaunchNestedContainer: call}
	var b []byte
	b, err = proto.Marshal(payload)
	if err != nil {
		return
	}
	var buf io.Reader = bytes.NewBuffer(b)
	var httpResponse *http.Response
	httpResponse, err = a.client.doProtoWrapper(ctx, buf, nil)
	defer httpResponse.Body.Close()
	return
}

// WaitNestedContainer waits for a nested container to terminate or exit. Any
// authorized entity, including the executor itself, its tasks, or the operator
// can use this API to wait on a nested container.
func (a *Agent) WaitNestedContainer(ctx context.Context, call *mesos_v1_agent.Call_WaitNestedContainer) (
	response *mesos_v1_agent.Response, err error,
) {
	// Build message
	var callType mesos_v1_agent.Call_Type = mesos_v1_agent.Call_WAIT_NESTED_CONTAINER
	var payload proto.Message = &mesos_v1_agent.Call{Type: &callType, WaitNestedContainer: call}
	var b []byte
	// Encode to protobuf
	b, err = proto.Marshal(payload)
	if err != nil {
		return
	}
	var buf io.Reader = bytes.NewBuffer(b)
	response = &mesos_v1_agent.Response{}
	// Send HTTP Request
	_, err = a.client.doProtoWrapper(ctx, buf, response)
	return
}

// KillNestedContainer initiates the destruction of a nested container. Any
// authorized entity, including the executor itself, its tasks, or the operator
// can use this API to kill a nested container.
func (a *Agent) KillNestedContainer(ctx context.Context, call *mesos_v1_agent.Call_KillNestedContainer) (err error) {
	var callType mesos_v1_agent.Call_Type = mesos_v1_agent.Call_KILL_NESTED_CONTAINER
	var payload proto.Message = &mesos_v1_agent.Call{Type: &callType, KillNestedContainer: call}
	var b []byte
	b, err = proto.Marshal(payload)
	if err != nil {
		return
	}
	var buf io.Reader = bytes.NewBuffer(b)
	_, err = a.client.doProtoWrapper(ctx, buf, nil)
	return
}

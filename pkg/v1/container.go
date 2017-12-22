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
)

// GetContainers returns a pointer to a GetContainers
//
// References:
//
// 	* http://mesos.apache.org/documentation/latest/operator-http-api/#get_containers
func (a *Agent) GetContainers(ctx context.Context) (response *agent.Response, err error) {
	response, _, err = a.sendSimpleCall(ctx, agent.Call_GET_CONTAINERS)
	return
}

// LaunchNestedContainer launches a nested container on the configured Agent
//
// References:
//
// 	* http://mesos.apache.org/documentation/latest/operator-http-api/#launch_nested_container
func (a *Agent) LaunchNestedContainer(ctx context.Context, call *agent.Call_LaunchNestedContainer) (err error) {
	var callType agent.Call_Type = agent.Call_LAUNCH_NESTED_CONTAINER
	var payload proto.Message = &agent.Call{Type: &callType, LaunchNestedContainer: call}
	var b []byte
	b, err = proto.Marshal(payload)
	if err != nil {
		return
	}
	var buf io.Reader = bytes.NewBuffer(b)
	_, err = a.client.doProtoWrapper(ctx, buf, nil)
	return
}

// WaitNestedContainer waits for a nested container to terminate or exit. Any
// authorized entity, including the executor itself, its tasks, or the operator
// can use this API to wait on a nested container.
//
// References:
//
// 	* http://mesos.apache.org/documentation/latest/operator-http-api/#wait_nested_container
func (a *Agent) WaitNestedContainer(ctx context.Context, call agent.Call_WaitNestedContainer) (
	response *agent.Response, err error,
) {
	// Build message
	var callType agent.Call_Type = agent.Call_WAIT_NESTED_CONTAINER
	var payload proto.Message = &agent.Call{Type: &callType, WaitNestedContainer: &call}
	var b []byte
	// Encode to protobuf
	b, err = proto.Marshal(payload)
	if err != nil {
		return
	}
	var buf io.Reader = bytes.NewBuffer(b)
	response = &agent.Response{}
	// Send HTTP Request
	_, err = a.client.doProtoWrapper(ctx, buf, response)
	return
}

// KillNestedContainer initiates the destruction of a nested container. Any
// authorized entity, including the executor itself, its tasks, or the operator
// can use this API to kill a nested container.
//
// References:
//
// 	* http://mesos.apache.org/documentation/latest/operator-http-api/#kill_nested_container
func (a *Agent) KillNestedContainer(ctx context.Context, call agent.Call_KillNestedContainer) (err error) {
	var callType agent.Call_Type = agent.Call_KILL_NESTED_CONTAINER
	var payload proto.Message = &agent.Call{Type: &callType, KillNestedContainer: &call}
	var b []byte
	b, err = proto.Marshal(payload)
	if err != nil {
		return
	}
	var buf io.Reader = bytes.NewBuffer(b)
	_, err = a.client.doProtoWrapper(ctx, buf, nil)
	return
}

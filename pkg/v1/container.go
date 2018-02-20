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
	"bufio"
	"context"
	"net/http"

	"github.com/gogo/protobuf/proto"

	"github.com/mesos/go-proto/mesos/v1/agent"
)

type ProcessIOStream chan *mesos_v1_agent.ProcessIO

// GetContainers retrieves information about containers running on this agent.
// It contains ContainerStatus and ResourceStatistics along with some metadata
// of the containers.
func (a *Agent) GetContainers(ctx context.Context) (response *mesos_v1_agent.Response, err error) {
	var httpResponse *http.Response
	response, httpResponse, err = a.sendSimpleCall(ctx, mesos_v1_agent.Call_GET_CONTAINERS)
	defer httpResponse.Body.Close()
	return
}

// LaunchContainer launches a nested container. Any authorized entity,
// including the executor itself, its tasks, or the operator can use this API to
// launch a nested container.
func (a *Agent) LaunchContainer(ctx context.Context, call *mesos_v1_agent.Call_LaunchContainer) (err error) {
	var callType mesos_v1_agent.Call_Type = mesos_v1_agent.Call_LAUNCH_CONTAINER
	var message proto.Message = &mesos_v1_agent.Call{Type: &callType, LaunchContainer: call}
	var httpResponse *http.Response
	httpResponse, err = a.client.makeCall(ctx, message, nil)
	defer httpResponse.Body.Close()
	return
}

// LaunchNestedContainer launches a nested container. Any authorized entity,
// including the executor itself, its tasks, or the operator can use this API to
// launch a nested container.
func (a *Agent) LaunchNestedContainer(ctx context.Context, call *mesos_v1_agent.Call_LaunchNestedContainer) (err error) {
	var callType mesos_v1_agent.Call_Type = mesos_v1_agent.Call_LAUNCH_NESTED_CONTAINER
	var message proto.Message = &mesos_v1_agent.Call{Type: &callType, LaunchNestedContainer: call}
	var httpResponse *http.Response
	httpResponse, err = a.client.makeCall(ctx, message, nil)
	defer httpResponse.Body.Close()
	return
}

// WaitNestedContainer waits for a nested container to terminate or exit. Any
// authorized entity, including the executor itself, its tasks, or the operator
// can use this API to wait on a nested container.
func (a *Agent) WaitNestedContainer(ctx context.Context, call *mesos_v1_agent.Call_WaitNestedContainer) (
	response *mesos_v1_agent.Response, err error,
) {
	var httpResponse *http.Response
	// Build message
	var callType mesos_v1_agent.Call_Type = mesos_v1_agent.Call_WAIT_NESTED_CONTAINER
	var message proto.Message = &mesos_v1_agent.Call{Type: &callType, WaitNestedContainer: call}
	response = &mesos_v1_agent.Response{}
	httpResponse, err = a.client.makeCall(ctx, message, response)
	defer httpResponse.Body.Close()
	return
}

// KillNestedContainer initiates the destruction of a nested container. Any
// authorized entity, including the executor itself, its tasks, or the operator
// can use this API to kill a nested container.
func (a *Agent) KillNestedContainer(ctx context.Context, call *mesos_v1_agent.Call_KillNestedContainer) (err error) {
	var httpResponse *http.Response
	var callType mesos_v1_agent.Call_Type = mesos_v1_agent.Call_KILL_NESTED_CONTAINER
	var message proto.Message = &mesos_v1_agent.Call{Type: &callType, KillNestedContainer: call}
	httpResponse, err = a.client.makeCall(ctx, message, nil)
	defer httpResponse.Body.Close()
	return
}

// LaunchNestedContainerSession launches a nested container whose lifetime is
// tied to the lifetime of the HTTP call establishing this connection. The
// STDOUT and STDERR of the nested container is streamed back to the client so
// long as the connection is active.
func (a *Agent) LaunchNestedContainerSession(
	ctx context.Context, call *mesos_v1_agent.Call_LaunchNestedContainerSession, procesIOStream ProcessIOStream,
) (err error) {
	var httpResponse *http.Response
	var callType mesos_v1_agent.Call_Type = mesos_v1_agent.Call_LAUNCH_NESTED_CONTAINER_SESSION
	var callMsg proto.Message = &mesos_v1_agent.Call{Type: &callType, LaunchNestedContainerSession: call}

	httpResponse, err = a.client.makeCall(ctx, callMsg, nil)
	if err != nil {
		return
	}
	var reader *bufio.Reader = bufio.NewReader(httpResponse.Body)
	defer httpResponse.Body.Close()
	for {
		select {
		case <-ctx.Done():
			err = ctx.Err()
			return
		default:
			var msg []byte
			msg, err = readRecordioMessage(reader)
			if err != nil {
				return
			}

			// Unmarshal data into a mesos_v1_master.Event
			processIO := &mesos_v1_agent.ProcessIO{}
			err = proto.Unmarshal(msg, processIO)
			if err != nil {
				return
			}
			procesIOStream <- processIO
		}
	}
}

// This call attaches to the STDIN of the primary process of a container and
// streams input to it. This call can only be made against containers that have
// been launched with an associated IOSwitchboard (i.e. nested containers
// launched via a LAUNCH_NESTED_CONTAINER_SESSION call or normal containers
// launched with a TTYInfo in their ContainerInfo). Only one
// ATTACH_CONTAINER_INPUT call can be active for a given container at a time.
// Subsequent attempts to attach will fail.
//
// The first message sent over an ATTACH_CONTAINER_INPUT stream must be of type
// CONTAINER_ID and contain the ContainerID of the container being attached to.
// Subsequent messages must be of type PROCESS_IO, but they may contain subtypes
// of either DATA or CONTROL. DATA messages must be of type STDIN and contain
// the actual data to stream to the STDIN of the container being attached to.
// Currently, the only valid CONTROL message sends a heartbeat to keep the
// connection alive. We may add more CONTROL messages in the future.
func (a *Agent) AttachContainerInput(
	ctx context.Context, call *mesos_v1_agent.Call_AttachContainerInput,
	procesIOStream ProcessIOStream,
) (err error) {
	var httpResponse *http.Response
	var callType mesos_v1_agent.Call_Type = mesos_v1_agent.Call_ATTACH_CONTAINER_INPUT
	var callMsg proto.Message = &mesos_v1_agent.Call{Type: &callType, AttachContainerInput: call}

	httpResponse, err = a.client.makeCall(ctx, callMsg, nil)
	if err != nil {
		return
	}
	var reader *bufio.Reader = bufio.NewReader(httpResponse.Body)
	defer httpResponse.Body.Close()
	for {
		select {
		case <-ctx.Done():
			err = ctx.Err()
			return
		default:
			var msg []byte
			msg, err = readRecordioMessage(reader)
			if err != nil {
				return
			}

			// Unmarshal data into a mesos_v1_master.Event
			processIO := &mesos_v1_agent.ProcessIO{}
			err = proto.Unmarshal(msg, processIO)
			if err != nil {
				return
			}
			procesIOStream <- processIO
		}
	}
}

// AtachContainerOutput attaches to the STDOUT and STDERR of the primary process of a
// container and streams its output back to the client. This call can only be
// made against containers that have been launched with an associated
// IOSwitchboard (i.e. nested containers launched via a
// AUNCH_NESTED_CONTAINER_SESSION call or normal containers launched with a
// TTYInfo in their ContainerInfo field). Multiple ATTACH_CONTAINER_OUTPUT
// calls can be active for a given container at once.
func (a *Agent) AttachContainerOutput(
	ctx context.Context, call *mesos_v1_agent.Call_AttachContainerOutput,
	procesIOStream ProcessIOStream,
) (err error) {
	var httpResponse *http.Response
	var callType mesos_v1_agent.Call_Type = mesos_v1_agent.Call_ATTACH_CONTAINER_OUTPUT
	var callMsg proto.Message = &mesos_v1_agent.Call{Type: &callType, AttachContainerOutput: call}

	httpResponse, err = a.client.makeCall(ctx, callMsg, nil)
	if err != nil {
		return
	}
	var reader *bufio.Reader = bufio.NewReader(httpResponse.Body)
	defer httpResponse.Body.Close()
	for {
		select {
		case <-ctx.Done():
			err = ctx.Err()
			return
		default:
			var msg []byte
			msg, err = readRecordioMessage(reader)
			if err != nil {
				return
			}

			// Unmarshal data into a mesos_v1_master.Event
			processIO := &mesos_v1_agent.ProcessIO{}
			err = proto.Unmarshal(msg, processIO)
			if err != nil {
				return
			}
			procesIOStream <- processIO
		}
	}
}

// RemoveNestedContainer initiates the destruction of a nested container. Any
// authorized entity, including the executor itself, its tasks, or the operator
// can use this API to kill a nested container.
func (a *Agent) RemoveNestedContainer(
	ctx context.Context, call *mesos_v1_agent.Call_RemoveNestedContainer,
) (err error) {
	var httpResponse *http.Response
	var callType mesos_v1_agent.Call_Type = mesos_v1_agent.Call_REMOVE_NESTED_CONTAINER
	var message proto.Message = &mesos_v1_agent.Call{Type: &callType, RemoveNestedContainer: call}
	httpResponse, err = a.client.makeCall(ctx, message, nil)
	defer httpResponse.Body.Close()
	return
}

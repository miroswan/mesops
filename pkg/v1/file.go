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
	"github.com/mesos/go-proto/mesos/v1/agent"
	"github.com/mesos/go-proto/mesos/v1/master"
)

// ListFiles retrieves the file listing for a directory in mesos_v1_master. You must
// pass a valid virtual file path. A mapping of the virtual file paths to actual
// paths can be found at the files/debug endpoint
func (m *Master) ListFiles(ctx context.Context, call *mesos_v1_master.Call_ListFiles) (response *mesos_v1_master.Response, err error) {
	var callType mesos_v1_master.Call_Type = mesos_v1_master.Call_LIST_FILES
	var payload proto.Message = &mesos_v1_master.Call{Type: &callType, ListFiles: call}
	var b []byte
	// Encode to protobuf
	b, err = proto.Marshal(payload)
	if err != nil {
		return
	}
	var buf io.Reader = bytes.NewBuffer(b)
	response = &mesos_v1_master.Response{}
	// Send HTTP Request
	_, err = m.client.doProtoWrapper(ctx, buf, response)
	return
}

// ListFiles retrieves the file listing for a directory in agent. You must
// pass a valid virtual file path. A mapping of the virtual file paths to actual
// paths can be found at the files/debug endpoint
func (a *Agent) ListFiles(ctx context.Context, call *mesos_v1_agent.Call_ListFiles) (response *mesos_v1_agent.Response, err error) {
	var callType mesos_v1_agent.Call_Type = mesos_v1_agent.Call_LIST_FILES
	var payload proto.Message = &mesos_v1_agent.Call{Type: &callType, ListFiles: call}
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

// ReadFile retrieves the file listing for a directory in mesos_v1_agent. This call takes
// the path of the file to be read, the offset to start reading, and the maximum
// number of bytes to read. The length member of the ReadFile is optional. The
// path must be a valid virtual file path.
func (m *Master) ReadFile(ctx context.Context, call *mesos_v1_master.Call_ReadFile) (response *mesos_v1_master.Response, err error) {
	var callType mesos_v1_master.Call_Type = mesos_v1_master.Call_READ_FILE
	var payload proto.Message = &mesos_v1_master.Call{Type: &callType, ReadFile: call}
	var b []byte
	// Encode to protobuf
	b, err = proto.Marshal(payload)
	if err != nil {
		return
	}
	var buf io.Reader = bytes.NewBuffer(b)
	response = &mesos_v1_master.Response{}
	// Send HTTP Request
	_, err = m.client.doProtoWrapper(ctx, buf, response)
	return
}

// ReadFile retrieves the file listing for a directory in mesos_v1_agent. This call takes
// the path of the file to be read, the offset to start reading, and the maximum
// number of bytes to read. The length member of the ReadFile is optional. The
// path must be a valid virtual file path.
func (a *Agent) ReadFile(ctx context.Context, call *mesos_v1_agent.Call_ReadFile) (response *mesos_v1_agent.Response, err error) {
	var callType mesos_v1_agent.Call_Type = mesos_v1_agent.Call_READ_FILE
	var payload proto.Message = &mesos_v1_agent.Call{Type: &callType, ReadFile: call}
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

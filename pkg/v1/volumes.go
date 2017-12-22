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
	"github.com/miroswan/mesops/pkg/v1/master"
)

// CreateVolumes creates persistent volumes on reserved resources. The request
// is forwarded asynchronously to the Mesos agent where the reserved resources
// are located. That asynchronous message may not be delivered or creating the
// volumes at the agent might fail.
//
// References:
//
//  * http://mesos.apache.org/documentation/latest/operator-http-api/#create_volumes
func (m *Master) CreateVolumes(ctx context.Context, call *master.Call_CreateVolumes) (err error) {
	var callType master.Call_Type = master.Call_CREATE_VOLUMES
	var payload proto.Message = &master.Call{Type: &callType, CreateVolumes: call}
	var b []byte
	b, err = proto.Marshal(payload)
	if err != nil {
		return
	}
	var buf io.Reader = bytes.NewBuffer(b)
	_, err = m.client.doProtoWrapper(ctx, buf, nil)
	return
}

// This call destroys persistent volumes. The request is forwarded
// asynchronously to the Mesos agent where the reserved resources are located.
//
// References:
//
//  * http://mesos.apache.org/documentation/latest/operator-http-api/#destroy_volumes
func (m *Master) DestroyVolumes(ctx context.Context, call *master.Call_DestroyVolumes) (err error) {
	var callType master.Call_Type = master.Call_DESTROY_VOLUMES
	var payload proto.Message = &master.Call{Type: &callType, DestroyVolumes: call}
	var b []byte
	b, err = proto.Marshal(payload)
	if err != nil {
		return
	}
	var buf io.Reader = bytes.NewBuffer(b)
	_, err = m.client.doProtoWrapper(ctx, buf, nil)
	return
}

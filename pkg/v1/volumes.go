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
	"github.com/mesos/go-proto/mesos/v1/master"
)

// CreateVolumes creates persistent volumes on reserved resources. The request
// is forwarded asynchronously to the Mesos agent where the reserved resources
// are located. That asynchronous message may not be delivered or creating the
// volumes at the agent might fail.
func (m *Master) CreateVolumes(ctx context.Context, call *mesos_v1_master.Call_CreateVolumes) (err error) {
	var callType mesos_v1_master.Call_Type = mesos_v1_master.Call_CREATE_VOLUMES
	var message proto.Message = &mesos_v1_master.Call{Type: &callType, CreateVolumes: call}
	var httpResponse *http.Response
	httpResponse, err = m.client.makeCall(ctx, message, nil)
	defer httpResponse.Body.Close()
	return
}

// DestroyVolumes destroys persistent volumes. The request is forwarded
// asynchronously to the Mesos agent where the reserved resources are located.
func (m *Master) DestroyVolumes(ctx context.Context, call *mesos_v1_master.Call_DestroyVolumes) (err error) {
	var callType mesos_v1_master.Call_Type = mesos_v1_master.Call_DESTROY_VOLUMES
	var message proto.Message = &mesos_v1_master.Call{Type: &callType, DestroyVolumes: call}
	var httpResponse *http.Response
	httpResponse, err = m.client.makeCall(ctx, message, nil)
	defer httpResponse.Body.Close()
	return
}

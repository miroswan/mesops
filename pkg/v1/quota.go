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

// GetQuota retrieves the cluster’s configured quotas.
func (m *Master) GetQuota(ctx context.Context) (response *mesos_v1_master.Response, err error) {
	var httpResponse *http.Response
	response, httpResponse, err = m.sendSimpleCall(ctx, mesos_v1_master.Call_GET_QUOTA)
	defer httpResponse.Body.Close()
	return
}

// SetQuota sets the quota for resources to be used by a particular role.
func (m *Master) SetQuota(ctx context.Context, call *mesos_v1_master.Call_SetQuota) (err error) {
	var callType mesos_v1_master.Call_Type = mesos_v1_master.Call_SET_QUOTA
	var payload proto.Message = &mesos_v1_master.Call{Type: &callType, SetQuota: call}
	var b []byte
	b, err = proto.Marshal(payload)
	if err != nil {
		return
	}
	var buf io.Reader = bytes.NewBuffer(b)
	var httpResponse *http.Response
	httpResponse, err = m.client.doProtoWrapper(ctx, buf, nil)
	defer httpResponse.Body.Close()
	return
}

// RemoveQuota  removes the quota for a particular role.
func (m *Master) RemoveQuota(ctx context.Context, call *mesos_v1_master.Call_RemoveQuota) (err error) {
	var callType mesos_v1_master.Call_Type = mesos_v1_master.Call_REMOVE_QUOTA
	var payload proto.Message = &mesos_v1_master.Call{Type: &callType, RemoveQuota: call}
	var b []byte
	b, err = proto.Marshal(payload)
	if err != nil {
		return
	}
	var buf io.Reader = bytes.NewBuffer(b)
	var httpResponse *http.Response
	httpResponse, err = m.client.doProtoWrapper(ctx, buf, nil)
	defer httpResponse.Body.Close()
	return
}

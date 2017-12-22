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
	"encoding/json"
	"io"
)

type CreateVolumesPayload struct {
	Type          *string        `json:"type"`
	CreateVolumes *CreateVolumes `json:"create_volumes"`
}

type CreateVolumes struct {
	AgentID *struct {
		Value *string `json:"value"`
	} `json:"agent_id"`
	Volumes []*struct {
		Type *string `json:"type"`
		Disk *struct {
			Persistence *struct {
				ID        *string `json:"id"`
				Principal *string `json:"principal"`
			} `json:"persistence"`
			Volume *struct {
				ContainerPath *string `json:"container_path"`
				Mode          *string `json:"mode"`
			} `json:"volume"`
		} `json:"disk"`
		Name   *string `json:"name"`
		Role   *string `json:"role"`
		Scalar *struct {
			Value *float64 `json:"value"`
		} `json:"scalar"`
	} `json:"volumes"`
}

type DestroyVolumesPayload struct {
	Type           *string         `json:"type"`
	DestroyVolumes *DestroyVolumes `json:"destroy_volumes"`
}

type DestroyVolumes struct {
	AgentID *struct {
		Value *string `json:"value"`
	} `json:"agent_id"`
	Volumes []*struct {
		Disk *struct {
			Persistence *struct {
				ID        *string `json:"id"`
				Principal *string `json:"principal"`
			} `json:"persistence"`
			Volume *struct {
				ContainerPath *string `json:"container_path"`
				Mode          *string `json:"mode"`
			} `json:"volume"`
		} `json:"disk"`
		Name   *string `json:"name"`
		Role   *string `json:"role"`
		Scalar *struct {
			Value *float64 `json:"value"`
		} `json:"scalar"`
		Type *string `json:"type"`
	} `json:"volumes"`
}

// CreateVolumes
// http://mesos.apache.org/documentation/latest/operator-http-api/#create_volumes
func (m *Master) CreateVolumes(ctx context.Context, createVolumes *CreateVolumes) (err error) {
	var t string = "CREATE_VOLUMES"
	var cvp *CreateVolumesPayload = &CreateVolumesPayload{
		Type:          &t,
		CreateVolumes: createVolumes,
	}
	var b []byte
	b, err = json.Marshal(cvp)
	if err != nil {
		return
	}
	var buf io.Reader = bytes.NewBuffer(b)
	err = m.client.doWithRetryAndLoad(ctx, buf, nil)
	return
}

// DestroyVolumes
// http://mesos.apache.org/documentation/latest/operator-http-api/#destroy_volumes
func (m *Master) DestroyVolumes(ctx context.Context, destroyVolumes *DestroyVolumes) (err error) {
	var t string = "DESTROY_VOLUMES"
	var cvp *DestroyVolumesPayload = &DestroyVolumesPayload{
		Type:           &t,
		DestroyVolumes: destroyVolumes,
	}
	var b []byte
	b, err = json.Marshal(cvp)
	if err != nil {
		return
	}
	var buf io.Reader = bytes.NewBuffer(b)
	err = m.client.doWithRetryAndLoad(ctx, buf, nil)
	return
}

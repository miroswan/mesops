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

type ReserveResourcesPayload struct {
	Type             *string           `json:"type"`
	ReserveResources *ReserveResources `json:"reserve_resources"`
}

type ReserveResources struct {
	AgentID *struct {
		Value *string `json:"value"`
	} `json:"agent_id"`
	Resources []*struct {
		Type        *string `json:"type"`
		Name        *string `json:"name"`
		Reservation *struct {
			Principal *string `json:"principal"`
		} `json:"reservation"`
		Role   *string `json:"role"`
		Scalar *struct {
			Value *float64 `json:"value"`
		} `json:"scalar"`
	} `json:"resources"`
}

type UnreserveResourcesPayload struct {
	Type               *string             `json:"type"`
	UnreserveResources *UnreserveResources `json:"unreserve_resources"`
}

type UnreserveResources struct {
	AgentID *struct {
		Value *string `json:"value"`
	} `json:"agent_id"`
	Resources []*struct {
		Type        *string `json:"type"`
		Name        *string `json:"name"`
		Reservation *struct {
			Principal *string `json:"principal"`
		} `json:"reservation"`
		Role   *string `json:"role"`
		Scalar *struct {
			Value *float64 `json:"value"`
		} `json:"scalar"`
	} `json:"resources"`
}

// ReserveResources reserves resources dynamically on a specific agent.
//
// References:
//
// 	* http://mesos.apache.org/documentation/latest/operator-http-api/#reserve_resources
func (m *Master) ReserveResources(ctx context.Context, reserveResources *ReserveResources) (err error) {
	var t string = "RESERVE_RESOURCES"
	var rrp *ReserveResourcesPayload = &ReserveResourcesPayload{
		Type:             &t,
		ReserveResources: reserveResources,
	}
	var b []byte
	b, err = json.Marshal(rrp)
	if err != nil {
		return
	}
	var buf io.Reader = bytes.NewBuffer(b)
	err = m.client.doWithRetryAndLoad(ctx, buf, nil)
	return
}

// UnreserveResources unreserves resources dynamically on a specific agent.
//
// References:
//
// 	* http://mesos.apache.org/documentation/latest/operator-http-api/#unreserve_resources
func (m *Master) UnreserveResources(ctx context.Context, unreserveResources *UnreserveResources) (err error) {
	var t string = "UNRESERVE_RESOURCES"
	var urp *UnreserveResourcesPayload = &UnreserveResourcesPayload{
		Type:               &t,
		UnreserveResources: unreserveResources,
	}
	var b []byte
	b, err = json.Marshal(urp)
	if err != nil {
		return
	}
	var buf io.Reader = bytes.NewBuffer(b)
	err = m.client.doWithRetryAndLoad(ctx, buf, nil)
	return
}

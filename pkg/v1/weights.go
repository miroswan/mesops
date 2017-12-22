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

type GetWeightsResponse struct {
	Type       *string `json:"type"`
	GetWeights *struct {
		WeightInfos []*struct {
			Role   *string  `json:"role"`
			Weight *float64 `json:"weight"`
		} `json:"weight_infos"`
	} `json:"get_weights"`
}

type UpdateWeightsPayload struct {
	Type          *string        `json:"type"`
	UpdateWeights *UpdateWeights `json:"update_weights"`
}

type UpdateWeights struct {
	WeightInfos []*struct {
		Role   *string  `json:"role"`
		Weight *float64 `json:"weight"`
	} `json:"weight_infos"`
}

// GetWeights returns a *GetWeights.
// http://mesos.apache.org/documentation/latest/operator-http-api/#get_weights
func (m *Master) GetWeights(ctx context.Context) (gw *GetWeightsResponse, err error) {
	var buf io.Reader = simpleRequestPayload("GET_WEIGHTS")
	gw = &GetWeightsResponse{}
	err = m.client.doWithRetryAndLoad(ctx, buf, gw)
	return
}

// UpdateWeights
// http://mesos.apache.org/documentation/latest/operator-http-api/#update_weights
func (m *Master) UpdateWeights(ctx context.Context, updateWeights *UpdateWeights) (err error) {
	var t string = "UPDATE_WEIGHTS"
	var payload *UpdateWeightsPayload = &UpdateWeightsPayload{
		Type:          &t,
		UpdateWeights: updateWeights,
	}
	var b []byte
	b, err = json.Marshal(payload)
	if err != nil {
		return
	}
	var buf io.Reader = bytes.NewBuffer(b)
	err = m.client.doWithRetryAndLoad(ctx, buf, nil)
	return
}

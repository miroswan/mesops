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
	"io"
)

type GetRolesResponse struct {
	Type     *string `json:"type"`
	GetRoles *struct {
		Roles []*struct {
			Name       *string  `json:"name"`
			Weight     *float64 `json:"weight"`
			Frameworks []*struct {
				Value *string `json:"value"`
			} `json:"frameworks,omitempty"`
			Resources []*struct {
				Name   *string `json:"name"`
				Role   *string `json:"role"`
				Scalar *struct {
					Value *float64 `json:"value"`
				} `json:"scalar,omitempty"`
				Type   *string `json:"type"`
				Ranges *struct {
					Range []*struct {
						Begin *int `json:"begin"`
						End   *int `json:"end"`
					} `json:"range"`
				} `json:"ranges,omitempty"`
			} `json:"resources,omitempty"`
		} `json:"roles"`
	} `json:"get_roles"`
}

// GetRoles returns a pointer to a GetRoles.
//
// References:
//
// 	* http://mesos.apache.org/documentation/latest/operator-http-api/#get_roles
func (m *Master) GetRoles(ctx context.Context) (gr *GetRolesResponse, err error) {
	var buf io.Reader = simpleRequestPayload("GET_ROLES")
	gr = &GetRolesResponse{}
	err = m.client.doWithRetryAndLoad(ctx, buf, gr)
	return
}

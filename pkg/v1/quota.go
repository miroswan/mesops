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
	"fmt"
	"io"
)

type GetQuotaResponse struct {
	Type     *string `json:"type"`
	GetQuota *struct {
		Status *struct {
			Infos []*struct {
				Guarantee []*struct {
					Name   *string `json:"name"`
					Role   *string `json:"role"`
					Scalar *struct {
						Value *float64 `json:"value"`
					} `json:"scalar"`
					Type *string `json:"type"`
				} `json:"guarantee"`
				Principal *string `json:"principal"`
				Role      *string `json:"role"`
			} `json:"infos"`
		} `json:"status"`
	} `json:"get_quota"`
}

type SetQuotaPayload struct {
	Type     *string `json:"type"`
	SetQuota *SetQuota
}

type SetQuota struct {
	QuotaRequest *struct {
		Force     *bool `json:"force"`
		Guarantee []*struct {
			Name   *string `json:"name"`
			Role   *string `json:"role"`
			Scalar *struct {
				Value *float64 `json:"value"`
			} `json:"scalar"`
			Type *string `json:"type"`
		} `json:"guarantee"`
		Role *string `json:"role"`
	} `json:"quota_request"`
}

// GetQuota returns a pointer to a GetQuota.
//
// References:
//
// 	* http://mesos.apache.org/documentation/latest/operator-http-api/#get_quota
func (m *Master) GetQuota(ctx context.Context) (gq *GetQuotaResponse, err error) {
	var buf io.Reader = simpleRequestPayload("GET_QUOTA")
	gq = &GetQuotaResponse{}
	err = m.client.doWithRetryAndLoad(ctx, buf, gq)
	return
}

// SetQuota sets the quota for resources to be used by a particular role.
//
// References:
//
// 	* http://mesos.apache.org/documentation/latest/operator-http-api/#set_quota
func (m *Master) SetQuota(ctx context.Context, setQuota *SetQuota) (err error) {
	var t string = "SET_QUOTA"
	var payload *SetQuotaPayload = &SetQuotaPayload{
		Type:     &t,
		SetQuota: setQuota,
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

// RemoveQuota removes the quota for a particular role.
//
// References:
//
// 	* http://mesos.apache.org/documentation/latest/operator-http-api/#remove_quota
func (m *Master) RemoveQuota(ctx context.Context, role string) (err error) {
	const tmpl string = `
  {
    "type": "REMOVE_QUOTA",
    "remove_quota": {
      "role": "%s"
    }
  }
  `
	var txt string = fmt.Sprintf(tmpl, role)
	var buf io.Reader = bytes.NewBuffer([]byte(txt))
	err = m.client.doWithRetryAndLoad(ctx, buf, nil)
	return
}

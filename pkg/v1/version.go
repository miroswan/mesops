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

type GetVersionResponse struct {
	Type       *string `json:"type"`
	GetVersion *struct {
		VersionInfo *struct {
			Version   *string  `json:"version"`
			BuildDate *string  `json:"build_date"`
			BuildTime *float64 `json:"build_time"`
			BuildUser *string  `json:"build_user"`
		} `json:"version_info"`
	} `json:"get_version"`
}

// GetVersion returns a *GetVersion.
// http://mesos.apache.org/documentation/latest/operator-http-api/#get_version
func (m *Master) GetVersion(ctx context.Context) (gv *GetVersionResponse, err error) {
	gv, err = getVersion(ctx, m.client)
	return
}

// GetVersion returns a *GetVersion.
// http://mesos.apache.org/documentation/latest/operator-http-api/#get_version-1
func (a *Agent) GetVersion(ctx context.Context) (gv *GetVersionResponse, err error) {
	gv, err = getVersion(ctx, a.client)
	return
}

func getVersion(ctx context.Context, client *client) (gv *GetVersionResponse, err error) {
	var buf io.Reader = simpleRequestPayload("GET_VERSION")
	gv = &GetVersionResponse{}
	err = client.doWithRetryAndLoad(ctx, buf, gv)
	return
}

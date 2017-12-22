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
	"fmt"
	"io"
	"strings"
)

type GetLoggingLevelResponse struct {
	Type            *string `json:"type"`
	GetLoggingLevel *struct {
		Level *int `json:"level"`
	} `json:"get_logging_level"`
}

// GetLoggingLevel returns a pointer to a GetLoggingLevel.
//
// References:
//
// 	* http://mesos.apache.org/documentation/latest/operator-http-api/#get_logging_level
func (m *Master) GetLoggingLevel(ctx context.Context) (gll *GetLoggingLevelResponse, err error) {
	gll, err = getLoggingLevel(ctx, m.client)
	return
}

// GetLoggingLevel returns a pointer to a GetLoggingLevel.
//
// References:
//
// 	* http://mesos.apache.org/documentation/latest/operator-http-api/#get_logging_level-1
func (a *Agent) GetLoggingLevel(ctx context.Context) (gll *GetLoggingLevelResponse, err error) {
	gll, err = getLoggingLevel(ctx, a.client)
	return
}

// SetLoggingLevel sets the logging level on the configured Master
//
// References:
// * http://mesos.apache.org/documentation/latest/operator-http-api/#set_logging_level
func (m *Master) SetLoggingLevel(ctx context.Context, ll int) (err error) {
	err = setLoggingLevel(ctx, m.client, ll)
	return
}

// SetLoggingLevel sets the logging level on the configured Agent
//
// References:
// * http://mesos.apache.org/documentation/latest/operator-http-api/#set_logging_level-1
func (a *Agent) SetLoggingLevel(ctx context.Context, ll int) (err error) {
	err = setLoggingLevel(ctx, a.client, ll)
	return
}

func getLoggingLevel(ctx context.Context, client *client) (gll *GetLoggingLevelResponse, err error) {
	var buf io.Reader = simpleRequestPayload("GET_LOGGING_LEVEL")
	gll = &GetLoggingLevelResponse{}
	err = client.doWithRetryAndLoad(ctx, buf, gll)
	return
}

func setLoggingLevel(ctx context.Context, client *client, ll int) (err error) {
	var j string = `
  {
    "type": "SET_LOGGING_LEVEL",
    "set_logging_level": {
      "duration": {
        "nanoseconds": 60000000000
      },
      "level": %d
    }
  }
  `
	j = strings.TrimSpace(fmt.Sprintf(j, ll))
	var buf io.Reader = bytes.NewBuffer([]byte(j))
	err = client.doWithRetryAndLoad(ctx, buf, nil)
	return
}

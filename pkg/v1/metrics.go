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

	"github.com/miroswan/mesops/pkg/v1/agent"
	"github.com/miroswan/mesops/pkg/v1/master"
)

// GetMetrics gives the snapshot of current metrics to the end user. If timeout
// is set in the call, it will be used to determine the maximum amount of time
// the API will take to respond. If the timeout is exceeded, some metrics may
// not be included in the response.
func (m *Master) GetMetrics(ctx context.Context) (response *master.Response, err error) {
	response, _, err = m.sendSimpleCall(ctx, master.Call_GET_METRICS)
	return
}

// GetMetrics gives the snapshot of current metrics to the end user. If timeout
// is set in the call, it will be used to determine the maximum amount of time
// the API will take to respond. If the timeout is exceeded, some metrics may
// not be included in the response.
func (a *Agent) GetMetrics(ctx context.Context) (response *agent.Response, err error) {
	response, _, err = a.sendSimpleCall(ctx, agent.Call_GET_METRICS)
	return
}

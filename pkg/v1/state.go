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

	"github.com/miroswan/mesops/pkg/v1/agent"
	"github.com/miroswan/mesops/pkg/v1/master"
)

// GetState returns a pointer to a master.GetState
//
// References:
//
// 	* http://mesos.apache.org/documentation/latest/operator-http-api/#get_state
func (m *Master) GetState(ctx context.Context) (gs *master.GetStateResponse, err error) {
	var buf io.Reader = simpleRequestPayload("GET_STATE")
	gs = &master.GetStateResponse{}
	err = m.client.doWithRetryAndLoad(ctx, buf, gs)
	return
}

// GetState returns a pointer to an agent.GetState
//
// References:
//
// 	* http://mesos.apache.org/documentation/latest/operator-http-api/#get_state-1
func (a *Agent) GetState(ctx context.Context) (gs *agent.GetStateResponse, err error) {
	var buf io.Reader = simpleRequestPayload("GET_STATE")
	gs = &agent.GetStateResponse{}
	err = a.client.doWithRetryAndLoad(ctx, buf, gs)
	return
}

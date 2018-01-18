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
	"net/http"

	"github.com/mesos/go-proto/mesos/v1/agent"
	"github.com/mesos/go-proto/mesos/v1/master"
)

// GetTasks queries about all the tasks known to the mesos_v1_master.
func (m *Master) GetTasks(ctx context.Context) (response *mesos_v1_master.Response, err error) {
	var httpResponse *http.Response
	response, httpResponse, err = m.sendSimpleCall(ctx, mesos_v1_master.Call_GET_TASKS)
	defer httpResponse.Body.Close()
	return
}

// GetTasks queries about all the tasks known to the agent.
func (a *Agent) GetTasks(ctx context.Context) (response *mesos_v1_agent.Response, err error) {
	var httpResponse *http.Response
	response, httpResponse, err = a.sendSimpleCall(ctx, mesos_v1_agent.Call_GET_TASKS)
	defer httpResponse.Body.Close()
	return
}

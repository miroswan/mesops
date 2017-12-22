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
)

type GetAgentsResponse struct {
	Type      *string `json:"type"`
	GetAgents *struct {
		Agents []*struct {
			Active    *bool `json:"active"`
			AgentInfo *struct {
				Hostname *string `json:"hostname"`
				ID       *struct {
					Value *string `json:"value"`
				} `json:"id"`
				Port      *int `json:"port"`
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
				} `json:"resources"`
			} `json:"agent_info"`
			Pid            *string `json:"pid"`
			RegisteredTime *struct {
				Nanoseconds *int64 `json:"nanoseconds"`
			} `json:"registered_time"`
			TotalResources []*struct {
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
			} `json:"total_resources"`
			Version *string `json:"version"`
		} `json:"agents"`
	} `json:"get_agents"`
}

// GetAgents returns a pointer to a GetAgents.
//
// References:
//
// 	* http://mesos.apache.org/documentation/latest/operator-http-api/#get_agents
func (m *Master) GetAgents(ctx context.Context) (ga *GetAgentsResponse, err error) {
	var buf io.Reader = simpleRequestPayload("GET_AGENTS")
	ga = &GetAgentsResponse{}
	err = m.client.doWithRetryAndLoad(ctx, buf, ga)
	return
}

// MarkAgentGone can be used by operators to assert that an agent instance has
// failed and is never coming back
//
// References:
//
// 	* http://mesos.apache.org/documentation/latest/operator-http-api/#mark_agent_gone
func (m *Master) MarkAgentGone(ctx context.Context, agentUUID string) (err error) {
	const tmpl string = `
	{
	  "type": "MARK_AGENT_GONE",
	  "mark_agent_gone": {
	    "agent_id": {
	      "value": "%s"
	    }
	  }
	}
	`
	var txt string = fmt.Sprintf(tmpl, agentUUID)
	var buf io.Reader = bytes.NewBuffer([]byte(txt))
	err = m.client.doWithRetryAndLoad(ctx, buf, nil)
	return
}

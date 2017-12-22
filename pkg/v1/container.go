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

type GetContainersResponse struct {
	Type          *string `json:"type"`
	GetContainers *struct {
		Containers []*struct {
			ContainerID *struct {
				Value *string `json:"value"`
			} `json:"container_id"`
			ContainerStatus *struct {
				NetworkInfos []*struct {
					IPAddresses []struct {
						IPAddress *string `json:"ip_address"`
					} `json:"ip_addresses"`
				} `json:"network_infos"`
			} `json:"container_status"`
			ExecutorID *struct {
				Value *string `json:"value"`
			} `json:"executor_id"`
			ExecutorName *string `json:"executor_name"`
			FrameworkID  *struct {
				Value *string `json:"value"`
			} `json:"framework_id"`
			ResourceStatistics *struct {
				MemLimitBytes *int     `json:"mem_limit_bytes"`
				Timestamp     *float64 `json:"timestamp"`
			} `json:"resource_statistics"`
		} `json:"containers"`
	} `json:"get_containers"`
}

type LaunchNestedContainerPayload struct {
	Type                  *string                `json:"type"`
	LaunchNestedContainer *LaunchNestedContainer `json:"launch_nested_container"`
}

type LaunchNestedContainer struct {
	ContainerID *struct {
		Parent *struct {
			Parent *struct {
				Value *string `json:"value"`
			} `json:"parent"`
			Value *string `json:"value"`
		} `json:"parent"`
		Value *string `json:"value"`
	} `json:"container_id"`
	Command *struct {
		Environment *struct {
			Variables []*struct {
				Name  *string `json:"name"`
				Type  *string `json:"type"`
				Value *string `json:"value"`
			} `json:"variables"`
		} `json:"environment"`
		Shell *bool   `json:"shell"`
		Value *string `json:"value"`
	} `json:"command"`
}

type WaitNestedContainerPayload struct {
	Type                *string              `json:"type"`
	WaitNestedContainer *WaitNestedContainer `json:"wait_nested_container"`
}

type WaitNestedContainer struct {
	ContainerID *struct {
		Parent *struct {
			Value *string `json:"value"`
		} `json:"parent"`
		Value *string `json:"value"`
	} `json:"container_id"`
}

type WaitNestedContainerResponse struct {
	Type                *string `json:"type"`
	WaitNestedContainer *struct {
		ExitStatus *int `json:"exit_status"`
	} `json:"wait_nested_container"`
}

type KillNestedContainerPayload struct {
	Type                *string              `json:"type"`
	KillNestedContainer *KillNestedContainer `json:"kill_nested_container"`
}

type KillNestedContainer struct {
	ContainerID *struct {
		Parent *struct {
			Value *string `json:"value"`
		} `json:"parent"`
		Value *string `json:"value"`
	} `json:"container_id"`
}

// GetContainers returns a pointer to a GetContainers
//
// References:
//
// 	* http://mesos.apache.org/documentation/latest/operator-http-api/#get_containers
func (a *Agent) GetContainers(ctx context.Context, showNested bool, showStandalone bool) (gc *GetContainersResponse, err error) {
	var tmpl string = `
  {
    "type": "GET_CONTAINERS",
    "get_containers": {
      "show_nested": %b,
      "show_standalone": %b
    }
  }
  `
	var txt string = fmt.Sprintf(tmpl, showNested, showStandalone)
	var buf io.Reader = bytes.NewBuffer([]byte(txt))
	gc = &GetContainersResponse{}
	err = a.client.doWithRetryAndLoad(ctx, buf, gc)
	return
}

// LaunchNestedContainer launches a nested container on the configured Agent
//
// References:
//
// 	* http://mesos.apache.org/documentation/latest/operator-http-api/#launch_nested_container
func (a *Agent) LaunchNestedContainer(ctx context.Context, l *LaunchNestedContainer) (err error) {
	var t string = "LAUNCH_NESTED_CONTAINER"
	var payload *LaunchNestedContainerPayload = &LaunchNestedContainerPayload{
		Type: &t,
		LaunchNestedContainer: l,
	}
	var b []byte
	b, err = json.Marshal(payload)
	if err != nil {
		return
	}
	var buf io.Reader = bytes.NewBuffer(b)
	err = a.client.doWithRetryAndLoad(ctx, buf, nil)
	return
}

// WaitNestedContainer waits for a nested container to terminate or exit. Any
// authorized entity, including the executor itself, its tasks, or the operator
// can use this API to wait on a nested container.
//
// References:
//
// 	* http://mesos.apache.org/documentation/latest/operator-http-api/#wait_nested_container
func (a *Agent) WaitNestedContainer(ctx context.Context, w *WaitNestedContainer) (err error, wnc *WaitNestedContainerResponse) {
	var t string = "WAIT_NESTED_CONTAINER"
	var payload *WaitNestedContainerPayload = &WaitNestedContainerPayload{
		Type:                &t,
		WaitNestedContainer: w,
	}
	var b []byte
	b, err = json.Marshal(payload)
	if err != nil {
		return
	}
	var buf io.Reader = bytes.NewBuffer(b)
	wnc = &WaitNestedContainerResponse{}
	err = a.client.doWithRetryAndLoad(ctx, buf, wnc)
	return
}

// KillNestedContainer initiates the destruction of a nested container. Any
// authorized entity, including the executor itself, its tasks, or the operator
// can use this API to kill a nested container.
//
// References:
//
// 	* http://mesos.apache.org/documentation/latest/operator-http-api/#kill_nested_container
func (a *Agent) KillNestedContainer(ctx context.Context, k *KillNestedContainer) (err error) {
	var t string = "KILL_NESTED_CONTAINER"
	var payload *KillNestedContainerPayload = &KillNestedContainerPayload{
		Type:                &t,
		KillNestedContainer: k,
	}
	var b []byte
	b, err = json.Marshal(payload)
	if err != nil {
		return
	}
	var buf io.Reader = bytes.NewBuffer(b)
	err = a.client.doWithRetryAndLoad(ctx, buf, nil)
	return
}

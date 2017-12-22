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

package master

type GetExecutorsResponse struct {
	Type         *string `json:"type"`
	GetExecutors *struct {
		Executors []*struct {
			AgentID *struct {
				Value *string `json:"value"`
			} `json:"agent_id"`
			ExecutorInfo *struct {
				Command *struct {
					Shell *bool   `json:"shell"`
					Value *string `json:"value"`
				} `json:"command"`
				ExecutorID *struct {
					Value *string `json:"value"`
				} `json:"executor_id"`
				FrameworkID *struct {
					Value *string `json:"value"`
				} `json:"framework_id"`
			} `json:"executor_info"`
		} `json:"executors"`
	} `json:"get_executors"`
}

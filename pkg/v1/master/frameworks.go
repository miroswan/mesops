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

type GetFrameworksResponse struct {
	Type          *string `json:"type"`
	GetFrameworks *struct {
		Frameworks []*struct {
			Active        *bool `json:"active"`
			Connected     *bool `json:"connected"`
			FrameworkInfo *struct {
				Checkpoint      *bool    `json:"checkpoint"`
				FailoverTimeout *float64 `json:"failover_timeout"`
				Hostname        *string  `json:"hostname"`
				ID              *struct {
					Value *string `json:"value"`
				} `json:"id"`
				Name      *string `json:"name"`
				Principal *string `json:"principal"`
				Role      *string `json:"role"`
				User      *string `json:"user"`
			} `json:"framework_info"`
			RegisteredTime *struct {
				Nanoseconds *int64 `json:"nanoseconds"`
			} `json:"registered_time"`
			ReregisteredTime *struct {
				Nanoseconds *int64 `json:"nanoseconds"`
			} `json:"reregistered_time"`
		} `json:"frameworks"`
	} `json:"get_frameworks"`
}

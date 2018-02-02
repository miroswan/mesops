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
	"bufio"
	"context"
	"net/http"

	"github.com/gogo/protobuf/proto"
	"github.com/mesos/go-proto/mesos/v1/master"
)

type EventStream chan *mesos_v1_master.Event

// Subscribe subscribes to events on the Mesos mesos_v1_master. This method blocks, so
// you. likely want to call it in a go routine. Process each *mesos_v1_master.Event on
// the EventStream by checking the type (you may call GetType() on the
// *mesos_v1_master.Event), then processing the data as needed. See the test/cmd
// package for an example.
func (m *Master) Subscribe(ctx context.Context, es EventStream) (err error) {
	var httpResponse *http.Response
	var callType mesos_v1_master.Call_Type = mesos_v1_master.Call_SUBSCRIBE
	var callMsg proto.Message = &mesos_v1_master.Call{Type: &callType}

	httpResponse, err = m.client.makeCall(ctx, callMsg, nil)
	if err != nil {
		return
	}
	var reader *bufio.Reader = bufio.NewReader(httpResponse.Body)
	defer httpResponse.Body.Close()
	for {
		select {
		case <-ctx.Done():
			err = ctx.Err()
			return
		default:
			var msg []byte
			msg, err = readRecordioMessage(reader)
			if err != nil {
				return
			}

			// Unmarshal data into a mesos_v1_master.Event
			event := &mesos_v1_master.Event{}
			err = proto.Unmarshal(msg, event)
			if err != nil {
				return
			}
			es <- event
		}
	}
}

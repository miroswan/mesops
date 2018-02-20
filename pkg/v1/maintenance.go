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

	"github.com/gogo/protobuf/proto"
	"github.com/mesos/go-proto/mesos/v1/master"
)

// GetMaintenanceStatus retrieves the cluster’s maintenance status.
func (m *Master) GetMaintenanceStatus(ctx context.Context) (response *mesos_v1_master.Response, err error) {
	var httpResponse *http.Response
	response, httpResponse, err = m.sendSimpleCall(ctx, mesos_v1_master.Call_GET_MAINTENANCE_STATUS)
	defer httpResponse.Body.Close()
	return
}

// GetMaintenanceSchedule retrieves the cluster’s maintenance schedule.
func (m *Master) GetMaintenanceSchedule(ctx context.Context) (response *mesos_v1_master.Response, err error) {
	var httpResponse *http.Response
	response, httpResponse, err = m.sendSimpleCall(ctx, mesos_v1_master.Call_GET_MAINTENANCE_SCHEDULE)
	defer httpResponse.Body.Close()
	return
}

// UpdateMaintenanceSchedule updates the cluster’s maintenance schedule.
func (m *Master) UpdateMaintenanceSchedule(ctx context.Context, call *mesos_v1_master.Call_UpdateMaintenanceSchedule) (err error) {
	var callType mesos_v1_master.Call_Type = mesos_v1_master.Call_UPDATE_MAINTENANCE_SCHEDULE
	var message proto.Message = &mesos_v1_master.Call{
		Type: &callType,
		UpdateMaintenanceSchedule: call,
	}
	var httpResponse *http.Response
	httpResponse, err = m.client.makeCall(ctx, message, nil)
	defer httpResponse.Body.Close()
	return
}

// StartMaintenance starts the maintenance of the cluster, this would bring a
// set of machines down.
func (m *Master) StartMaintenance(ctx context.Context, call *mesos_v1_master.Call_StartMaintenance) (err error) {
	var callType mesos_v1_master.Call_Type = mesos_v1_master.Call_START_MAINTENANCE
	var message proto.Message = &mesos_v1_master.Call{
		Type:             &callType,
		StartMaintenance: call,
	}
	var httpResponse *http.Response
	httpResponse, err = m.client.makeCall(ctx, message, nil)
	defer httpResponse.Body.Close()
	return
}

// StopMaintenace stops the maintenance of the cluster, this would bring a set of
// machines back up.
func (m *Master) StopMaintenance(ctx context.Context, call *mesos_v1_master.Call_StopMaintenance) (err error) {
	var callType mesos_v1_master.Call_Type = mesos_v1_master.Call_STOP_MAINTENANCE
	var message proto.Message = &mesos_v1_master.Call{
		Type:            &callType,
		StopMaintenance: call,
	}
	var httpResponse *http.Response
	httpResponse, err = m.client.makeCall(ctx, message, nil)
	defer httpResponse.Body.Close()
	return
}

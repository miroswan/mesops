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
	"io"
)

type GetMaintenanceStatusResponse struct {
	Type                 *string `json:"type"`
	GetMaintenanceStatus *struct {
		Status *struct {
			DrainingMachines []*struct {
				ID *struct {
					IP *string `json:"ip"`
				} `json:"id"`
			} `json:"draining_machines"`
		} `json:"status"`
	} `json:"get_maintenance_status"`
}

type GetMaintenanceScheduleResponse struct {
	Type                   *string `json:"type"`
	GetMaintenanceSchedule *struct {
		Schedule *struct {
			Windows []*struct {
				MachineIDs []*struct {
					Hostname *string `json:"hostname,omitempty"`
					IP       *string `json:"ip,omitempty"`
				} `json:"machine_ids"`
				Unavailability *struct {
					Start *struct {
						Nanoseconds *int64 `json:"nanoseconds"`
					} `json:"start"`
				} `json:"unavailability"`
			} `json:"windows"`
		} `json:"schedule"`
	} `json:"get_maintenance_schedule"`
}

type StartMaintenancePayload struct {
	Type             *string           `json:"type"`
	StartMaintenance *StartMaintenance `json:"start_maintenance"`
}

// StartMaintenance for StartMaintenancePayload
type StartMaintenance struct {
	Machines []*struct {
		Hostname *string `json:"hostname"`
		IP       *string `json:"ip"`
	} `json:"machines"`
}

type UpdateMaintenanceSchedulePayload struct {
	Type                      *string                    `json:"type"`
	UpdateMaintenanceSchedule *UpdateMaintenanceSchedule `json:"update_maintenance_schedule"`
}

type UpdateMaintenanceSchedule struct {
	Schedule *struct {
		Windows []*struct {
			MachineIds []*struct {
				Hostname *string `json:"hostname,omitempty"`
				IP       *string `json:"ip,omitempty"`
			} `json:"machine_ids"`
			Unavailability *struct {
				Start *struct {
					Nanoseconds *int64 `json:"nanoseconds"`
				} `json:"start"`
			} `json:"unavailability"`
		} `json:"windows"`
	}
}

// StopMaintenancePayload is the data strucrture encoded to JSON within a
// StopMaintenance call on a Master
type StopMaintenancePayload struct {
	Type            *string          `json:"type"`
	StopMaintenance *StopMaintenance `json:"stop_maintenance"`
}

// StopMaintenance for StopMaintenancePayload
type StopMaintenance struct {
	Machines []*struct {
		Hostname *string `json:"hostname"`
		IP       *string `json:"ip"`
	} `json:"machines"`
}

// GetMaintenanceSchedule returns a pointer to a GetMaintenanceSchedule.
//
// References:
//
// 	* http://mesos.apache.org/documentation/latest/operator-http-api/#get_maintenance_schedule
func (m *Master) GetMaintenanceSchedule(ctx context.Context) (gs *GetMaintenanceScheduleResponse, err error) {
	var buf io.Reader = simpleRequestPayload("GET_MAINTENANCE_SCHEDULE")
	gs = &GetMaintenanceScheduleResponse{}
	err = m.client.doWithRetryAndLoad(ctx, buf, gs)
	return
}

// GetMaintenanceStatus returns a pointer to a GetMaintenanceStatus
//
// References:
//
// 	* http://mesos.apache.org/documentation/latest/operator-http-api/#get_maintenance_status
func (m *Master) GetMaintenanceStatus(ctx context.Context) (gs *GetMaintenanceStatusResponse, err error) {
	var buf io.Reader = simpleRequestPayload("GET_MAINTENANCE_STATUS")
	gs = &GetMaintenanceStatusResponse{}
	err = m.client.doWithRetryAndLoad(ctx, buf, gs)
	return
}

// UpdateMaintenanceSchedule updates the cluster’s maintenance schedule.
//
// References:
//
// 	* http://mesos.apache.org/documentation/latest/operator-http-api/#update_maintenance_schedule
func (m *Master) UpdateMaintenanceSchedule(ctx context.Context, updateMaintenanceSchedule *UpdateMaintenanceSchedule) (err error) {
	var t string = "UPDATE_MAINTENANCE_SCHEDULE"
	var ums *UpdateMaintenanceSchedulePayload = &UpdateMaintenanceSchedulePayload{
		Type: &t,
		UpdateMaintenanceSchedule: updateMaintenanceSchedule,
	}
	var b []byte
	b, err = json.Marshal(ums)
	if err != nil {
		return
	}
	var buf io.Reader = bytes.NewBuffer(b)
	err = m.client.doWithRetryAndLoad(ctx, buf, nil)
	return
}

// StartMaintenance starts the maintenance of the cluster, this would bring a
// set of machines down.
//
// References:
//
// 	* http://mesos.apache.org/documentation/latest/operator-http-api/#start_maintenance
func (m *Master) StartMaintenance(ctx context.Context, startMaintenance *StartMaintenance) (err error) {
	var t string = "START_MAINTENANCE"
	var payload *StartMaintenancePayload = &StartMaintenancePayload{
		Type:             &t,
		StartMaintenance: startMaintenance,
	}
	var b []byte
	b, err = json.Marshal(payload)
	if err != nil {
		return
	}
	var buf io.Reader = bytes.NewBuffer(b)
	err = m.client.doWithRetryAndLoad(ctx, buf, nil)
	return
}

// StopMaintenance stops the maintenance of the cluster, this would bring a set
// of machines back up.
//
// References:
//
// 	* http://mesos.apache.org/documentation/latest/operator-http-api/#stop_maintenance
func (m *Master) StopMaintenance(ctx context.Context, stopMaintenance *StopMaintenance) (err error) {
	var t string = "STOP_MAINTENANCE"
	var payload *StopMaintenancePayload = &StopMaintenancePayload{
		Type:            &t,
		StopMaintenance: stopMaintenance,
	}
	var b []byte
	b, err = json.Marshal(payload)
	if err != nil {
		return
	}
	var buf io.Reader = bytes.NewBuffer(b)
	err = m.client.doWithRetryAndLoad(ctx, buf, nil)
	return
}

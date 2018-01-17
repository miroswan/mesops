package v1

import (
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/mesos/go-proto/mesos/v1"
	"github.com/mesos/go-proto/mesos/v1/allocator"
	"github.com/mesos/go-proto/mesos/v1/maintenance"
	"github.com/mesos/go-proto/mesos/v1/master"
)

func TestGetMaintenanceStatus(t *testing.T) {
	s := NewTestProtobufServer(MasterClient)
	defer s.Teardown()

	responseType := mesos_v1_master.Response_GET_MAINTENANCE_STATUS
	hostname := "test-node"
	ip := "127.0.0.1"
	inverseOfferStatus := mesos_v1_allocator.InverseOfferStatus_Status(1)
	frameworkID := "test-framework"
	nanoseconds := int64(2000000)
	response := mesos_v1_master.Response{
		Type: &responseType,
		GetMaintenanceStatus: &mesos_v1_master.Response_GetMaintenanceStatus{
			Status: &mesos_v1_maintenance.ClusterStatus{
				DrainingMachines: []*mesos_v1_maintenance.ClusterStatus_DrainingMachine{
					&mesos_v1_maintenance.ClusterStatus_DrainingMachine{
						Id: &mesos_v1.MachineID{
							Hostname: &hostname,
							Ip:       &ip,
						},
						Statuses: []*mesos_v1_allocator.InverseOfferStatus{
							&mesos_v1_allocator.InverseOfferStatus{
								Status: &inverseOfferStatus,
								FrameworkId: &mesos_v1.FrameworkID{
									Value: &frameworkID,
								},
								Timestamp: &mesos_v1.TimeInfo{
									Nanoseconds: &nanoseconds,
								},
							},
						},
					},
				},
			},
		},
	}

	b, err := proto.Marshal(&response)
	if err != nil {
		t.Fatal(err)
	}

	s.SetOutput(b).Handle()

	data, err := s.Master().GetMaintenanceStatus(s.Ctx())
	if err != nil {
		t.Error(err)
	}

	respHostname := data.GetGetMaintenanceStatus().
		GetStatus().
		GetDrainingMachines()[0].
		GetId().
		GetHostname()

	respIp := data.GetGetMaintenanceStatus().
		GetStatus().
		GetDrainingMachines()[0].
		GetId().
		GetIp()

	respFrameworkID := data.GetGetMaintenanceStatus().
		GetStatus().
		GetDrainingMachines()[0].
		GetStatuses()[0].
		GetFrameworkId().
		GetValue()

	if hostname != respHostname {
		t.Errorf("expected %s, got %s", hostname, respHostname)
	}

	if ip != respIp {
		t.Errorf("expected %s, got %s", ip, respIp)
	}

	if frameworkID != respFrameworkID {
		t.Errorf("expected %s, got %s", frameworkID, respFrameworkID)
	}
}

func TestGetMaintenanceSchedule(t *testing.T) {
	s := NewTestProtobufServer(MasterClient)
	defer s.Teardown()

	responseType := mesos_v1_master.Response_GET_MAINTENANCE_SCHEDULE
	hostname := "test-node"
	ip := "127.0.0.1"
	nanoseconds := int64(2000000)
	response := mesos_v1_master.Response{
		Type: &responseType,
		GetMaintenanceSchedule: &mesos_v1_master.Response_GetMaintenanceSchedule{
			Schedule: &mesos_v1_maintenance.Schedule{
				Windows: []*mesos_v1_maintenance.Window{
					&mesos_v1_maintenance.Window{
						MachineIds: []*mesos_v1.MachineID{
							&mesos_v1.MachineID{
								Hostname: &hostname,
								Ip:       &ip,
							},
						},
						Unavailability: &mesos_v1.Unavailability{
							Start: &mesos_v1.TimeInfo{
								Nanoseconds: &nanoseconds,
							},
						},
					},
				},
			},
		},
	}

	b, err := proto.Marshal(&response)
	if err != nil {
		t.Fatal(err)
	}
	s.SetOutput(b).Handle()

	data, err := s.Master().GetMaintenanceSchedule(s.Ctx())
	if err != nil {
		t.Fatal(err)
	}
	resHostname := data.GetGetMaintenanceSchedule().
		GetSchedule().
		GetWindows()[0].
		GetMachineIds()[0].
		GetHostname()

	resIP := data.GetGetMaintenanceSchedule().
		GetSchedule().
		GetWindows()[0].
		GetMachineIds()[0].
		GetIp()

	if hostname != resHostname {
		t.Errorf("expected %s, got %s", hostname, resHostname)
	}

	if ip != resIP {
		t.Errorf("expected %s, got %s", ip, resIP)
	}
}

func TestUpdateMaintenanceSchedule(t *testing.T) {
	s := NewTestProtobufServer(MasterClient)
	defer s.Teardown()

	s.Handle()
	hostname := "test-hostname"
	ip := "127.0.0.1"
	nanoseconds := int64(2000000)

	call := &mesos_v1_master.Call_UpdateMaintenanceSchedule{
		Schedule: &mesos_v1_maintenance.Schedule{
			Windows: []*mesos_v1_maintenance.Window{
				&mesos_v1_maintenance.Window{
					MachineIds: []*mesos_v1.MachineID{
						&mesos_v1.MachineID{
							Hostname: &hostname,
							Ip:       &ip,
						},
					},
					Unavailability: &mesos_v1.Unavailability{
						Start: &mesos_v1.TimeInfo{
							Nanoseconds: &nanoseconds,
						},
					},
				},
			},
		},
	}

	err := s.Master().UpdateMaintenanceSchedule(s.Ctx(), call)
	if err != nil {
		t.Error(err)
	}
}

func TestStartMaintenace(t *testing.T) {
	s := NewTestProtobufServer(MasterClient)
	defer s.Teardown()

	s.Handle()

	hostname := "test-node"
	ip := "127.0.0.1"
	call := &mesos_v1_master.Call_StartMaintenance{
		Machines: []*mesos_v1.MachineID{
			&mesos_v1.MachineID{
				Hostname: &hostname,
				Ip:       &ip,
			},
		},
	}

	err := s.Master().StartMaintenance(s.Ctx(), call)
	if err != nil {
		t.Error(err)
	}
}

func TestStopMaintenace(t *testing.T) {
	s := NewTestProtobufServer(MasterClient)
	defer s.Teardown()

	s.Handle()

	hostname := "test-node"
	ip := "127.0.0.1"
	call := &mesos_v1_master.Call_StopMaintenance{
		Machines: []*mesos_v1.MachineID{
			&mesos_v1.MachineID{
				Hostname: &hostname,
				Ip:       &ip,
			},
		},
	}

	err := s.Master().StopMaintenance(s.Ctx(), call)
	if err != nil {
		t.Error(err)
	}
}

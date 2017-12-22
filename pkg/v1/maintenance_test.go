package v1

import (
	"testing"

	"github.com/gogo/protobuf/proto"

	"github.com/miroswan/mesops/pkg/v1/allocator"
	"github.com/miroswan/mesops/pkg/v1/maintenance"
	"github.com/miroswan/mesops/pkg/v1/master"
	"github.com/miroswan/mesops/pkg/v1/mesos"
)

func TestGetMaintenanceStatus(t *testing.T) {
	s := NewTestProtobufServer(MasterClient)
	defer s.Teardown()

	responseType := master.Response_GET_MAINTENANCE_STATUS
	hostname := "test-node"
	ip := "127.0.0.1"
	inverseOfferStatus := allocator.InverseOfferStatus_Status(1)
	frameworkID := "test-framework"
	nanoseconds := int64(2000000)
	response := master.Response{
		Type: &responseType,
		GetMaintenanceStatus: &master.Response_GetMaintenanceStatus{
			Status: &maintenance.ClusterStatus{
				DrainingMachines: []*maintenance.ClusterStatus_DrainingMachine{
					&maintenance.ClusterStatus_DrainingMachine{
						Id: &mesos.MachineID{
							Hostname: &hostname,
							Ip:       &ip,
						},
						Statuses: []*allocator.InverseOfferStatus{
							&allocator.InverseOfferStatus{
								Status: &inverseOfferStatus,
								FrameworkId: &mesos.FrameworkID{
									Value: &frameworkID,
								},
								Timestamp: &mesos.TimeInfo{
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

	responseType := master.Response_GET_MAINTENANCE_SCHEDULE
	hostname := "test-node"
	ip := "127.0.0.1"
	nanoseconds := int64(2000000)
	response := master.Response{
		Type: &responseType,
		GetMaintenanceSchedule: &master.Response_GetMaintenanceSchedule{
			Schedule: &maintenance.Schedule{
				Windows: []*maintenance.Window{
					&maintenance.Window{
						MachineIds: []*mesos.MachineID{
							&mesos.MachineID{
								Hostname: &hostname,
								Ip:       &ip,
							},
						},
						Unavailability: &mesos.Unavailability{
							Start: &mesos.TimeInfo{
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

	call := &master.Call_UpdateMaintenanceSchedule{
		Schedule: &maintenance.Schedule{
			Windows: []*maintenance.Window{
				&maintenance.Window{
					MachineIds: []*mesos.MachineID{
						&mesos.MachineID{
							Hostname: &hostname,
							Ip:       &ip,
						},
					},
					Unavailability: &mesos.Unavailability{
						Start: &mesos.TimeInfo{
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
	call := &master.Call_StartMaintenance{
		Machines: []*mesos.MachineID{
			&mesos.MachineID{
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
	call := &master.Call_StopMaintenance{
		Machines: []*mesos.MachineID{
			&mesos.MachineID{
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

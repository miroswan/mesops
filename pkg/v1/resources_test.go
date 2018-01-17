package v1

import (
	"testing"

	"github.com/mesos/go-proto/mesos/v1"
	"github.com/mesos/go-proto/mesos/v1/master"
)

func TestReserveResources(t *testing.T) {
	s := NewTestProtobufServer(MasterClient)
	defer s.Teardown()

	s.Handle()

	slaveId := "test-slave"
	resourceName := "test-mem"
	resourceValue := mesos_v1.Value_Type(1.0)

	call := &mesos_v1_master.Call_ReserveResources{
		AgentId: &mesos_v1.AgentID{Value: &slaveId},
		Resources: []*mesos_v1.Resource{
			&mesos_v1.Resource{
				Name: &resourceName,
				Type: &resourceValue,
			},
		},
	}

	err := s.Master().ReserveResource(s.Ctx(), call)
	if err != nil {
		t.Error(err)
	}
}

func TestUnreserveResources(t *testing.T) {
	s := NewTestProtobufServer(MasterClient)
	defer s.Teardown()

	s.Handle()

	slaveId := "test-slave"
	resourceName := "test-mem"
	resourceValue := mesos_v1.Value_Type(1.0)

	call := &mesos_v1_master.Call_UnreserveResources{
		AgentId: &mesos_v1.AgentID{Value: &slaveId},
		Resources: []*mesos_v1.Resource{
			&mesos_v1.Resource{
				Name: &resourceName,
				Type: &resourceValue,
			},
		},
	}

	err := s.Master().UnreserveResource(s.Ctx(), call)
	if err != nil {
		t.Error(err)
	}
}

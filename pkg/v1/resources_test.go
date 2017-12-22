package v1

import (
	"testing"

	"github.com/miroswan/mesops/pkg/v1/master"
	"github.com/miroswan/mesops/pkg/v1/mesos"
)

func TestReserveResources(t *testing.T) {
	s := NewTestProtobufServer(MasterClient)
	defer s.Teardown()

	s.Handle()

	slaveId := "test-slave"
	resourceName := "test-mem"
	resourceValue := mesos.Value_Type(1.0)

	call := &master.Call_ReserveResources{
		SlaveId: &mesos.SlaveID{Value: &slaveId},
		Resources: []*mesos.Resource{
			&mesos.Resource{
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
	resourceValue := mesos.Value_Type(1.0)

	call := &master.Call_UnreserveResources{
		SlaveId: &mesos.SlaveID{Value: &slaveId},
		Resources: []*mesos.Resource{
			&mesos.Resource{
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

package v1

import (
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/mesos/go-proto/mesos/v1"
	"github.com/mesos/go-proto/mesos/v1/master"
)

func TestGetRoles(t *testing.T) {
	s := NewTestProtobufServer(MasterClient)
	defer s.Teardown()

	responseType := mesos_v1_master.Response_GET_ROLES
	name := "test-role"
	weight := float64(1.0)
	response := &mesos_v1_master.Response{
		Type: &responseType,
		GetRoles: &mesos_v1_master.Response_GetRoles{
			Roles: []*mesos_v1.Role{
				&mesos_v1.Role{
					Name:   &name,
					Weight: &weight,
				},
			},
		},
	}

	output, err := proto.Marshal(response)
	if err != nil {
		t.Fatal(err)
	}

	s.SetOutput(output).Handle()

	data, err := s.Master().GetRoles(s.Ctx())
	if err != nil {
		t.Fatal(err)
	}

	respName := data.GetGetRoles().GetRoles()[0].GetName()
	respWeight := data.GetGetRoles().GetRoles()[0].GetWeight()

	if name != respName {
		t.Errorf("expected %s, got %s", name, respName)
	}

	if weight != respWeight {
		t.Errorf("expected %f, got %f", weight, respWeight)
	}

}

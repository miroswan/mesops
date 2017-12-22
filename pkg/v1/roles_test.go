package v1

import (
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/miroswan/mesops/pkg/v1/master"
	"github.com/miroswan/mesops/pkg/v1/mesos"
)

func TestGetRoles(t *testing.T) {
	s := NewTestProtobufServer(MasterClient)
	defer s.Teardown()

	responseType := master.Response_GET_ROLES
	name := "test-role"
	weight := float64(1.0)
	response := &master.Response{
		Type: &responseType,
		GetRoles: &master.Response_GetRoles{
			Roles: []*mesos.Role{
				&mesos.Role{
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

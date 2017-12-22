package v1

import (
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/miroswan/mesops/pkg/v1/agent"
	"github.com/miroswan/mesops/pkg/v1/master"
	"github.com/miroswan/mesops/pkg/v1/mesos"
)

func TestMasterGetFlags(t *testing.T) {
	s := NewTestProtobufServer(MasterClient)
	defer s.Teardown()

	responseType := master.Response_GET_FLAGS
	responseName := "authenticate_http_readonly"
	responseValue := "true"
	response := &master.Response{
		Type: &responseType,
		GetFlags: &master.Response_GetFlags{
			Flags: []*mesos.Flag{
				&mesos.Flag{Name: &responseName, Value: &responseValue},
			},
		},
	}

	output, err := proto.Marshal(response)
	if err != nil {
		t.Fatal(err)
	}

	s.SetOutput(output).Handle()

	res, err := s.Master().GetFlags(s.Ctx())
	if err != nil {
		t.Fatal(err)
	}

	if res.GetFlags.Flags[0].GetValue() != responseValue {
		t.Errorf("expected %s, got %s", res.GetFlags.Flags[0].GetValue(), responseValue)
	}
}

func TestAgentGetFlags(t *testing.T) {
	s := NewTestProtobufServer(AgentClient)
	defer s.Teardown()

	responseType := agent.Response_GET_FLAGS
	responseName := "authenticate_http_readonly"
	responseValue := "true"
	response := &agent.Response{
		Type: &responseType,
		GetFlags: &agent.Response_GetFlags{
			Flags: []*mesos.Flag{
				&mesos.Flag{Name: &responseName, Value: &responseValue},
			},
		},
	}

	output, err := proto.Marshal(response)
	if err != nil {
		t.Fatal(err)
	}

	s.SetOutput(output).Handle()

	res, err := s.Agent().GetFlags(s.Ctx())
	if err != nil {
		t.Fatal(err)
	}

	if res.GetFlags.Flags[0].GetValue() != responseValue {
		t.Errorf("expected %s, got %s", res.GetFlags.Flags[0].GetValue(), responseValue)
	}
}

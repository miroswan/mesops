package v1

import (
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/mesos/go-proto/mesos/v1"
	"github.com/mesos/go-proto/mesos/v1/agent"
	"github.com/mesos/go-proto/mesos/v1/master"
)

func TestMasterGetFlags(t *testing.T) {
	s := NewTestProtobufServer(MasterClient)
	defer s.Teardown()

	responseType := mesos_v1_master.Response_GET_FLAGS
	responseName := "authenticate_http_readonly"
	responseValue := "true"
	response := &mesos_v1_master.Response{
		Type: &responseType,
		GetFlags: &mesos_v1_master.Response_GetFlags{
			Flags: []*mesos_v1.Flag{
				&mesos_v1.Flag{Name: &responseName, Value: &responseValue},
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

	responseType := mesos_v1_agent.Response_GET_FLAGS
	responseName := "authenticate_http_readonly"
	responseValue := "true"
	response := &mesos_v1_agent.Response{
		Type: &responseType,
		GetFlags: &mesos_v1_agent.Response_GetFlags{
			Flags: []*mesos_v1.Flag{
				&mesos_v1.Flag{Name: &responseName, Value: &responseValue},
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

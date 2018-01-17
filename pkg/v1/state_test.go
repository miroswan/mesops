package v1

import (
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/mesos/go-proto/mesos/v1/agent"
	"github.com/mesos/go-proto/mesos/v1/master"
)

func TestMasterGetState(t *testing.T) {
	s := NewTestProtobufServer(MasterClient)
	defer s.Teardown()

	responseType := mesos_v1_master.Response_GET_STATE
	response := &mesos_v1_master.Response{
		Type: &responseType,
		// GetTasks, GetExecutors, GetFrameworks, and GetAgents are all optional.
		// TODO Add these in later
		GetState: &mesos_v1_master.Response_GetState{},
	}

	output, err := proto.Marshal(response)
	if err != nil {
		t.Fatal(err)
	}

	s.SetOutput(output).Handle()

	data, err := s.Master().GetState(s.Ctx())
	if err != nil {
		t.Fatal(err)
	}

	if *data.Type != mesos_v1_master.Response_GET_STATE {
		t.Errorf("expected %v, got %v", mesos_v1_master.Response_GET_STATE, *data.Type)
	}
}

func TestAgentGetState(t *testing.T) {
	s := NewTestProtobufServer(MasterClient)
	defer s.Teardown()

	responseType := mesos_v1_agent.Response_GET_STATE
	response := &mesos_v1_agent.Response{
		Type: &responseType,
		// GetTasks, GetExecutors, GetFrameworks, and GetAgents are all optional.
		// TODO Add these in later
		GetState: &mesos_v1_agent.Response_GetState{},
	}

	output, err := proto.Marshal(response)
	if err != nil {
		t.Fatal(err)
	}

	s.SetOutput(output).Handle()

	data, err := s.Agent().GetState(s.Ctx())
	if err != nil {
		t.Fatal(err)
	}

	if *data.Type != mesos_v1_agent.Response_GET_STATE {
		t.Errorf("expected %v, got %v", mesos_v1_agent.Response_GET_STATE, *data.Type)
	}
}

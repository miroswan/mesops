package v1

import (
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/miroswan/mesops/pkg/v1/agent"
	"github.com/miroswan/mesops/pkg/v1/master"
)

func TestMasterGetState(t *testing.T) {
	s := NewTestProtobufServer(MasterClient)
	defer s.Teardown()

	responseType := master.Response_GET_STATE
	response := &master.Response{
		Type: &responseType,
		// GetTasks, GetExecutors, GetFrameworks, and GetAgents are all optional.
		// TODO Add these in later
		GetState: &master.Response_GetState{},
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

	if *data.Type != master.Response_GET_STATE {
		t.Errorf("expected %v, got %v", master.Response_GET_STATE, *data.Type)
	}
}

func TestAgentGetState(t *testing.T) {
	s := NewTestProtobufServer(MasterClient)
	defer s.Teardown()

	responseType := agent.Response_GET_STATE
	response := &agent.Response{
		Type: &responseType,
		// GetTasks, GetExecutors, GetFrameworks, and GetAgents are all optional.
		// TODO Add these in later
		GetState: &agent.Response_GetState{},
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

	if *data.Type != agent.Response_GET_STATE {
		t.Errorf("expected %v, got %v", agent.Response_GET_STATE, *data.Type)
	}
}

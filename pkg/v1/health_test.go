package v1

import (
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/mesos/go-proto/mesos/v1/agent"
	"github.com/mesos/go-proto/mesos/v1/master"
)

func TestMasterHGetealth(t *testing.T) {
	s := NewTestProtobufServer(MasterClient)
	defer s.Teardown()
	// Setup Response
	callType := mesos_v1_master.Response_GET_HEALTH
	healthy := true
	response := &mesos_v1_master.Response{
		Type:      &callType,
		GetHealth: &mesos_v1_master.Response_GetHealth{Healthy: &healthy},
	}
	output, err := proto.Marshal(response)
	if err != nil {
		t.Fatal(err)
	}

	// Set Response Handler
	s.SetOutput(output).Handle()

	// Call
	data, err := s.Master().GetHealth(s.Ctx())
	if err != nil {
		t.Fatal(err)
	}

	// Assert
	healthyResponse := data.GetGetHealth().GetHealthy()
	if healthy != healthyResponse {
		t.Errorf("expected true: got %b", healthyResponse)
	}
}

func TestAgentGetHealth(t *testing.T) {
	s := NewTestProtobufServer(AgentClient)
	defer s.Teardown()
	// Setup Response
	callType := mesos_v1_agent.Response_GET_HEALTH
	healthy := true
	response := &mesos_v1_agent.Response{
		Type:      &callType,
		GetHealth: &mesos_v1_agent.Response_GetHealth{Healthy: &healthy},
	}
	output, err := proto.Marshal(response)
	if err != nil {
		t.Fatal(err)
	}

	// Set Response Handler
	s.SetOutput(output).Handle()

	// Call
	data, err := s.Agent().GetHealth(s.Ctx())
	if err != nil {
		t.Fatal(err)
	}

	// Assert
	healthyResponse := *data.GetHealth.Healthy
	if healthy != healthyResponse {
		t.Errorf("expected true: got %b", healthyResponse)
	}
}

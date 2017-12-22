package v1

import (
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/miroswan/mesops/pkg/v1/agent"
	"github.com/miroswan/mesops/pkg/v1/master"
)

func TestMasterHGetealth(t *testing.T) {
	s := NewTestProtobufServer(MasterClient)
	defer s.Teardown()
	// Setup Response
	callType := master.Response_GET_HEALTH
	healthy := true
	response := &master.Response{
		Type:      &callType,
		GetHealth: &master.Response_GetHealth{Healthy: &healthy},
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
	callType := agent.Response_GET_HEALTH
	healthy := true
	response := &agent.Response{
		Type:      &callType,
		GetHealth: &agent.Response_GetHealth{Healthy: &healthy},
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

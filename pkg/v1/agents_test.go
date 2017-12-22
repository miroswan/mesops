package v1

import (
	"testing"

	"github.com/gogo/protobuf/proto"

	"github.com/miroswan/mesops/pkg/v1/master"
	"github.com/miroswan/mesops/pkg/v1/mesos"
)

func TestGetAgents(t *testing.T) {
	s := NewTestProtobufServer(MasterClient)
	defer s.Teardown()

	// Setup Response
	active := true
	hostname := "test"
	responseType := master.Response_GET_AGENTS
	version := "1.4.1"
	response := &master.Response{
		Type: &responseType,
		GetAgents: &master.Response_GetAgents{
			Agents: []*master.Response_GetAgents_Agent{
				&master.Response_GetAgents_Agent{
					Active:  &active,
					Version: &version,
					AgentInfo: &mesos.SlaveInfo{
						Hostname: &hostname,
					},
				},
			},
		},
	}

	output, err := proto.Marshal(response)
	if err != nil {
		t.Fatal(err)
	}

	// Setup Handler
	s.SetOutput(output).Handle()

	// Call
	data, err := s.Master().GetAgents(s.Ctx())
	if err != nil {
		t.Fatal(err)
	}
	if !data.GetAgents.Agents[0].GetActive() {
		t.Errorf("Expected true: got %b", data.GetAgents.Agents[0].GetActive())
	}
}

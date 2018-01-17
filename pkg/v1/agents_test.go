package v1

import (
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/mesos/go-proto/mesos/v1"
	"github.com/mesos/go-proto/mesos/v1/master"
)

func TestGetAgents(t *testing.T) {
	s := NewTestProtobufServer(MasterClient)
	defer s.Teardown()

	// Setup Response
	active := true
	hostname := "test"
	responseType := mesos_v1_master.Response_GET_AGENTS
	version := "1.4.1"
	response := &mesos_v1_master.Response{
		Type: &responseType,
		GetAgents: &mesos_v1_master.Response_GetAgents{
			Agents: []*mesos_v1_master.Response_GetAgents_Agent{
				&mesos_v1_master.Response_GetAgents_Agent{
					Active:  &active,
					Version: &version,
					AgentInfo: &mesos_v1.AgentInfo{
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

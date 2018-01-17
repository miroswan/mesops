package v1

import (
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/mesos/go-proto/mesos/v1"
	"github.com/mesos/go-proto/mesos/v1/agent"
	"github.com/mesos/go-proto/mesos/v1/master"
)

func TestMasterGetFramework(t *testing.T) {
	s := NewTestProtobufServer(MasterClient)
	defer s.Teardown()

	frameworkUser := "root"
	frameworkName := "test-framework"
	frameworkActive := true
	frameworkConnected := true
	frameworkRecovered := true

	responseType := mesos_v1_master.Response_GET_FRAMEWORKS
	response := mesos_v1_master.Response{
		Type: &responseType,
		GetFrameworks: &mesos_v1_master.Response_GetFrameworks{
			Frameworks: []*mesos_v1_master.Response_GetFrameworks_Framework{
				&mesos_v1_master.Response_GetFrameworks_Framework{
					FrameworkInfo: &mesos_v1.FrameworkInfo{
						User: &frameworkUser,
						Name: &frameworkName,
					},
					Active:    &frameworkActive,
					Connected: &frameworkConnected,
					Recovered: &frameworkRecovered,
				},
			},
		},
	}

	output, err := proto.Marshal(&response)
	if err != nil {
		t.Fatal(err)
	}

	s.SetOutput(output).Handle()

	data, err := s.Master().GetFrameworks(s.Ctx())
	if err != nil {
		t.Fatal(err)
	}

	user := data.GetGetFrameworks().GetFrameworks()[0].GetFrameworkInfo().GetUser()
	name := data.GetGetFrameworks().GetFrameworks()[0].GetFrameworkInfo().GetName()
	active := data.GetGetFrameworks().GetFrameworks()[0].GetActive()
	connected := data.GetGetFrameworks().GetFrameworks()[0].GetConnected()
	recovered := data.GetGetFrameworks().GetFrameworks()[0].GetRecovered()

	if user != frameworkUser {
		t.Errorf("expected %s, got %s", frameworkUser, user)
	}
	if name != frameworkName {
		t.Errorf("expected %s, got %s", frameworkName, name)
	}
	if active != frameworkActive {
		t.Errorf("expected %s, got %s", frameworkActive, active)
	}
	if connected != frameworkConnected {
		t.Errorf("expected %s, got %s", frameworkConnected, connected)
	}
	if recovered != frameworkRecovered {
		t.Errorf("expected %s, got %s", frameworkRecovered, recovered)
	}
}

func TestAgentGetFramework(t *testing.T) {
	s := NewTestProtobufServer(AgentClient)
	defer s.Teardown()

	frameworkUser := "root"
	frameworkName := "test-framework"

	responseType := mesos_v1_agent.Response_GET_FRAMEWORKS
	response := mesos_v1_agent.Response{
		Type: &responseType,
		GetFrameworks: &mesos_v1_agent.Response_GetFrameworks{
			Frameworks: []*mesos_v1_agent.Response_GetFrameworks_Framework{
				&mesos_v1_agent.Response_GetFrameworks_Framework{
					FrameworkInfo: &mesos_v1.FrameworkInfo{
						User: &frameworkUser,
						Name: &frameworkName,
					},
				},
			},
		},
	}

	output, err := proto.Marshal(&response)
	if err != nil {
		t.Fatal(err)
	}

	s.SetOutput(output).Handle()

	data, err := s.Agent().GetFrameworks(s.Ctx())
	if err != nil {
		t.Fatal(err)
	}

	user := data.GetGetFrameworks().GetFrameworks()[0].GetFrameworkInfo().GetUser()
	name := data.GetGetFrameworks().GetFrameworks()[0].GetFrameworkInfo().GetName()

	if user != frameworkUser {
		t.Errorf("expected %s, got %s", frameworkUser, user)
	}
	if name != frameworkName {
		t.Errorf("expected %s, got %s", frameworkName, name)
	}
}

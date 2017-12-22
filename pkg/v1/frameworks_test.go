package v1

import (
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/miroswan/mesops/pkg/v1/agent"
	"github.com/miroswan/mesops/pkg/v1/master"
	"github.com/miroswan/mesops/pkg/v1/mesos"
)

func TestMasterGetFramework(t *testing.T) {
	s := NewTestProtobufServer(MasterClient)
	defer s.Teardown()

	frameworkUser := "root"
	frameworkName := "test-framework"
	frameworkActive := true
	frameworkConnected := true
	frameworkRecovered := true

	responseType := master.Response_GET_FRAMEWORKS
	response := master.Response{
		Type: &responseType,
		GetFrameworks: &master.Response_GetFrameworks{
			Frameworks: []*master.Response_GetFrameworks_Framework{
				&master.Response_GetFrameworks_Framework{
					FrameworkInfo: &mesos.FrameworkInfo{
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

	responseType := agent.Response_GET_FRAMEWORKS
	response := agent.Response{
		Type: &responseType,
		GetFrameworks: &agent.Response_GetFrameworks{
			Frameworks: []*agent.Response_GetFrameworks_Framework{
				&agent.Response_GetFrameworks_Framework{
					FrameworkInfo: &mesos.FrameworkInfo{
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

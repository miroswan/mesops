package v1

import (
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/miroswan/mesops/pkg/v1/agent"
	"github.com/miroswan/mesops/pkg/v1/master"
	"github.com/miroswan/mesops/pkg/v1/mesos"
)

func TestMasterGetTasks(t *testing.T) {
	s := NewTestProtobufServer(MasterClient)
	defer s.Teardown()

	responseType := master.Response_GET_TASKS
	response := &master.Response{
		Type: &responseType,
		// TODO Fill these in with data
		GetTasks: &master.Response_GetTasks{
			PendingTasks:     []*mesos.Task{},
			Tasks:            []*mesos.Task{},
			UnreachableTasks: []*mesos.Task{},
			CompletedTasks:   []*mesos.Task{},
		},
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

	if *data.Type != master.Response_GET_TASKS {
		t.Errorf("expected %v, got %v", master.Response_GET_TASKS, *data.Type)
	}
}

func TestAgentGetTasks(t *testing.T) {
	s := NewTestProtobufServer(MasterClient)
	defer s.Teardown()

	responseType := agent.Response_GET_TASKS
	response := &agent.Response{
		Type: &responseType,
		// TODO Fill these in with data
		GetTasks: &agent.Response_GetTasks{
			PendingTasks:    []*mesos.Task{},
			QueuedTasks:     []*mesos.Task{},
			LaunchedTasks:   []*mesos.Task{},
			TerminatedTasks: []*mesos.Task{},
			CompletedTasks:  []*mesos.Task{},
		},
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

	if *data.Type != agent.Response_GET_TASKS {
		t.Errorf("expected %v, got %v", agent.Response_GET_TASKS, *data.Type)
	}
}

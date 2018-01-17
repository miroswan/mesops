package v1

import (
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/mesos/go-proto/mesos/v1"
	"github.com/mesos/go-proto/mesos/v1/agent"
	"github.com/mesos/go-proto/mesos/v1/master"
)

func TestMasterGetTasks(t *testing.T) {
	s := NewTestProtobufServer(MasterClient)
	defer s.Teardown()

	responseType := mesos_v1_master.Response_GET_TASKS
	response := &mesos_v1_master.Response{
		Type: &responseType,
		// TODO Fill these in with data
		GetTasks: &mesos_v1_master.Response_GetTasks{
			PendingTasks:     []*mesos_v1.Task{},
			Tasks:            []*mesos_v1.Task{},
			UnreachableTasks: []*mesos_v1.Task{},
			CompletedTasks:   []*mesos_v1.Task{},
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

	if *data.Type != mesos_v1_master.Response_GET_TASKS {
		t.Errorf("expected %v, got %v", mesos_v1_master.Response_GET_TASKS, *data.Type)
	}
}

func TestAgentGetTasks(t *testing.T) {
	s := NewTestProtobufServer(MasterClient)
	defer s.Teardown()

	responseType := mesos_v1_agent.Response_GET_TASKS
	response := &mesos_v1_agent.Response{
		Type: &responseType,
		// TODO Fill these in with data
		GetTasks: &mesos_v1_agent.Response_GetTasks{
			PendingTasks:    []*mesos_v1.Task{},
			QueuedTasks:     []*mesos_v1.Task{},
			LaunchedTasks:   []*mesos_v1.Task{},
			TerminatedTasks: []*mesos_v1.Task{},
			CompletedTasks:  []*mesos_v1.Task{},
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

	if *data.Type != mesos_v1_agent.Response_GET_TASKS {
		t.Errorf("expected %v, got %v", mesos_v1_agent.Response_GET_TASKS, *data.Type)
	}
}

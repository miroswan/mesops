package v1

import (
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/mesos/go-proto/mesos/v1"
	"github.com/mesos/go-proto/mesos/v1/agent"
	"github.com/mesos/go-proto/mesos/v1/master"
)

func TestMasterGetExecutors(t *testing.T) {
	s := NewTestProtobufServer(MasterClient)
	defer s.Teardown()

	callType := mesos_v1_master.Response_GET_EXECUTORS
	slaveID := "test-id"
	executorID := "test-id"
	response := &mesos_v1_master.Response{
		Type: &callType,
		GetExecutors: &mesos_v1_master.Response_GetExecutors{
			Executors: []*mesos_v1_master.Response_GetExecutors_Executor{
				&mesos_v1_master.Response_GetExecutors_Executor{
					ExecutorInfo: &mesos_v1.ExecutorInfo{
						ExecutorId: &mesos_v1.ExecutorID{
							Value: &executorID,
						},
					},
					AgentId: &mesos_v1.AgentID{
						Value: &slaveID,
					},
				},
			},
		},
	}

	output, err := proto.Marshal(response)
	if err != nil {
		t.Fatal(err)
	}
	s.SetOutput(output).Handle()

	// Call
	data, err := s.Master().GetExecutors(s.Ctx())
	if err != nil {
		t.Fatal(err)
	}

	// Set
	slaveIDResponse := data.GetGetExecutors().GetExecutors()[0].GetAgentId().GetValue()

	// Assert
	if slaveID != slaveIDResponse {
		t.Errorf("expected %s: got %s", slaveID, slaveIDResponse)
	}
}

func TestAgentGetExecutors(t *testing.T) {
	s := NewTestProtobufServer(AgentClient)
	defer s.Teardown()

	callType := mesos_v1_agent.Response_GET_EXECUTORS
	slaveID := "test-id"
	executorID := "test-id"
	response := &mesos_v1_agent.Response{
		Type: &callType,
		GetExecutors: &mesos_v1_agent.Response_GetExecutors{
			Executors: []*mesos_v1_agent.Response_GetExecutors_Executor{
				&mesos_v1_agent.Response_GetExecutors_Executor{
					ExecutorInfo: &mesos_v1.ExecutorInfo{
						ExecutorId: &mesos_v1.ExecutorID{
							Value: &executorID,
						},
					},
				},
			},
		},
	}

	output, err := proto.Marshal(response)
	if err != nil {
		t.Fatal(err)
	}
	s.SetOutput(output).Handle()

	// Call
	data, err := s.Agent().GetExecutors(s.Ctx())
	if err != nil {
		t.Fatal(err)
	}

	// Set
	slaveIDResponse := data.GetGetExecutors().
		GetExecutors()[0].
		GetExecutorInfo().
		GetExecutorId().
		GetValue()

	// Assert
	if slaveID != slaveIDResponse {
		t.Errorf("expected %s: got %s", executorID, slaveIDResponse)
	}
}

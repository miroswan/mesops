package v1

import (
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/miroswan/mesops/pkg/v1/agent"
	"github.com/miroswan/mesops/pkg/v1/master"
	"github.com/miroswan/mesops/pkg/v1/mesos"
)

func TestMasterGetExecutors(t *testing.T) {
	s := NewTestProtobufServer(MasterClient)
	defer s.Teardown()

	callType := master.Response_GET_EXECUTORS
	slaveID := "test-id"
	executorID := "test-id"
	response := &master.Response{
		Type: &callType,
		GetExecutors: &master.Response_GetExecutors{
			Executors: []*master.Response_GetExecutors_Executor{
				&master.Response_GetExecutors_Executor{
					ExecutorInfo: &mesos.ExecutorInfo{
						ExecutorId: &mesos.ExecutorID{
							Value: &executorID,
						},
					},
					SlaveId: &mesos.SlaveID{
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
	slaveIDResponse := data.GetGetExecutors().GetExecutors()[0].GetSlaveId().GetValue()

	// Assert
	if slaveID != slaveIDResponse {
		t.Errorf("expected %s: got %s", slaveID, slaveIDResponse)
	}
}

func TestAgentGetExecutors(t *testing.T) {
	s := NewTestProtobufServer(AgentClient)
	defer s.Teardown()

	callType := agent.Response_GET_EXECUTORS
	slaveID := "test-id"
	executorID := "test-id"
	response := &agent.Response{
		Type: &callType,
		GetExecutors: &agent.Response_GetExecutors{
			Executors: []*agent.Response_GetExecutors_Executor{
				&agent.Response_GetExecutors_Executor{
					ExecutorInfo: &mesos.ExecutorInfo{
						ExecutorId: &mesos.ExecutorID{
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

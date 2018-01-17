package v1

import (
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/mesos/go-proto/mesos/v1"
	"github.com/mesos/go-proto/mesos/v1/agent"
)

func TestAgentGetContainers(t *testing.T) {
	s := NewTestProtobufServer(AgentClient)

	// Setup Response
	responseType := mesos_v1_agent.Response_GET_CONTAINERS
	executorName := "fake-executor"
	frameworkID := "fake-framework-id"
	executorID := "fake-executor-id"
	containerID := "fake-container-id"
	response := mesos_v1_agent.Response{
		Type: &responseType,
		GetContainers: &mesos_v1_agent.Response_GetContainers{
			Containers: []*mesos_v1_agent.Response_GetContainers_Container{
				&mesos_v1_agent.Response_GetContainers_Container{
					FrameworkId: &mesos_v1.FrameworkID{
						Value: &frameworkID,
					},
					ExecutorId: &mesos_v1.ExecutorID{
						Value: &executorID,
					},
					ContainerId: &mesos_v1.ContainerID{
						Value: &containerID,
					},
					ExecutorName: &executorName,
				},
			},
		},
	}

	// Marshal to byte string
	output, err := proto.Marshal(&response)
	if err != nil {
		t.Fatal(err)
	}

	// Set Response Handler
	s.SetOutput(output).Handle()

	// Call it
	data, err := s.Agent().GetContainers(s.Ctx())
	if err != nil {
		t.Fatal(err)
	}

	// Test member in response
	executorNameResponse := data.GetContainers.Containers[0].GetExecutorName()
	if executorName != executorNameResponse {
		t.Errorf("expected f0f97041-1860-4b4a-b279-91fec4e0ebd8, got %s", executorNameResponse)
	}
}

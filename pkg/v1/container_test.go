package v1

import (
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/miroswan/mesops/pkg/v1/agent"
	"github.com/miroswan/mesops/pkg/v1/mesos"
)

func TestAgentGetContainers(t *testing.T) {
	s := NewTestProtobufServer(AgentClient)

	// Setup Response
	responseType := agent.Response_GET_CONTAINERS
	executorName := "fake-executor"
	frameworkID := "fake-framework-id"
	executorID := "fake-executor-id"
	containerID := "fake-container-id"
	response := agent.Response{
		Type: &responseType,
		GetContainers: &agent.Response_GetContainers{
			Containers: []*agent.Response_GetContainers_Container{
				&agent.Response_GetContainers_Container{
					FrameworkId: &mesos.FrameworkID{
						Value: &frameworkID,
					},
					ExecutorId: &mesos.ExecutorID{
						Value: &executorID,
					},
					ContainerId: &mesos.ContainerID{
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

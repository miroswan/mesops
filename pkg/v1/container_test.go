package v1

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/mesos/go-proto/mesos/v1"
	"github.com/mesos/go-proto/mesos/v1/agent"
)

func TestAgentGetContainers(t *testing.T) {
	s := NewTestProtobufServer(AgentClient)
	defer s.Teardown()

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

func TestLaunchContainer(t *testing.T) {
	s := NewTestProtobufServer(AgentClient)
	defer s.Teardown()

	containerIDValue := "test-id"
	commandValue := "/bin/true"
	call := &mesos_v1_agent.Call_LaunchContainer{
		ContainerId: &mesos_v1.ContainerID{Value: &containerIDValue},
		Command:     &mesos_v1.CommandInfo{Value: &commandValue},
	}

	s.SetOutput(make([]byte, 0)).Handle()

	err := s.Agent().LaunchContainer(s.Ctx(), call)
	if err != nil {
		t.Error("expected nil, got %s", err)
	}
}

func TestLaunchNestedContainer(t *testing.T) {
	s := NewTestProtobufServer(AgentClient)
	defer s.Teardown()

	containerIDValue := "test-id"
	commandValue := "/bin/true"
	hostname := "test-host"
	image := "test-image"
	dockerType := mesos_v1.ContainerInfo_DOCKER
	call := &mesos_v1_agent.Call_LaunchNestedContainer{
		ContainerId: &mesos_v1.ContainerID{Value: &containerIDValue},
		Command:     &mesos_v1.CommandInfo{Value: &commandValue},
		Container: &mesos_v1.ContainerInfo{
			Type:     &dockerType,
			Hostname: &hostname,
			Docker:   &mesos_v1.ContainerInfo_DockerInfo{Image: &image},
		},
	}
	s.SetOutput(make([]byte, 0)).Handle()
	err := s.Agent().LaunchNestedContainer(s.Ctx(), call)
	if err != nil {
		t.Error("expected nil, got %s", err)
	}
}

func TestWaitNestedContainer(t *testing.T) {
	s := NewTestProtobufServer(AgentClient)
	defer s.Teardown()

	// Request
	containerIDValue := "test-id"
	call := &mesos_v1_agent.Call_WaitNestedContainer{
		ContainerId: &mesos_v1.ContainerID{Value: &containerIDValue},
	}

	// Response
	responseType := mesos_v1_agent.Response_WAIT_NESTED_CONTAINER
	exitStatus := int32(0)
	response := &mesos_v1_agent.Response{
		Type: &responseType,
		WaitNestedContainer: &mesos_v1_agent.Response_WaitNestedContainer{
			ExitStatus: &exitStatus,
		},
	}

	// Marshal to byte string
	output, err := proto.Marshal(response)
	if err != nil {
		t.Fatal(err)
	}

	s.SetOutput(output).Handle()

	data, err := s.Agent().WaitNestedContainer(s.Ctx(), call)
	if err != nil {
		t.Fatal("expected nil, got %s", err)
	}
	if data.GetWaitNestedContainer().GetExitStatus() != int32(0) {
		t.Errorf("expected 0, got %d", data.GetWaitNestedContainer().GetExitStatus())
	}
}

func TestKillNestedContainer(t *testing.T) {
	s := NewTestProtobufServer(AgentClient)
	defer s.Teardown()

	containerIDValue := "test-id"
	call := &mesos_v1_agent.Call_KillNestedContainer{
		ContainerId: &mesos_v1.ContainerID{Value: &containerIDValue},
	}
	s.SetOutput(make([]byte, 0)).Handle()
	err := s.Agent().KillNestedContainer(s.Ctx(), call)
	if err != nil {
		t.Error("expected nil, got %s", err)
	}
}

func TestLaunchNestedContainerSession(t *testing.T) {
	s := NewTestProtobufServer(AgentClient)
	defer s.Teardown()

	// Request
	containerIDValue := "test-id"
	commandValue := "/bin/true"
	hostname := "test-host"
	image := "test-image"
	dockerType := mesos_v1.ContainerInfo_DOCKER
	call := &mesos_v1_agent.Call_LaunchNestedContainerSession{
		ContainerId: &mesos_v1.ContainerID{Value: &containerIDValue},
		Command:     &mesos_v1.CommandInfo{Value: &commandValue},
		Container: &mesos_v1.ContainerInfo{
			Type:     &dockerType,
			Hostname: &hostname,
			Docker:   &mesos_v1.ContainerInfo_DockerInfo{Image: &image},
		},
	}

	// Stdout message
	stdoutValue := []byte("stdout")
	stdoutType := mesos_v1_agent.ProcessIO_Data_STDOUT
	processIOType := mesos_v1_agent.ProcessIO_DATA
	processIOStdout := &mesos_v1_agent.ProcessIO{
		Type: &processIOType,
		Data: &mesos_v1_agent.ProcessIO_Data{
			Type: &stdoutType,
			Data: stdoutValue,
		},
	}
	stdoutData, err := proto.Marshal(processIOStdout)
	if err != nil {
		t.Fatal(err)
	}
	stdoutSizeBytes := []byte(strconv.Itoa(len(stdoutData)))

	// stderr message
	stderrValue := []byte("stderr")
	stderrType := mesos_v1_agent.ProcessIO_Data_STDERR
	processIOStderr := &mesos_v1_agent.ProcessIO{
		Type: &processIOType,
		Data: &mesos_v1_agent.ProcessIO_Data{
			Type: &stderrType,
			Data: stderrValue,
		},
	}
	stderrData, err := proto.Marshal(processIOStderr)
	if err != nil {
		t.Fatal(err)
	}
	stderrSizeBytes := []byte(strconv.Itoa(len(stderrData)))

	// Create message
	msg := make([]byte, 0)
	msg = append(msg, stdoutSizeBytes...)
	msg = append(msg, '\n')
	msg = append(msg, stdoutData...)
	msg = append(msg, stderrSizeBytes...)
	msg = append(msg, '\n')
	msg = append(msg, stderrData...)

	// Set Handler and call
	s.SetOutput(msg).Handle()
	ctx, _ := context.WithTimeout(s.Ctx(), time.Duration(5)*time.Second)

	processIOStream := make(ProcessIOStream)
	processIOCollection := make([]*mesos_v1_agent.ProcessIO, 0)

	// Run in background
	go func() {
		s.Agent().LaunchNestedContainerSession(ctx, call, processIOStream)
	}()

	// Collect results in background
	go func() {
		for {
			select {
			case p := <-processIOStream:
				processIOCollection = append(processIOCollection, p)
			}
		}
	}()

	for len(processIOCollection) != 2 {
		// Wait for both processIO messages to come in
	}

	// Retrieve results
	stdoutResult := string(processIOCollection[0].GetData().GetData())
	stderrResult := string(processIOCollection[1].GetData().GetData())

	// Verify
	if stdoutResult != "stdout" {
		t.Errorf("expected stdout, got %s", stdoutResult)
	}

	if stderrResult != "stderr" {
		t.Errorf("expected stderr, got %s", stderrResult)
	}
}

func TestAttachContainerInput(t *testing.T) {
	s := NewTestProtobufServer(AgentClient)
	defer s.Teardown()

	// Request
	containerIDValue := "test-id"
	call := &mesos_v1_agent.Call_AttachContainerInput{
		ContainerId: &mesos_v1.ContainerID{Value: &containerIDValue},
	}

	// Stdin message
	stdinValue := []byte("stdin")
	stdinType := mesos_v1_agent.ProcessIO_Data_STDIN
	processIOType := mesos_v1_agent.ProcessIO_DATA
	processIOStdin := &mesos_v1_agent.ProcessIO{
		Type: &processIOType,
		Data: &mesos_v1_agent.ProcessIO_Data{
			Type: &stdinType,
			Data: stdinValue,
		},
	}
	stdinData, err := proto.Marshal(processIOStdin)
	if err != nil {
		t.Fatal(err)
	}
	stdinSizeBytes := []byte(strconv.Itoa(len(stdinData)))

	// Create message
	msg := make([]byte, 0)
	msg = append(msg, stdinSizeBytes...)
	msg = append(msg, '\n')
	msg = append(msg, stdinData...)

	// Set Handler and call
	s.SetOutput(msg).Handle()
	ctx, _ := context.WithTimeout(s.Ctx(), time.Duration(5)*time.Second)

	processIOStream := make(ProcessIOStream)
	processIOCollection := make([]*mesos_v1_agent.ProcessIO, 0)

	// Run in background
	go func() {
		s.Agent().AttachContainerInput(ctx, call, processIOStream)
	}()

	// Collect results in background
	go func() {
		for {
			select {
			case p := <-processIOStream:
				processIOCollection = append(processIOCollection, p)
			}
		}
	}()

	for len(processIOCollection) != 1 {
		// Wait for both processIO messages to come in
	}

	// Retrieve results
	stdinResult := string(processIOCollection[0].GetData().GetData())

	// Verify
	if stdinResult != "stdin" {
		t.Errorf("expected stdin, got %s", stdinResult)
	}
}

func TestAttachContainerOutput(t *testing.T) {
	s := NewTestProtobufServer(AgentClient)
	defer s.Teardown()

	// Request
	containerIDValue := "test-id"
	call := &mesos_v1_agent.Call_AttachContainerOutput{
		ContainerId: &mesos_v1.ContainerID{Value: &containerIDValue},
	}

	// Stdin message
	stdoutValue := []byte("stdout")
	stdoutType := mesos_v1_agent.ProcessIO_Data_STDOUT
	processIOType := mesos_v1_agent.ProcessIO_DATA
	processIOStdout := &mesos_v1_agent.ProcessIO{
		Type: &processIOType,
		Data: &mesos_v1_agent.ProcessIO_Data{
			Type: &stdoutType,
			Data: stdoutValue,
		},
	}
	stdoutData, err := proto.Marshal(processIOStdout)
	if err != nil {
		t.Fatal(err)
	}
	stdoutSizeBytes := []byte(strconv.Itoa(len(stdoutData)))

	// Create message
	msg := make([]byte, 0)
	msg = append(msg, stdoutSizeBytes...)
	msg = append(msg, '\n')
	msg = append(msg, stdoutData...)

	// Set Handler and call
	s.SetOutput(msg).Handle()
	ctx, _ := context.WithTimeout(s.Ctx(), time.Duration(5)*time.Second)

	processIOStream := make(ProcessIOStream)
	processIOCollection := make([]*mesos_v1_agent.ProcessIO, 0)

	// Run in background
	go func() {
		s.Agent().AttachContainerOutput(ctx, call, processIOStream)
	}()

	// Collect results in background
	go func() {
		for {
			select {
			case p := <-processIOStream:
				processIOCollection = append(processIOCollection, p)
			}
		}
	}()

	for len(processIOCollection) != 1 {
		// Wait for both processIO messages to come in
	}

	// Retrieve results
	stdoutResult := string(processIOCollection[0].GetData().GetData())

	// Verify
	if stdoutResult != "stdout" {
		t.Errorf("expected stdout, got %s", stdoutResult)
	}
}

func TestRemoveContainer(t *testing.T) {
	s := NewTestProtobufServer(AgentClient)
	defer s.Teardown()

	containerIDValue := "test-id"
	call := &mesos_v1_agent.Call_RemoveNestedContainer{
		ContainerId: &mesos_v1.ContainerID{Value: &containerIDValue},
	}

	s.SetOutput(make([]byte, 0)).Handle()

	err := s.Agent().RemoveNestedContainer(s.Ctx(), call)
	if err != nil {
		t.Error("expected nil, got %s", err)
	}
}

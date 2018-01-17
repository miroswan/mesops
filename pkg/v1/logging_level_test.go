package v1

import (
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/mesos/go-proto/mesos/v1"
	"github.com/mesos/go-proto/mesos/v1/agent"
	"github.com/mesos/go-proto/mesos/v1/master"
)

func TestMasterGetLoggingLevel(t *testing.T) {
	s := NewTestProtobufServer(MasterClient)
	defer s.Teardown()

	responseType := mesos_v1_master.Response_GET_LOGGING_LEVEL
	responseLoggingLevel := uint32(1)
	response := mesos_v1_master.Response{
		Type: &responseType,
		GetLoggingLevel: &mesos_v1_master.Response_GetLoggingLevel{
			Level: &responseLoggingLevel,
		},
	}

	output, err := proto.Marshal(&response)
	if err != nil {
		t.Fatal(err)
	}

	s.SetOutput(output).Handle()

	data, err := s.Master().GetLoggingLevel(s.Ctx())
	if err != nil {
		t.Fatal(err)
	}

	level := data.GetGetLoggingLevel().GetLevel()
	if level != responseLoggingLevel {
		t.Errorf("expected %d, got %d", responseLoggingLevel, level)
	}
}

func TestAgentGetLoggingLevel(t *testing.T) {
	s := NewTestProtobufServer(AgentClient)
	defer s.Teardown()

	responseType := mesos_v1_agent.Response_GET_LOGGING_LEVEL
	responseLoggingLevel := uint32(1)
	response := mesos_v1_agent.Response{
		Type: &responseType,
		GetLoggingLevel: &mesos_v1_agent.Response_GetLoggingLevel{
			Level: &responseLoggingLevel,
		},
	}

	output, err := proto.Marshal(&response)
	if err != nil {
		t.Fatal(err)
	}

	s.SetOutput(output).Handle()

	data, err := s.Agent().GetLoggingLevel(s.Ctx())
	if err != nil {
		t.Fatal(err)
	}

	level := data.GetGetLoggingLevel().GetLevel()
	if level != responseLoggingLevel {
		t.Errorf("expected %d, got %d", responseLoggingLevel, level)
	}
}

func TestMasterSetLoggingLevel(t *testing.T) {
	s := NewTestProtobufServer(MasterClient)
	defer s.Teardown()

	s.Handle()

	level := uint32(1)
	nanoseconds := int64(2000000)
	err := s.Master().SetLoggingLevel(s.Ctx(), &mesos_v1_master.Call_SetLoggingLevel{
		Level:    &level,
		Duration: &mesos_v1.DurationInfo{Nanoseconds: &nanoseconds},
	})

	if err != nil {
		t.Error(err)
	}
}

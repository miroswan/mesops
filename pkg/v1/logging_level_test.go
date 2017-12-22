package v1

import (
	"testing"

	"github.com/gogo/protobuf/proto"

	"github.com/miroswan/mesops/pkg/v1/agent"
	"github.com/miroswan/mesops/pkg/v1/master"
	"github.com/miroswan/mesops/pkg/v1/mesos"
)

func TestMasterGetLoggingLevel(t *testing.T) {
	s := NewTestProtobufServer(MasterClient)
	defer s.Teardown()

	responseType := master.Response_GET_LOGGING_LEVEL
	responseLoggingLevel := uint32(1)
	response := master.Response{
		Type: &responseType,
		GetLoggingLevel: &master.Response_GetLoggingLevel{
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

	responseType := agent.Response_GET_LOGGING_LEVEL
	responseLoggingLevel := uint32(1)
	response := agent.Response{
		Type: &responseType,
		GetLoggingLevel: &agent.Response_GetLoggingLevel{
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
	err := s.Master().SetLoggingLevel(s.Ctx(), &master.Call_SetLoggingLevel{
		Level:    &level,
		Duration: &mesos.DurationInfo{Nanoseconds: &nanoseconds},
	})

	if err != nil {
		t.Error(err)
	}
}

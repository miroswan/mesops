package v1

import (
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/mesos/go-proto/mesos/v1"
	"github.com/mesos/go-proto/mesos/v1/agent"
	"github.com/mesos/go-proto/mesos/v1/master"
)

func TestMasterGetMetrics(t *testing.T) {
	s := NewTestProtobufServer(MasterClient)
	defer s.Teardown()

	responseType := mesos_v1_master.Response_GET_METRICS
	name := "fake-metric"
	value := 1.0
	response := &mesos_v1_master.Response{
		Type: &responseType,
		GetMetrics: &mesos_v1_master.Response_GetMetrics{
			Metrics: []*mesos_v1.Metric{
				&mesos_v1.Metric{
					Name:  &name,
					Value: &value,
				},
			},
		},
	}

	output, err := proto.Marshal(response)
	if err != nil {
		t.Fatal(err)
	}

	s.SetOutput(output).Handle()

	data, err := s.Master().GetMetrics(s.Ctx())
	if err != nil {
		t.Fatal(err)
	}

	respName := data.GetGetMetrics().GetMetrics()[0].GetName()
	respValue := data.GetGetMetrics().GetMetrics()[0].GetValue()

	if name != respName {
		t.Fatal("expected %s, got %s", name, respName)
	}

	if value != respValue {
		t.Fatal("expected %s, got %s", value, respValue)
	}
}

func TestAgentGetMetrics(t *testing.T) {
	s := NewTestProtobufServer(AgentClient)
	defer s.Teardown()

	responseType := mesos_v1_agent.Response_GET_METRICS
	name := "fake-metric"
	value := 1.0
	response := &mesos_v1_agent.Response{
		Type: &responseType,
		GetMetrics: &mesos_v1_agent.Response_GetMetrics{
			Metrics: []*mesos_v1.Metric{
				&mesos_v1.Metric{
					Name:  &name,
					Value: &value,
				},
			},
		},
	}

	output, err := proto.Marshal(response)
	if err != nil {
		t.Fatal(err)
	}

	s.SetOutput(output).Handle()

	data, err := s.Agent().GetMetrics(s.Ctx())
	if err != nil {
		t.Fatal(err)
	}

	respName := data.GetGetMetrics().GetMetrics()[0].GetName()
	respValue := data.GetGetMetrics().GetMetrics()[0].GetValue()

	if name != respName {
		t.Fatal("expected %s, got %s", name, respName)
	}

	if value != respValue {
		t.Fatal("expected %s, got %s", value, respValue)
	}

}

package v1

import (
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/mesos/go-proto/mesos/v1"
	"github.com/mesos/go-proto/mesos/v1/agent"
	"github.com/mesos/go-proto/mesos/v1/master"
)

func TestMasterGetVersion(t *testing.T) {
	// Setup
	s := NewTestProtobufServer(MasterClient)
	defer s.Teardown()

	// Create Mock Response
	masterCallType := mesos_v1_master.Response_GET_VERSION
	version := "1.0.0"
	buildDate := "2016-06-24 23:18:37"
	buildTime := 1466810317.0
	buildUser := "root"
	var masterResponse = &mesos_v1_master.Response{
		Type: &masterCallType,
		GetVersion: &mesos_v1_master.Response_GetVersion{
			VersionInfo: &mesos_v1.VersionInfo{
				Version:   &version,
				BuildDate: &buildDate,
				BuildTime: &buildTime,
				BuildUser: &buildUser,
			},
		},
	}
	output, err := proto.Marshal(masterResponse)
	if err != nil {
		t.Fatal(err)
	}

	// Set Response Handler
	s.SetOutput(output).Handle()

	// Call
	res, err := s.Master().GetVersion(s.Ctx())
	if err != nil {
		t.Fatal(err)
	}
	// Assert
	if *res.GetVersion.VersionInfo.Version != version {
		t.Errorf("expected 1.0.0: got %s", *res.GetVersion.VersionInfo.Version)
	}
	if *res.GetVersion.VersionInfo.BuildDate != buildDate {
		t.Errorf("expected 2016-06-24 23:18:37: got %s", *res.GetVersion.VersionInfo.BuildDate)
	}
	if *res.GetVersion.VersionInfo.BuildTime != buildTime {
		t.Errorf("expected 1466810317: got %f", *res.GetVersion.VersionInfo.BuildTime)
	}
	if *res.GetVersion.VersionInfo.BuildUser != buildUser {
		t.Errorf("expected root: got %s", *res.GetVersion.VersionInfo.BuildUser)
	}
}

func TestclientGetVersion(t *testing.T) {
	// Setup
	s := NewTestProtobufServer(AgentClient)
	defer s.Teardown()

	// Create Mock Response
	mesos_v1_agentCallType := mesos_v1_agent.Response_GET_VERSION
	version := "1.0.0"
	buildDate := "2016-06-24 23:18:37"
	buildTime := 1466810317.0
	buildUser := "root"
	var mesos_v1_agentResponse = &mesos_v1_agent.Response{
		Type: &mesos_v1_agentCallType,
		GetVersion: &mesos_v1_agent.Response_GetVersion{
			VersionInfo: &mesos_v1.VersionInfo{
				Version:   &version,
				BuildDate: &buildDate,
				BuildTime: &buildTime,
				BuildUser: &buildUser,
			},
		},
	}
	output, err := proto.Marshal(mesos_v1_agentResponse)
	if err != nil {
		t.Fatal(err)
	}

	// Set Response Handler
	s.SetOutput(output).Handle()

	// Call
	res, err := s.Agent().GetVersion(s.Ctx())
	if err != nil {
		t.Fatal(err)
	}
	// Assert
	if *res.GetVersion.VersionInfo.Version != version {
		t.Errorf("expected 1.0.0: got %s", *res.GetVersion.VersionInfo.Version)
	}
	if *res.GetVersion.VersionInfo.BuildDate != buildDate {
		t.Errorf("expected 2016-06-24 23:18:37: got %s", *res.GetVersion.VersionInfo.BuildDate)
	}
	if *res.GetVersion.VersionInfo.BuildTime != buildTime {
		t.Errorf("expected 1466810317: got %f", *res.GetVersion.VersionInfo.BuildTime)
	}
	if *res.GetVersion.VersionInfo.BuildUser != buildUser {
		t.Errorf("expected root: got %s", *res.GetVersion.VersionInfo.BuildUser)
	}
}

package v1

import (
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/miroswan/mesops/pkg/v1/agent"
	"github.com/miroswan/mesops/pkg/v1/master"
	"github.com/miroswan/mesops/pkg/v1/mesos"
)

func TestMasterGetVersion(t *testing.T) {
	// Setup
	s := NewTestProtobufServer(MasterClient)
	defer s.Teardown()

	// Create Mock Response
	masterCallType := master.Response_GET_VERSION
	version := "1.0.0"
	buildDate := "2016-06-24 23:18:37"
	buildTime := 1466810317.0
	buildUser := "root"
	var masterResponse = &master.Response{
		Type: &masterCallType,
		GetVersion: &master.Response_GetVersion{
			VersionInfo: &mesos.VersionInfo{
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
	agentCallType := agent.Response_GET_VERSION
	version := "1.0.0"
	buildDate := "2016-06-24 23:18:37"
	buildTime := 1466810317.0
	buildUser := "root"
	var agentResponse = &agent.Response{
		Type: &agentCallType,
		GetVersion: &agent.Response_GetVersion{
			VersionInfo: &mesos.VersionInfo{
				Version:   &version,
				BuildDate: &buildDate,
				BuildTime: &buildTime,
				BuildUser: &buildUser,
			},
		},
	}
	output, err := proto.Marshal(agentResponse)
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

package v1

import (
	"context"
	"net/http"
	"testing"
)

func testGetVersion(t *testing.T, c API, mux *http.ServeMux) {
	// Setup Handler
	output := `
  {
    "type": "GET_VERSION",
    "get_version": {
      "version_info": {
        "version": "1.0.0",
        "build_date": "2016-06-24 23:18:37",
        "build_time": 1466810317,
        "build_user": "root"
      }
    }
  }
  `
	SetOutput(mux, output)

	// Call
	getVersion, err := c.GetVersion(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	version := *getVersion.GetVersion.VersionInfo.Version
	buildDate := *getVersion.GetVersion.VersionInfo.BuildDate
	buildTime := *getVersion.GetVersion.VersionInfo.BuildTime
	buildUser := *getVersion.GetVersion.VersionInfo.BuildUser

	// Assert
	if version != "1.0.0" {
		t.Errorf("expected 1.0.0: got %s", version)
	}
	if buildDate != "2016-06-24 23:18:37" {
		t.Errorf("expected 2016-06-24 23:18:37: got %s", buildDate)
	}
	if buildTime != 1466810317.0 {
		t.Errorf("expected 1466810317: got %f", buildTime)
	}
	if buildUser != "root" {
		t.Errorf("expected root: got %s", buildUser)
	}
}

func TestMasterGetVersion(t *testing.T) {
	master, mux, teardown := MasterSetup()
	defer teardown()
	testGetVersion(t, master, mux)
}

func TestclientGetVersion(t *testing.T) {
	agent, mux, teardown := AgentSetup()
	defer teardown()
	testGetVersion(t, agent, mux)
}

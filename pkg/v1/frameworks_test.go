package v1

import (
	"context"
	"testing"
)

func TestMasterGetFrameworks(t *testing.T) {
	master, mux, teardown := MasterSetup()
	defer teardown()

	// Setup Handler
	output := `
  {
    "type": "GET_FRAMEWORKS",
    "get_frameworks": {
      "frameworks": [
        {
          "active": true,
          "connected": true,
          "framework_info": {
            "checkpoint": false,
            "failover_timeout": 0.0,
            "hostname": "myhost",
            "id": {
              "value": "361be53a-4d1b-42c1-bec3-e3979eff90bd-0000"
            },
            "name": "default",
            "principal": "my-principal",
            "role": "*",
            "user": "root"
          },
          "registered_time": {
            "nanoseconds": 1470820171578306816
          },
          "reregistered_time": {
            "nanoseconds": 1470820171578306816
          }
        }
      ]
    }
  }
  `
	SetOutput(mux, output)

	// Call
	data, err := master.GetFrameworks(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	// Assert
	id := *data.GetFrameworks.Frameworks[0].FrameworkInfo.ID.Value
	if id != "361be53a-4d1b-42c1-bec3-e3979eff90bd-0000" {
		t.Errorf("expected 361be53a-4d1b-42c1-bec3-e3979eff90bd-0000: got %s", id)
	}
	name := *data.GetFrameworks.Frameworks[0].FrameworkInfo.Name
	if name != "default" {
		t.Errorf("expected default: got %s", name)
	}
	registeredTime := *data.GetFrameworks.Frameworks[0].RegisteredTime.Nanoseconds
	if registeredTime != 1470820171578306816 {
		t.Errorf("expected 1470820171578306816: got %d", registeredTime)
	}
}

func TestAgentGetFrameworks(t *testing.T) {
	api, mux, teardown := AgentSetup()
	defer teardown()

	// Setup Handler
	output := `
	{
	  "type": "GET_FRAMEWORKS",
	  "get_frameworks": {
	    "frameworks": [
	      {
	        "framework_info": {
	          "checkpoint": false,
	          "failover_timeout": 0.0,
	          "hostname": "myhost",
	          "id": {
	            "value": "17e8c0d4-5ee2-4937-bc1c-06c39eddb004-0000"
	          },
	          "name": "default",
	          "principal": "my-principal",
	          "role": "*",
	          "user": "root"
	        }
	      }
	    ]
	  }
	}
  `
	SetOutput(mux, output)

	// Call
	data, err := api.GetFrameworks(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	// Assert
	checkpoint := *data.GetFrameworks.Frameworks[0].FrameworkInfo.Checkpoint
	id := *data.GetFrameworks.Frameworks[0].FrameworkInfo.ID.Value
	if checkpoint != false {
		t.Errorf("expected false: got %b", checkpoint)
	}
	if id != "17e8c0d4-5ee2-4937-bc1c-06c39eddb004-0000" {
		t.Errorf("expected 17e8c0d4-5ee2-4937-bc1c-06c39eddb004-0000: got %s", id)
	}
}

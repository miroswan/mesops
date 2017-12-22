package v1

import (
	"context"
	"testing"
)

func TestMasterGetExecutors(t *testing.T) {
	client, mux, teardown := MasterSetup()
	defer teardown()

	// Setup Handler
	output := `
  {
    "type": "GET_EXECUTORS",
    "get_executors": {
      "executors": [
        {
          "agent_id": {
            "value": "f2ddc41d-6284-405e-8642-34953093140f-S0"
          },
          "executor_info": {
            "command": {
              "shell": true,
              "value": "exit 1"
            },
            "executor_id": {
              "value": "default"
            },
            "framework_id": {
              "value": "f2ddc41d-6284-405e-8642-34953093140f-0000"
            }
          }
        }
      ]
    }
  }
  `
	SetOutput(mux, output)

	// Call
	getExecutors, err := client.GetExecutors(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	// Set
	agentID := *getExecutors.GetExecutors.Executors[0].AgentID.Value
	commandShell := *getExecutors.GetExecutors.Executors[0].ExecutorInfo.Command.Shell
	commandValue := *getExecutors.GetExecutors.Executors[0].ExecutorInfo.Command.Value
	executorID := *getExecutors.GetExecutors.Executors[0].ExecutorInfo.ExecutorID.Value
	frameworkID := *getExecutors.GetExecutors.Executors[0].ExecutorInfo.FrameworkID.Value

	// Assert
	if agentID != "f2ddc41d-6284-405e-8642-34953093140f-S0" {
		t.Errorf("expected f2ddc41d-6284-405e-8642-34953093140f-S0: got %s", agentID)
	}
	if commandShell != true {
		t.Errorf("expected true: got %b", commandShell)
	}
	if commandValue != "exit 1" {
		t.Errorf("expected exit 1: got %s", commandValue)
	}
	if executorID != "default" {
		t.Errorf("expected default: got %s", executorID)
	}
	if frameworkID != "f2ddc41d-6284-405e-8642-34953093140f-0000" {
		t.Errorf("expected f2ddc41d-6284-405e-8642-34953093140f-0000: got %s", frameworkID)
	}
}

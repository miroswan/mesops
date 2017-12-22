package v1

import (
	"context"
	"testing"

	"github.com/miroswan/mesops/pkg/v1/agent"
)

func TestMasterGetTasks(t *testing.T) {
	master, mux, teardown := MasterSetup()
	defer teardown()

	// Setup Handler
	output := `
{
  "type": "GET_TASKS",
  "get_tasks": {
    "tasks": [
      {
        "agent_id": {
          "value": "d4bd102f-e25f-46dc-bb5d-8b10bca133d8-S0"
        },
        "executor_id": {
          "value": "default"
        },
        "framework_id": {
          "value": "d4bd102f-e25f-46dc-bb5d-8b10bca133d8-0000"
        },
        "name": "test",
        "resources": [
          {
            "name": "cpus",
            "role": "*",
            "scalar": {
              "value": 2.0
            },
            "type": "SCALAR"
          },
          {
            "name": "mem",
            "role": "*",
            "scalar": {
              "value": 1024.0
            },
            "type": "SCALAR"
          },
          {
            "name": "disk",
            "role": "*",
            "scalar": {
              "value": 1024.0
            },
            "type": "SCALAR"
          },
          {
            "name": "ports",
            "ranges": {
              "range": [
                {
                  "begin": 31000,
                  "end": 32000
                }
              ]
            },
            "role": "*",
            "type": "RANGES"
          }
        ],
        "state": "TASK_RUNNING",
        "status_update_state": "TASK_RUNNING",
        "status_update_uuid": "ycLTRBo8TjKFTrh4vsBERg==",
        "statuses": [
          {
            "agent_id": {
              "value": "d4bd102f-e25f-46dc-bb5d-8b10bca133d8-S0"
            },
            "container_status": {
              "network_infos": [
                {
                  "ip_addresses": [
                    {
                      "ip_address": "127.0.1.1"
                    }
                  ]
                }
              ]
            },
            "executor_id": {
              "value": "default"
            },
            "source": "SOURCE_EXECUTOR",
            "state": "TASK_RUNNING",
            "task_id": {
              "value": "1"
            },
            "timestamp": 1470820172.32565,
            "uuid": "ycLTRBo8TjKFTrh4vsBERg=="
          }
        ],
        "task_id": {
          "value": "1"
        }
      }
    ]
  }
}
  `
	SetOutput(mux, output)

	// Call
	getTasks, err := master.GetTasks(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	agentID := *getTasks.GetTasks.Tasks[0].AgentID.Value
	executorID := *getTasks.GetTasks.Tasks[0].ExecutorID.Value
	frameworkID := *getTasks.GetTasks.Tasks[0].FrameworkID.Value
	name := *getTasks.GetTasks.Tasks[0].Name
	state := *getTasks.GetTasks.Tasks[0].State
	statusUpdateState := *getTasks.GetTasks.Tasks[0].StatusUpdateState
	statusUpdateUUID := *getTasks.GetTasks.Tasks[0].StatusUpdateUUID
	statusesAgentID := *getTasks.GetTasks.Tasks[0].Statuses[0].AgentID.Value
	agentIPAddress := *getTasks.GetTasks.Tasks[0].Statuses[0].ContainerStatus.NetworkInfos[0].IPAddresses[0].IPAddress
	var cpus float64
	var mem float64
	var disk float64
	for _, resource := range getTasks.GetTasks.Tasks[0].Resources {
		switch *resource.Name {
		case "cpus":
			cpus = *resource.Scalar.Value
		case "mem":
			mem = *resource.Scalar.Value
		case "disk":
			disk = *resource.Scalar.Value
		}
	}

	// Assert
	if agentID != "d4bd102f-e25f-46dc-bb5d-8b10bca133d8-S0" {
		t.Errorf("expected d4bd102f-e25f-46dc-bb5d-8b10bca133d8-S0: got %s", agentID)
	}
	if executorID != "default" {
		t.Errorf("expected default: got %s", executorID)
	}
	if frameworkID != "d4bd102f-e25f-46dc-bb5d-8b10bca133d8-0000" {
		t.Errorf("expected d4bd102f-e25f-46dc-bb5d-8b10bca133d8-0000: got %f", frameworkID)
	}
	if name != "test" {
		t.Errorf("expected test: got %s", name)
	}
	if cpus != 2.0 {
		t.Errorf("expected 2.0: got %f", cpus)
	}
	if mem != 1024.0 {
		t.Errorf("expected 1024.0: got %f", mem)
	}
	if disk != 1024.0 {
		t.Errorf("expected 1024.0: got %f", disk)
	}
	if state != "TASK_RUNNING" {
		t.Errorf("expected TASK_RUNNING: got %s", state)
	}
	if statusUpdateState != "TASK_RUNNING" {
		t.Errorf("expected TASK_RUNNING: got %s", statusUpdateState)
	}
	if statusUpdateUUID != "ycLTRBo8TjKFTrh4vsBERg==" {
		t.Errorf("expected ycLTRBo8TjKFTrh4vsBERg==: got %s", statusUpdateUUID)
	}
	if statusesAgentID != "d4bd102f-e25f-46dc-bb5d-8b10bca133d8-S0" {
		t.Errorf("expcted d4bd102f-e25f-46dc-bb5d-8b10bca133d8-S0: got %s", statusesAgentID)
	}
	if agentIPAddress != "127.0.1.1" {
		t.Errorf("expcted 127.0.1.1: got %s", agentIPAddress)
	}
}

func TestAgentGetTasks(t *testing.T) {
	api, mux, teardown := AgentSetup()
	defer teardown()

	output := `
	{
	  "type": "GET_TASKS",
	  "get_tasks": {
	    "launched_tasks": [
	      {
	        "agent_id": {
	          "value": "70770d61-d666-4547-a808-787f63b00cf2-S0"
	        },
	        "framework_id": {
	          "value": "70770d61-d666-4547-a808-787f63b00cf2-0000"
	        },
	        "name": "",
	        "resources": [
	          {
	            "name": "cpus",
	            "role": "*",
	            "scalar": {
	              "value": 2.0
	            },
	            "type": "SCALAR"
	          },
	          {
	            "name": "mem",
	            "role": "*",
	            "scalar": {
	              "value": 1024.0
	            },
	            "type": "SCALAR"
	          },
	          {
	            "name": "disk",
	            "role": "*",
	            "scalar": {
	              "value": 1024.0
	            },
	            "type": "SCALAR"
	          },
	          {
	            "name": "ports",
	            "ranges": {
	              "range": [
	                {
	                  "begin": 31000,
	                  "end": 32000
	                }
	              ]
	            },
	            "role": "*",
	            "type": "RANGES"
	          }
	        ],
	        "state": "TASK_RUNNING",
	        "status_update_state": "TASK_RUNNING",
	        "status_update_uuid": "0RC72iyRTQefoUL0ClcL0g==",
	        "statuses": [
	          {
	            "agent_id": {
	              "value": "70770d61-d666-4547-a808-787f63b00cf2-S0"
	            },
	            "container_status": {
	              "executor_pid": 27140,
	              "network_infos": [
	                {
	                  "ip_addresses": [
	                    {
	                      "ip_address": "127.0.1.1"
	                    }
	                  ]
	                }
	              ]
	            },
	            "executor_id": {
	              "value": "1"
	            },
	            "source": "SOURCE_EXECUTOR",
	            "state": "TASK_RUNNING",
	            "task_id": {
	              "value": "1"
	            },
	            "timestamp": 1470900791.21577,
	            "uuid": "0RC72iyRTQefoUL0ClcL0g=="
	          }
	        ],
	        "task_id": {
	          "value": "1"
	        }
	      }
	    ]
	  }
	}
	`
	SetOutput(mux, output)

	data, err := api.GetTasks(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	gt := func(i interface{}) interface{} {
		return i
	}(data)

	if _, ok := gt.(*agent.GetTasksResponse); !ok {
		t.Error("the type of the object returned is not an *agent.GetTasks")
	}
}

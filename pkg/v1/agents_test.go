package v1

import (
	"context"
	"testing"
	"time"
)

func TestGetAgents(t *testing.T) {
	master, mux, teardown := MasterSetup()
	defer teardown()

	// Setup Handler
	output := `
{
  "type": "GET_AGENTS",
  "get_agents": {
    "agents": [
      {
        "active": true,
        "agent_info": {
          "hostname": "host",
          "id": {
            "value": "3669ea49-c3c4-4b13-adee-05b8f9cb2562-S0"
          },
          "port": 34626,
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
          ]
        },
        "pid": "slave(1)@127.0.1.1:34626",
        "registered_time": {
          "nanoseconds": 1470820171393027072
        },
        "total_resources": [
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
        "version": "1.1.0"
      }
    ]
  }
}
  `
	SetOutput(mux, output)

	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)

	// Call
	data, err := master.GetAgents(ctx)
	if err != nil {
		t.Fatal(err)
	}

	// Set
	active := *data.GetAgents.Agents[0].Active
	hostname := *data.GetAgents.Agents[0].AgentInfo.Hostname
	id := *data.GetAgents.Agents[0].AgentInfo.ID.Value
	port := *data.GetAgents.Agents[0].AgentInfo.Port
	pid := *data.GetAgents.Agents[0].Pid
	registeredTime := *data.GetAgents.Agents[0].RegisteredTime.Nanoseconds
	version := *data.GetAgents.Agents[0].Version
	var cpus float64
	var mem float64
	var disk float64

	for _, resource := range data.GetAgents.Agents[0].AgentInfo.Resources {
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
	if active != true {
		t.Errorf("expected true: got %b", active)
	}
	if hostname != "host" {
		t.Errorf("expcted host: got %s", hostname)
	}
	if id != "3669ea49-c3c4-4b13-adee-05b8f9cb2562-S0" {
		t.Errorf("expected 3669ea49-c3c4-4b13-adee-05b8f9cb2562-S0: got %s", id)
	}
	if port != 34626 {
		t.Errorf("expected 34626: got %d", port)
	}
	if pid != "slave(1)@127.0.1.1:34626" {
		t.Errorf("expected slave(1)@127.0.1.1:34626: got %s", pid)
	}
	if registeredTime != 1470820171393027072 {
		t.Errorf("expected 1470820171393027072: got %d", registeredTime)
	}
	if version != "1.1.0" {
		t.Errorf("expected 1.1.0: got %s", version)
	}
	if cpus != 2.0 {
		t.Errorf("expcted 2.0: got %f", cpus)
	}
	if mem != 1024.0 {
		t.Errorf("expected 1024.0: got %f", mem)
	}
	if disk != 1024.0 {
		t.Errorf("expected 1024.0: got %f", disk)
	}
}

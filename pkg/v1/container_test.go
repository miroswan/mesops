package v1

import (
	"context"
	"testing"
)

func TestAgentGetContainers(t *testing.T) {
	api, mux, teardown := AgentSetup()
	defer teardown()

	output := `
  {
    "type": "GET_CONTAINERS",
    "get_containers": {
      "containers": [
        {
          "container_id": {
            "value": "f0f97041-1860-4b4a-b279-91fec4e0ebd8"
          },
          "container_status": {
            "network_infos": [
              {
                "ip_addresses": [
                  {
                    "ip_address": "192.168.1.20"
                  }
                ]
              }
            ]
          },
          "executor_id": {
            "value": "default"
          },
          "executor_name": "",
          "framework_id": {
            "value": "cbe3c0f1-5655-4110-bc01-ae658a9dbab9-0000"
          },
          "resource_statistics": {
            "mem_limit_bytes": 2048,
            "timestamp": 0.0
          }
        }
      ]
    }
  }
  `
	SetOutput(mux, output)

	data, err := api.GetContainers(context.Background(), true, true)
	if err != nil {
		t.Fatal(err)
	}
	firstID := *data.GetContainers.Containers[0].ContainerID.Value
	if firstID != "f0f97041-1860-4b4a-b279-91fec4e0ebd8" {
		t.Errorf("expected f0f97041-1860-4b4a-b279-91fec4e0ebd8, got %s", firstID)
	}
}

func TestAgentLaunchNestedContainer(t *testing.T) {

}

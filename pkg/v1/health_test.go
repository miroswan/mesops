package v1

import (
	"context"
	"net/http"
	"testing"
)

func testGetHealth(t *testing.T, client API, mux *http.ServeMux) {
	// Setup Handler
	output := `
  {
    "type": "GET_HEALTH",
    "get_health": {
      "healthy": true
    }
  }
  `
	SetOutput(mux, output)

	// Call
	getHealth, err := client.GetHealth(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	// Assert
	healthy := *getHealth.GetHealth.Healthy
	if healthy != true {
		t.Errorf("expected true: got %b", healthy)
	}
}

func TestMasterHGetealth(t *testing.T) {
	master, mux, teardown := MasterSetup()
	defer teardown()
	testGetHealth(t, master, mux)
}

func TestAgentGetHealth(t *testing.T) {
	agent, mux, teardown := AgentSetup()
	defer teardown()
	testGetHealth(t, agent, mux)
}

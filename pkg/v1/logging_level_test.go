package v1

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

type SetLoggingLevelPayload struct {
	Type            *string `json:"type"`
	SetLoggingLevel *struct {
		Duration *struct {
			Nanoseconds *int64 `json:"nanoseconds"`
		} `json:"duration"`
		Level *int `json:"level"`
	} `json:"set_logging_level"`
}

func TestMasterGetLoggingLevel(t *testing.T) {
	master, mux, teardown := MasterSetup()
	defer teardown()
	testGetLoggingLevel(t, master, mux)
}

func TestAgentGetLoggingLevel(t *testing.T) {
	agent, mux, teardown := AgentSetup()
	defer teardown()
	testGetLoggingLevel(t, agent, mux)
}

func TestMasterSetLoggingLevel(t *testing.T) {
	master, mux, teardown := MasterSetup()
	defer teardown()
	testSetLoggingLevel(t, master, mux)
}

func TestAgentSetLoggingLevel(t *testing.T) {
	agent, mux, teardown := AgentSetup()
	defer teardown()
	testSetLoggingLevel(t, agent, mux)
}

func testGetLoggingLevel(t *testing.T, c API, mux *http.ServeMux) {
	// Setup Handler
	output := `
  {
    "type": "GET_LOGGING_LEVEL",
    "get_logging_level": {
      "level": 0
    }
  }
  `
	SetOutput(mux, output)

	// Call
	getLoggingLevel, err := c.GetLoggingLevel(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	// Assert
	level := *getLoggingLevel.GetLoggingLevel.Level
	if level != 0 {
		t.Errorf("expected 0: got %d", level)
	}
}

func testSetLoggingLevel(t *testing.T, c API, mux *http.ServeMux) {
	mux.HandleFunc("/api/v1", func(rw http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodPost {
			// Read Body
			body := &SetLoggingLevelPayload{}
			b, err := ioutil.ReadAll(req.Body)
			if err != nil {
				t.Error(err)
			}
			// Validate request body
			err = json.Unmarshal(b, body)
			if err != nil {
				t.Error(err)
			}
		}
	})
	err := c.SetLoggingLevel(context.Background(), 1)
	if err != nil {
		t.Fatal(err)
	}
}

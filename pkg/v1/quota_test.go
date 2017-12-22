package v1

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

type RemoveQuotaPayload struct {
	Type        string `json:"type"`
	RemoveQuota struct {
		Role string `json:"role"`
	} `json:"remove_quota"`
}

func TestMasterGetQuota(t *testing.T) {
	master, mux, teardown := MasterSetup()
	defer teardown()

	output := `
  {
    "type": "GET_QUOTA",
    "get_quota": {
      "status": {
        "infos": [
          {
            "guarantee": [
              {
                "name": "cpus",
                "role": "*",
                "scalar": {
                  "value": 1.0
                },
                "type": "SCALAR"
              },
              {
                "name": "mem",
                "role": "*",
                "scalar": {
                  "value": 512.0
                },
                "type": "SCALAR"
              }
            ],
            "principal": "my-principal",
            "role": "role1"
          }
        ]
      }
    }
  }
  `
	SetOutput(mux, output)

	data, err := master.GetQuota(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	principle := *data.GetQuota.Status.Infos[0].Principal
	role := *data.GetQuota.Status.Infos[0].Role
	if principle != "my-principal" {
		t.Errorf("expected my-principal: got %s", principle)
	}
	if role != "role1" {
		t.Errorf("expected role1: got %s", role)
	}
}

func TestMasterRemoveQuota(t *testing.T) {
	master, mux, teardown := MasterSetup()
	defer teardown()

	mux.HandleFunc("/api/v1", func(rw http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodPost {
			b, err := ioutil.ReadAll(req.Body)
			if err != nil {
				t.Fatal(err)
			}
			err = json.Unmarshal(b, &RemoveQuotaPayload{})
			if err != nil {
				t.Fatal(err)
			}
		}
	})
	err := master.RemoveQuota(context.Background(), "test-role")
	if err != nil {
		t.Fatal(err)
	}
}

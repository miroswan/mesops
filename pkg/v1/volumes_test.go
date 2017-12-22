package v1

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestMasterCreateVolumes(t *testing.T) {
	master, mux, teardown := MasterSetup()
	defer teardown()

	mux.HandleFunc("/api/v1", func(rw http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodPost {
			b, err := ioutil.ReadAll(req.Body)
			if err != nil {
				t.Fatal(err)
			}
			err = json.Unmarshal(b, &CreateVolumesPayload{})
			if err != nil {
				t.Fatal(err)
			}
		}
	})

	txt := `
	{
	  "create_volumes": {
	    "agent_id": {
	      "value": "919141a8-b434-4946-86b9-e1b65c8171f6-S0"
	    },
	    "volumes": [
	      {
	        "type": "SCALAR",
	        "disk": {
	          "persistence": {
	            "id": "id1",
	            "principal": "my-principal"
	          },
	          "volume": {
	            "container_path": "path1",
	            "mode": "RW"
	          }
	        },
	        "name": "disk",
	        "role": "role1",
	        "scalar": {
	          "value": 64.0
	        }
	      }
	    ]
	  }
	}
	`
	cv := &CreateVolumes{}
	err := json.Unmarshal([]byte(txt), cv)
	if err != nil {
		t.Fatal(err)
	}

	err = master.CreateVolumes(context.Background(), cv)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMasterDestroyVolumes(t *testing.T) {
	master, mux, teardown := MasterSetup()
	defer teardown()

	mux.HandleFunc("/api/v1", func(rw http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodPost {
			b, err := ioutil.ReadAll(req.Body)
			if err != nil {
				t.Fatal(err)
			}
			err = json.Unmarshal(b, &DestroyVolumesPayload{})
			if err != nil {
				t.Fatal(err)
			}
		}
	})

	txt := `
	{
	  "destroy_volumes": {
	    "agent_id": {
	      "value": "919141a8-b434-4946-86b9-e1b65c8171f6-S0"
	    },
	    "volumes": [
	      {
	        "disk": {
	          "persistence": {
	            "id": "id1",
	            "principal": "my-principal"
	          },
	          "volume": {
	            "container_path": "path1",
	            "mode": "RW"
	          }
	        },
	        "name": "disk",
	        "role": "role1",
	        "scalar": {
	          "value": 64.0
	        },
	        "type": "SCALAR"
	      }
	    ]
	  }
	}
	`
	dv := &DestroyVolumes{}
	err := json.Unmarshal([]byte(txt), dv)
	if err != nil {
		t.Fatal(err)
	}

	err = master.DestroyVolumes(context.Background(), dv)
	if err != nil {
		t.Fatal(err)
	}
}

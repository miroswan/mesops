package v1

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestMasterReserveResources(t *testing.T) {
	master, mux, teardown := MasterSetup()
	defer teardown()

	mux.HandleFunc("/api/v1", func(rw http.ResponseWriter, req *http.Request) {
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			t.Fatal(err)
		}
		err = json.Unmarshal(b, &ReserveResourcesPayload{})
		if err != nil {
			t.Fatal(err)
		}
	})

	txt := `
	{
	  "reserve_resources": {
	    "agent_id": {
	      "value": "1557de7d-547c-48db-b5d3-6bef9c9640ef-S0"
	    },
	    "resources": [
	      {
	        "type": "SCALAR",
	        "name": "cpus",
	        "reservation": {
	          "principal": "my-principal"
	        },
	        "role": "role",
	        "scalar": {
	          "value": 1.0
	        }
	      },
	      {
	        "type": "SCALAR",
	        "name": "mem",
	        "reservation": {
	          "principal": "my-principal"
	        },
	        "role": "role",
	        "scalar": {
	          "value": 512.0
	        }
	      }
	    ]
	  }
	}
	`
	rr := &ReserveResources{}
	err := json.Unmarshal([]byte(txt), rr)
	if err != nil {
		t.Fatal(err)
	}

	err = master.ReserveResources(context.Background(), rr)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMasterUnreserveResources(t *testing.T) {
	master, mux, teardown := MasterSetup()
	defer teardown()

	mux.HandleFunc("/api/v1", func(rw http.ResponseWriter, req *http.Request) {
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			t.Fatal(err)
		}
		err = json.Unmarshal(b, &UnreserveResourcesPayload{})
		if err != nil {
			t.Fatal(err)
		}
	})
	txt := `
	{
	  "unreserve_resources": {
	    "agent_id": {
	      "value": "1557de7d-547c-48db-b5d3-6bef9c9640ef-S0"
	    },
	    "resources": [
	      {
	        "type": "SCALAR",
	        "name": "cpus",
	        "reservation": {
	          "principal": "my-principal"
	        },
	        "role": "role",
	        "scalar": {
	          "value": 1.0
	        }
	      },
	      {
	        "type": "SCALAR",
	        "name": "mem",
	        "reservation": {
	          "principal": "my-principal"
	        },
	        "role": "role",
	        "scalar": {
	          "value": 512.0
	        }
	      }
	    ]
	  }
	}
	`
	rr := &UnreserveResources{}
	err := json.Unmarshal([]byte(txt), rr)
	if err != nil {
		t.Fatal(err)
	}

	err = master.UnreserveResources(context.Background(), rr)
	if err != nil {
		t.Fatal(err)
	}
}

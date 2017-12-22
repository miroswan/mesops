package v1

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestMasterGetWeights(t *testing.T) {
	master, mux, teardown := MasterSetup()
	defer teardown()

	output := `
  {
    "type": "GET_WEIGHTS",
    "get_weights": {
      "weight_infos": [
        {
          "role": "role",
          "weight": 2.0
        }
      ]
    }
  }
  `
	SetOutput(mux, output)

	data, err := master.GetWeights(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	role := *data.GetWeights.WeightInfos[0].Role
	weight := *data.GetWeights.WeightInfos[0].Weight
	if role != "role" {
		t.Errorf("expected role: got %s", role)
	}
	if weight != 2.0 {
		t.Errorf("expected 2.0: got %f", weight)
	}
}

func TestMasterUpdateWeights(t *testing.T) {
	master, mux, teardown := MasterSetup()
	defer teardown()

	mux.HandleFunc("/api/v1", func(rw http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodPost {
			b, err := ioutil.ReadAll(req.Body)
			if err != nil {
				t.Fatal(err)
			}
			err = json.Unmarshal(b, &UpdateWeightsPayload{})
			if err != nil {
				t.Fatal(err)
			}
		}
	})

	txt := `
	{
	  "type": "UPDATE_WEIGHTS",
	  "update_weights": {
	    "weight_infos": [
	      {
	        "role": "role",
	        "weight": 4.0
	      }
	    ]
	  }
	}
	`
	uw := &UpdateWeights{}
	err := json.Unmarshal([]byte(txt), uw)
	if err != nil {
		t.Fatal(err)
	}
	err = master.UpdateWeights(context.Background(), uw)
	if err != nil {
		t.Fatal(err)
	}
}

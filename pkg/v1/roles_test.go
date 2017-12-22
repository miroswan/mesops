package v1

import (
	"context"
	"testing"
)

func TestMasterGetRoles(t *testing.T) {
	master, mux, teardown := MasterSetup()
	defer teardown()

	output := `
  {
    "type": "GET_ROLES",
    "get_roles": {
      "roles": [
        {
          "name": "*",
          "weight": 1.0
        },
        {
          "frameworks": [
            {
              "value": "74bddcbc-4a02-4d64-b291-aed52032055f-0000"
            }
          ],
          "name": "role1",
          "resources": [
            {
              "name": "cpus",
              "role": "role1",
              "scalar": {
                "value": 0.5
              },
              "type": "SCALAR"
            },
            {
              "name": "mem",
              "role": "role1",
              "scalar": {
                "value": 512.0
              },
              "type": "SCALAR"
            },
            {
              "name": "ports",
              "ranges": {
                "range": [
                  {
                    "begin": 31000,
                    "end": 31001
                  }
                ]
              },
              "role": "role1",
              "type": "RANGES"
            },
            {
              "name": "disk",
              "role": "role1",
              "scalar": {
                "value": 1024.0
              },
              "type": "SCALAR"
            }
          ],
          "weight": 2.5
        }
      ]
    }
  }
  `

	SetOutput(mux, output)
	data, err := master.GetRoles(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	firstWeight := *data.GetRoles.Roles[0].Weight
	firstName := *data.GetRoles.Roles[0].Name
	secondFrameworkID := *data.GetRoles.Roles[1].Frameworks[0].Value
	firstResourceValue := *data.GetRoles.Roles[1].Resources[0].Scalar.Value

	if firstWeight != 1.0 {
		t.Errorf("excpected 1.0: got %f", firstWeight)
	}
	if firstName != "*" {
		t.Errorf("expected *: got %s", firstName)
	}
	if secondFrameworkID != "74bddcbc-4a02-4d64-b291-aed52032055f-0000" {
		t.Errorf("expected 74bddcbc-4a02-4d64-b291-aed52032055f-0000: got %s", secondFrameworkID)
	}
	if firstResourceValue != 0.5 {
		t.Errorf("expected 0.5: got %f", firstResourceValue)
	}
}

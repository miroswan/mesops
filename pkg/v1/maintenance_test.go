package v1

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestMasterGetMaintenanceStatus(t *testing.T) {
	master, mux, teardown := MasterSetup()
	defer teardown()

	// Setup Handler
	output := `
  {
    "type": "GET_MAINTENANCE_STATUS",
    "get_maintenance_status": {
      "status": {
        "draining_machines": [
          {
            "id": {
              "ip": "0.0.0.2"
            }
          },
          {
            "id": {
              "hostname": "myhost"
            }
          }
        ]
      }
    }
  }
  `
	SetOutput(mux, output)

	// Call
	getMaintenanceStatus, err := master.GetMaintenanceStatus(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	// Assert
	first_ip := *getMaintenanceStatus.GetMaintenanceStatus.Status.DrainingMachines[0].ID.IP

	if first_ip != "0.0.0.2" {
		t.Errorf("expected 0.0.0.2: got %s", first_ip)
	}
}

func TestMasterGetMaintenanceSchedule(t *testing.T) {
	master, mux, teardown := MasterSetup()
	defer teardown()

	// Setup Handler
	output := `
{
  "type": "GET_MAINTENANCE_SCHEDULE",
  "get_maintenance_schedule": {
    "schedule": {
      "windows": [
        {
          "machine_ids": [
            {
              "hostname": "myhost"
            },
            {
              "ip": "0.0.0.2"
            }
          ],
          "unavailability": {
            "start": {
              "nanoseconds": 1470849373150643200
            }
          }
        }
      ]
    }
  }
}
  `
	SetOutput(mux, output)

	// Call
	data, err := master.GetMaintenanceSchedule(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	// Assert
	firstUnavailability := *data.GetMaintenanceSchedule.Schedule.Windows[0].Unavailability.Start.Nanoseconds

	if firstUnavailability != 1470849373150643200 {
		t.Errorf("expected 1470849373150643200: got %s", firstUnavailability)
	}
}

func TestMasterUpdateMaintenanceSchedule(t *testing.T) {
	master, mux, teardown := MasterSetup()
	defer teardown()

	mux.HandleFunc("/api/v1", func(rw http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodPost {
			b, err := ioutil.ReadAll(req.Body)
			if err != nil {
				t.Fatal(err)
			}
			err = json.Unmarshal(b, &UpdateMaintenanceSchedulePayload{})
			if err != nil {
				t.Fatal(err)
			}
		}
	})
	txt := `
	{
		"update_maintenance_schedule": {
	    "schedule": {
	      "windows": [
	        {
	          "machine_ids": [
	            {
	              "hostname": "myhost"
	            },
	            {
	              "ip": "0.0.0.2"
	            }
	          ],
	          "unavailability": {
	            "start": {
	              "nanoseconds": 1470820233192017920
	            }
	          }
	        }
	      ]
	    }
	  }
	}
	`
	ums := &UpdateMaintenanceSchedule{}
	err := json.Unmarshal([]byte(txt), ums)
	if err != nil {
		t.Fatal(err)
	}
	err = master.UpdateMaintenanceSchedule(context.Background(), ums)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMasterStartMaintenance(t *testing.T) {
	master, mux, teardown := MasterSetup()
	defer teardown()

	mux.HandleFunc("/api/v1", func(rw http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodPost {
			b, err := ioutil.ReadAll(req.Body)
			if err != nil {
				t.Fatal(err)
			}
			json.Unmarshal(b, &StartMaintenancePayload{})
			if err != nil {
				t.Fatal(err)
			}
		}
	})

	txt := `
	{
		"start_maintenance": {
	    "machines": [
	      {
	        "hostname": "myhost",
	        "ip": "0.0.0.3"
	      }
	    ]
	  }
	}
	`
	sm := &StartMaintenance{}
	err := json.Unmarshal([]byte(txt), sm)
	if err != nil {
		t.Fatal(err)
	}
	err = master.StartMaintenance(context.Background(), sm)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMasterStopMaintenance(t *testing.T) {
	master, mux, teardown := MasterSetup()
	defer teardown()

	mux.HandleFunc("/api/v1", func(rw http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodPost {
			b, err := ioutil.ReadAll(req.Body)
			if err != nil {
				t.Fatal(err)
			}
			json.Unmarshal(b, &StopMaintenancePayload{})
			if err != nil {
				t.Fatal(err)
			}
		}
	})

	txt := `
	{
		"stop_maintenance": {
	    "machines": [
	      {
	        "hostname": "myhost",
	        "ip": "0.0.0.3"
	      }
	    ]
	  }
	}
	`
	sm := &StopMaintenance{}
	err := json.Unmarshal([]byte(txt), sm)
	if err != nil {
		t.Fatal(err)
	}
	err = master.StopMaintenance(context.Background(), sm)
	if err != nil {
		t.Fatal(err)
	}
}

package v1

import (
	"context"
	"net/http"
	"testing"
)

func TestNewMasterBuilder(t *testing.T) {
	a := NewMasterBuilder("test-url")
	i := func(i interface{}) interface{} {
		return i
	}(a)
	if _, ok := i.(*MasterBuilder); !ok {
		t.Error("expected returned type to be a pointer to an MasterBuilder")
	}
}

func TestMasterSetHTTPClient(t *testing.T) {
	a := NewMasterBuilder("test-url").SetHTTPClient(http.DefaultClient)
	i := func(i interface{}) interface{} {
		return i
	}(a)
	if _, ok := i.(*MasterBuilder); !ok {
		t.Error("expected returned type to be a pointer to an MasterBuilder")
	}
}

func TestMasterMaxRetries(t *testing.T) {
	a := NewMasterBuilder("test-url").SetMaxRetries(5)
	i := func(i interface{}) interface{} {
		return i
	}(a)
	if _, ok := i.(*MasterBuilder); !ok {
		t.Error("expected returned type to be a pointer to an MasterBuilder")
	}
}

func TestMasterBuild(t *testing.T) {
	a, err := NewMasterBuilder("test-url").Build()
	if err != nil {
		t.Fatal(err)
	}
	i := func(i interface{}) interface{} {
		return i
	}(a)
	if _, ok := i.(*Master); !ok {
		t.Error("expected returned type to be a pointer to an Master")
	}
}

func TestGetMaster(t *testing.T) {
	master, mux, teardown := MasterSetup()
	defer teardown()

	output := `
	{
	  "type": "GET_MASTER",
	  "get_master": {
	    "master_info": {
	      "address": {
	        "hostname": "myhost",
	        "ip": "127.0.1.1",
	        "port": 34626
	      },
	      "hostname": "myhost",
	      "id": "310ffdac-0b73-408d-acf0-2adcd21cb4b7",
	      "ip": 16842879,
	      "pid": "master@127.0.1.1:34626",
	      "port": 34626,
	      "version": "1.1.0"
	    }
	  }
	}
	`
	SetOutput(mux, output)

	data, err := master.GetMaster(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	hostname := *data.GetMaster.MasterInfo.Address.Hostname
	ip := *data.GetMaster.MasterInfo.Address.IP
	port := *data.GetMaster.MasterInfo.Address.Port
	version := *data.GetMaster.MasterInfo.Version

	if hostname != "myhost" {
		t.Errorf("expected myhost: got %s", hostname)
	}
	if ip != "127.0.1.1" {
		t.Errorf("expected 127.0.1.1: got %s", ip)
	}
	if port != 34626 {
		t.Errorf("expected 34626: got %d", port)
	}
	if version != "1.1.0" {
		t.Errorf("expected 1.1.0: got %s", version)
	}

}

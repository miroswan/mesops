package v1

import (
	"net/http"
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/mesos/go-proto/mesos/v1"
	"github.com/mesos/go-proto/mesos/v1/master"
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
	s := NewTestProtobufServer(MasterClient)
	defer s.Teardown()

	responseType := mesos_v1_master.Response_GET_MASTER
	id := "test-id"
	ip, _ := IPv4toUint32("127.0.0.1")
	ip32 := uint32(ip)
	port := uint32(65000)
	response := mesos_v1_master.Response{
		Type: &responseType,
		GetMaster: &mesos_v1_master.Response_GetMaster{
			MasterInfo: &mesos_v1.MasterInfo{
				Id:   &id,
				Ip:   &ip32,
				Port: &port,
			},
		},
	}

	output, err := proto.Marshal(&response)
	if err != nil {
		t.Fatal(err)
	}

	s.SetOutput(output).Handle()

	data, err := s.Master().GetMaster(s.Ctx())
	if err != nil {
		t.Fatal(err)
	}

	resID := data.GetGetMaster().GetMasterInfo().GetId()
	if id != resID {
		t.Errorf("expected %s, got %s", id, resID)
	}
}

package v1

import (
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/mesos/go-proto/mesos/v1"
	"github.com/mesos/go-proto/mesos/v1/agent"
	"github.com/mesos/go-proto/mesos/v1/master"
)

func TestMasterListFiles(t *testing.T) {
	s := NewTestProtobufServer(MasterClient)
	defer s.Teardown()

	// Setup Response
	callType := mesos_v1_master.Response_LIST_FILES
	gid := "root"
	mode := uint32(16877)
	nlink := int32(2)
	path := "one/2"
	size := uint64(4096)
	uid := "root"
	nanoseconds := int64(1470820172000000000)

	response := &mesos_v1_master.Response{
		Type: &callType,
		ListFiles: &mesos_v1_master.Response_ListFiles{
			FileInfos: []*mesos_v1.FileInfo{
				&mesos_v1.FileInfo{
					Gid:   &gid,
					Mode:  &mode,
					Nlink: &nlink,
					Path:  &path,
					Size:  &size,
					Uid:   &uid,
					Mtime: &mesos_v1.TimeInfo{Nanoseconds: &nanoseconds},
				},
			},
		},
	}

	// Marshal to byte string
	output, err := proto.Marshal(response)
	if err != nil {
		t.Fatal(err)
	}

	s.SetOutput(output).Handle()

	// Setup Payload
	respPath := "one"
	newResponse, err := s.Master().ListFiles(s.Ctx(), &mesos_v1_master.Call_ListFiles{Path: &respPath})
	if err != nil {
		t.Fatal(err)
	}
	respMode := newResponse.GetListFiles().GetFileInfos()[0].GetMode()
	respNlink := newResponse.GetListFiles().GetFileInfos()[0].GetNlink()
	respGetPath := newResponse.GetListFiles().GetFileInfos()[0].GetPath()
	respGetSize := newResponse.GetListFiles().GetFileInfos()[0].GetSize()
	respUID := newResponse.GetListFiles().GetFileInfos()[0].GetUid()
	respNanoseconds := newResponse.GetListFiles().GetFileInfos()[0].GetMtime().GetNanoseconds()

	if respMode != mode {
		t.Error("expected %d, got %d", mode, respMode)
	}
	if respNlink != nlink {
		t.Error("expected %d, got %d", nlink, respNlink)
	}
	if respGetPath != path {
		t.Errorf("expected %s, got %s", path, respGetPath)
	}
	if respGetSize != size {
		t.Error("expected %d, got %d", size, respGetSize)
	}
	if respUID != uid {
		t.Error("expected %d, got %d", uid, respUID)
	}
	if respNanoseconds != nanoseconds {
		t.Error("expected %d, got %d", nanoseconds, respNanoseconds)
	}
}

func TestAgentListFiles(t *testing.T) {
	s := NewTestProtobufServer(AgentClient)
	defer s.Teardown()

	// Setup Response
	callType := mesos_v1_agent.Response_LIST_FILES
	gid := "root"
	mode := uint32(16877)
	nlink := int32(2)
	path := "one/2"
	size := uint64(4096)
	uid := "root"
	nanoseconds := int64(1470820172000000000)

	response := &mesos_v1_agent.Response{
		Type: &callType,
		ListFiles: &mesos_v1_agent.Response_ListFiles{
			FileInfos: []*mesos_v1.FileInfo{
				&mesos_v1.FileInfo{
					Gid:   &gid,
					Mode:  &mode,
					Nlink: &nlink,
					Path:  &path,
					Size:  &size,
					Uid:   &uid,
					Mtime: &mesos_v1.TimeInfo{Nanoseconds: &nanoseconds},
				},
			},
		},
	}

	// Marshal to byte string
	output, err := proto.Marshal(response)
	if err != nil {
		t.Fatal(err)
	}

	// Set Response Handler
	s.SetOutput(output).Handle()

	// Setup Payload
	respPath := "one"
	newResponse, err := s.Agent().ListFiles(s.Ctx(), &mesos_v1_agent.Call_ListFiles{Path: &respPath})
	if err != nil {
		t.Fatal(err)
	}
	respMode := newResponse.ListFiles.FileInfos[0].GetMode()
	respNlink := newResponse.ListFiles.FileInfos[0].GetNlink()
	respGetPath := newResponse.ListFiles.FileInfos[0].GetPath()
	respGetSize := newResponse.ListFiles.FileInfos[0].GetSize()
	respUID := newResponse.ListFiles.FileInfos[0].GetUid()
	respNanoseconds := newResponse.ListFiles.FileInfos[0].GetMtime().GetNanoseconds()

	if respMode != mode {
		t.Error("expected %d, got %d", mode, respMode)
	}
	if respNlink != nlink {
		t.Error("expected %d, got %d", nlink, respNlink)
	}
	if respGetPath != path {
		t.Errorf("expected %s, got %s", path, respGetPath)
	}
	if respGetSize != size {
		t.Error("expected %d, got %d", size, respGetSize)
	}
	if respUID != uid {
		t.Error("expected %d, got %d", uid, respUID)
	}
	if respNanoseconds != nanoseconds {
		t.Error("expected %d, got %d", nanoseconds, respNanoseconds)
	}
}

func TestMasterReadFile(t *testing.T) {
	s := NewTestProtobufServer(MasterClient)
	defer s.Teardown()

	// Setup Response
	responseType := mesos_v1_master.Response_READ_FILE
	size := uint64(4)
	response := &mesos_v1_master.Response{
		Type: &responseType,
		ReadFile: &mesos_v1_master.Response_ReadFile{
			Size: &size,
			Data: []byte("test"),
		},
	}

	// Marshal to byte string
	output, err := proto.Marshal(response)
	if err != nil {
		t.Fatal(err)
	}

	// Set Response Handler
	s.SetOutput(output).Handle()

	// Setup Payload
	callPath := "file.txt"
	callOffset := uint64(0)
	call := &mesos_v1_master.Call_ReadFile{Path: &callPath, Offset: &callOffset}

	res, err := s.Master().ReadFile(s.Ctx(), call)
	if err != nil {
		t.Fatal(err)
	}

	respData := string(res.ReadFile.GetData())
	if respData != "test" {
		t.Errorf("expected test, got %s", respData)
	}
}

func TestAgentReadFile(t *testing.T) {
	s := NewTestProtobufServer(AgentClient)
	defer s.Teardown()

	// Setup Response
	responseType := mesos_v1_agent.Response_READ_FILE
	size := uint64(4)
	response := &mesos_v1_agent.Response{
		Type: &responseType,
		ReadFile: &mesos_v1_agent.Response_ReadFile{
			Size: &size,
			Data: []byte("test"),
		},
	}

	// Marshal to byte string
	output, err := proto.Marshal(response)
	if err != nil {
		t.Fatal(err)
	}

	// Set Response Handler
	s.SetOutput(output).Handle()

	// Setup Payload
	callPath := "file.txt"
	callOffset := uint64(0)
	call := &mesos_v1_agent.Call_ReadFile{Path: &callPath, Offset: &callOffset}

	res, err := s.Agent().ReadFile(s.Ctx(), call)
	if err != nil {
		t.Fatal(err)
	}

	respData := string(res.ReadFile.GetData())
	if respData != "test" {
		t.Errorf("expected test, got %s", respData)
	}
}

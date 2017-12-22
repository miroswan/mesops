package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

type ListFilesPayload struct {
	Type      *string `json:"type"`
	ListFiles *struct {
		Path *string `json:"path"`
	} `json:"list_files"`
}

func TestMasterListFiles(t *testing.T) {
	master, mux, teardown := MasterSetup()
	defer teardown()
	testListFiles(t, master, mux)
}

func TestAgentListFiles(t *testing.T) {
	agent, mux, teardown := AgentSetup()
	defer teardown()
	testListFiles(t, agent, mux)
}

func testListFiles(t *testing.T, c API, mux *http.ServeMux) {
	// Setup Handler
	output := `
  {
    "type": "LIST_FILES",
    "list_files": {
      "file_infos": [
        {
          "gid": "root",
          "mode": 16877,
          "mtime": {
            "nanoseconds": 1470820172000000000
          },
          "nlink": 2,
          "path": "one/2",
          "size": 4096,
          "uid": "root"
        },
        {
          "gid": "root",
          "mode": 16877,
          "mtime": {
            "nanoseconds": 1470820172000000000
          },
          "nlink": 2,
          "path": "one/3",
          "size": 4096,
          "uid": "root"
        },
        {
          "gid": "root",
          "mode": 33188,
          "mtime": {
            "nanoseconds": 1470820172000000000
          },
          "nlink": 1,
          "path": "one/two",
          "size": 3,
          "uid": "root"
        }
      ]
    }
  }
  `
	mux.HandleFunc("/api/v1", func(rw http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodPost {
			// Read Body
			body := &ListFilesPayload{}
			b, err := ioutil.ReadAll(req.Body)
			if err != nil {
				t.Fatal(err)
			}
			// Validate request body
			err = json.Unmarshal(b, body)
			if err != nil {
				t.Fatal(err)
			}
			fmt.Fprint(rw, output)
		}
	})

	// Call
	listFiles, err := c.ListFiles(context.Background(), "one/")
	if err != nil {
		t.Fatal(err)
	}

	// Assert
	firstPath := *listFiles.ListFiles.FileInfos[0].Path

	if firstPath != "one/2" {
		t.Errorf("expecrted one/2: got %s", firstPath)
	}
}

func TestMasterReadFile(t *testing.T) {
	master, mux, teardown := MasterSetup()
	defer teardown()
	testReadFile(t, master, mux)
}

func TestAgentReadFile(t *testing.T) {
	master, mux, teardown := MasterSetup()
	defer teardown()
	testReadFile(t, master, mux)
}

func testReadFile(t *testing.T, c API, mux *http.ServeMux) {
	// Setup Handler
	output := `
	{
	  "type": "READ_FILE",
	  "read_file": {
	    "data": "b2R5",
	    "size": 4
	  }
	}
  `
	mux.HandleFunc("/api/v1", func(rw http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodPost {
			// Read Body
			body := &ReadFilePayload{}
			b, err := ioutil.ReadAll(req.Body)
			if err != nil {
				t.Fatal(err)
			}
			// Validate request body
			err = json.Unmarshal(b, body)
			if err != nil {
				t.Fatal(err)
			}
			fmt.Fprint(rw, output)
		}
	})

	txt := `
	{
	  "type": "READ_FILE",
	  "read_file": {
	    "length": 6,
	    "offset": 1,
	    "path": "myname"
	  }
	}
	`

	rf := &ReadFile{}
	err := json.Unmarshal([]byte(txt), rf)
	if err != nil {
		t.Fatal(err)
	}

	// Call
	data, err := c.ReadFile(context.Background(), rf)
	if err != nil {
		t.Fatal(err)
	}

	// Assert
	contents := *data.ReadFile.Data

	if contents != "b2R5" {
		t.Errorf("expecrted b2R5: got %s", contents)
	}
}

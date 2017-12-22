// MIT License
//
// Copyright (c) [2017-2018] [Demitri Swan]
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package v1

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
)

type ListFilesResponse struct {
	Type      *string `json:"type"`
	ListFiles *struct {
		FileInfos []*struct {
			Gid   *string `json:"gid"`
			Mode  *int    `json:"mode"`
			Mtime *struct {
				Nanoseconds *int64 `json:"nanoseconds"`
			} `json:"mtime"`
			Nlink *int    `json:"nlink"`
			Path  *string `json:"path"`
			Size  *int    `json:"size"`
			UID   *string `json:"uid"`
		} `json:"file_infos"`
	} `json:"list_files"`
}

type ReadFilePayload struct {
	Type     *string   `json:"type"`
	ReadFile *ReadFile `json:"read_file"`
}

type ReadFile struct {
	// Legnth is optional. Defaults to the entire length of the file.
	Length *int `json:"length"`
	// Offset is required. Set a pointer to 0 to read from the beginning of the file
	Offset *int `json:"offset"`
	// Path is the virual file path to the file.
	//
	// References:
	//
	// 	* http://mesos.apache.org/documentation/latest/endpoints/files/debug
	Path *string `json:"path"`
}

type ReadFileResponse struct {
	Type     *string `json:"type"`
	ReadFile *struct {
		Data *string `json:"data"`
		Size *int    `json:"size"`
	} `json:"read_file"`
}

// ListFiles returns a pointer to a ListFiles. You must pass a valid virtual
// file path. A mapping of the virtual file paths to actual paths can be found
// at the files/debug endpoint
//
// References:
//
// 	* http://mesos.apache.org/documentation/latest/operator-http-api/#list_files
// 	* http://mesos.apache.org/documentation/latest/endpoints/files/debug
func (m *Master) ListFiles(ctx context.Context, path string) (lf *ListFilesResponse, err error) {
	lf, err = listFiles(ctx, m.client, path)
	return
}

// ListFiles retrieves the file listing for a directory in master. You must pass
// a valid virtual file path.
//
// References:
//
// 	* http://mesos.apache.org/documentation/latest/operator-http-api/#list_files-1
// 	* http://mesos.apache.org/documentation/latest/endpoints/files/debug
func (a *Agent) ListFiles(ctx context.Context, path string) (lf *ListFilesResponse, err error) {
	lf, err = listFiles(ctx, a.client, path)
	return
}

// ReadFile reads data from a file on the master. This call takes the path of the
// file to be read, the offset to start reading, and the maximum number of bytes
// to read. The length member of the ReadFile is optional. The path must be a
// valid virtual file path.
//
// References:
//
// 	* http://mesos.apache.org/documentation/latest/operator-http-api/#read_file
// 	* http://mesos.apache.org/documentation/latest/endpoints/files/debug
func (m *Master) ReadFile(ctx context.Context, readFile *ReadFile) (rf *ReadFileResponse, err error) {
	rf, err = _readFile(ctx, m.client, readFile)
	return
}

// ReadFile reads data from a file on the master. This call takes the path of the
// file to be read, the offset to start reading, and the maximum number of bytes
// to read. The length member of the ReadFile is optional. The path must be a
// valid virtual file path.
//
// References:
//
// 	* http://mesos.apache.org/documentation/latest/operator-http-api/#read_file
// 	* http://mesos.apache.org/documentation/latest/endpoints/files/debug
func (a *Agent) ReadFile(ctx context.Context, readFile *ReadFile) (rf *ReadFileResponse, err error) {
	rf, err = _readFile(ctx, a.client, readFile)
	return
}

func _readFile(ctx context.Context, client *client, readFile *ReadFile) (rf *ReadFileResponse, err error) {
	var b []byte
	b, err = json.Marshal(readFile)
	if err != nil {
		return
	}
	var buf io.Reader = bytes.NewBuffer(b)
	rf = &ReadFileResponse{}
	err = client.doWithRetryAndLoad(ctx, buf, rf)
	return
}

func listFiles(ctx context.Context, client *client, path string) (lf *ListFilesResponse, err error) {
	const txtTmpl string = `
  {
    "type": "LIST_FILES",
    "list_files": {
    	"path": "%s"
    }
  }
  `
	var txt string = fmt.Sprintf(txtTmpl, path)
	var buf io.Reader = bytes.NewBuffer([]byte(txt))
	if err != nil {
		return
	}
	lf = &ListFilesResponse{}
	err = client.doWithRetryAndLoad(ctx, buf, lf)
	return
}

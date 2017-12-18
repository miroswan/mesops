package mesosops

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	version    string = "0.0.0"
	name       string = "mesops"
	apiVersion int    = 1
)

var (
	userAgent             string        = fmt.Sprintf("%s/%s", name, version)
	baseURN               string        = fmt.Sprintf("api/v%d", apiVersion)
	defaultRequestTimeout time.Duration = 15 * time.Second
)

type (
	// Client contains all of the methods required to interact with the
	// Mesos Operator HTTP API
	Client struct {
		client  *http.Client
		BaseURL *url.URL
	}

	ErrorResponse struct {
		*http.Response
		Message string `json:"message"`
	}
)

//
// Client
//

// NewClient returns a pointer to a new Client
func NewClient(httpClient *http.Client, baseURL string) (client *Client, err error) {
	var parsedURL *url.URL
	if httpClient == nil {
		httpClient = http.DefaultClient
		httpClient.Timeout = defaultRequestTimeout
	}
	if !strings.HasSuffix(baseURL, "/") {
		baseURL += "/"
	}
	baseURL += baseURN
	parsedURL, err = url.Parse(baseURL)
	if err != nil {
		return
	}
	client = &Client{client: httpClient, BaseURL: parsedURL}
	return
}

// newRequest returns a *http.Request. The Mesos HTTP API only uses the POST
// method for all requests, so just pass in the uniform resource name (path) and
// a struct to be loaded.
func (c *Client) newRequest(urn string, body interface{}) (request *http.Request, err error) {
	var finalURL *url.URL
	var buf io.ReadWriter
	finalURL, err = c.BaseURL.Parse(urn)
	if err != nil {
		return
	}
	if body != nil {
		buf = new(bytes.Buffer)
		var enc *json.Encoder = json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err = enc.Encode(body)
		if err != nil {
			return
		}
	}
	request, err = http.NewRequest(http.MethodPost, finalURL.String(), buf)
	if err != nil {
		return
	}
	if body != nil {
		request.Header.Set("Content-Type", "application/json")
	}
	request.Header.Set("User-Agent", userAgent)
	return request, nil
}

func (c *Client) do(
	ctx context.Context,
	request *http.Request,
	v interface{},
) (response *http.Response, err error) {
	request = request.WithContext(ctx)
	response, err = c.client.Do(request)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		var e *url.Error
		var ok bool
		// type assertion on e
		if e, ok = err.(*url.Error); ok {
			var url *url.URL
			if url, err = url.Parse(e.URL); err == nil {
				e.URL = url.String()
			}
		}
		err = e
		return
	}
	// Check for non-200 response.
	err = checkResponse(response)
	if err != nil {
		return
	}
	// If it did not result in error, populate v from body if v exists
	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, response.Body)
		} else {
			err = json.NewDecoder(response.Body).Decode(v)
			if err == io.EOF {
				err = nil
			}
		}
	}
	return
}

func checkResponse(response *http.Response) (errorResponse error) {
	if r := response.StatusCode; r >= 200 && r <= 299 {
		return
	}
	errorResponse = &ErrorResponse{Response: response}
	var err error
	var data []byte
	var defaultErrorMsg string = "failed to create ErrorResponse"
	data, err = ioutil.ReadAll(response.Body)
	if err == nil && data != nil {
		err = json.Unmarshal(data, errorResponse)
		if err != nil {
			errorResponse = fmt.Errorf("%s", defaultErrorMsg)
		}
	} else {
		errorResponse = fmt.Errorf("%s", defaultErrorMsg)
	}
	return
}

func (e *ErrorResponse) Error() (errMsg string) {
	errMsg = fmt.Sprintf("%v %v: %d %v",
		e.Response.Request.Method, e.Response.Request.URL.String(),
		e.Response.StatusCode, e.Message)
	return
}

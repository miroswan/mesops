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
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gogo/protobuf/proto"

	"github.com/miroswan/mesops/pkg"
	"github.com/miroswan/mesops/pkg/v1/agent"
	"github.com/miroswan/mesops/pkg/v1/master"
)

type MasterAPI interface {
	GetExecutors(ctx context.Context) (response *master.Response, err error)
	ListFiles(ctx context.Context, call *master.Call_ListFiles) (response *master.Response, err error)
	ReadFile(ctx context.Context, call *master.Call_ReadFile) (response *master.Response, err error)
	GetFlags(ctx context.Context) (response *master.Response, err error)
	GetFrameworks(ctx context.Context) (response *master.Response, err error)
	GetHealth(ctx context.Context) (response *master.Response, err error)
	GetLoggingLevel(ctx context.Context) (response *master.Response, err error)
	SetLoggingLevel(ctx context.Context, call *master.Call_SetLoggingLevel) (err error)
	GetMaintenanceStatus(ctx context.Context) (response *master.Response, err error)
	GetMaintenanceSchedule(ctx context.Context) (response *master.Response, err error)
	UpdateMaintenanceSchedule(ctx context.Context, call *master.Call_UpdateMaintenanceSchedule)
	StartMaintenance(ctx context.Context, call *master.Call_StartMaintenance) (err error)
	StopMaintenance(ctx context.Context, call *master.Call_StopMaintenance) (err error)
	GetMaster(ctx context.Context) (response *master.Response, err error)
	GetMetrics(ctx context.Context) (response *master.Response, err error)
	GetQuota(ctx context.Context) (response *master.Response, err error)
	SetQuota(ctx context.Context, call *master.Call_SetQuota) (err error)
	RemoveQuota(ctx context.Context, call *master.Call_RemoveQuota) (err error)
	ReserveResource(ctx context.Context, call *master.Call_ReserveResources) (err error)
	UnreserveResource(ctx context.Context, call *master.Call_UnreserveResources) (err error)
	GetRoles(ctx context.Context) (response *master.Response, err error)
	GetState(ctx context.Context) (response *master.Response, err error)
	Subscribe(ctx context.Context, es EventStream) (err error)
	GetTasks(ctx context.Context) (response *master.Response, err error)
	GetVersion(ctx context.Context) (response *master.Response, err error)
	CreateVolumes(ctx context.Context, call *master.Call_CreateVolumes) (err error)
	DestroyVolumes(ctx context.Context, call *master.Call_DestroyVolumes) (err error)
	GetWeights(ctx context.Context) (response *master.Response, err error)
}

type AgentAPI interface {
	GetAgents(ctx context.Context) (response *master.Response, err error)
	GetContainers(ctx context.Context) (response *agent.Response, err error)
	LaunchNestedContainer(ctx context.Context, call *agent.Call_LaunchNestedContainer) (err error)
	WaitNestedContainer(ctx context.Context, call agent.Call_WaitNestedContainer)
	KillNestedContainer(ctx context.Context, call agent.Call_KillNestedContainer) (err error)
	GetExecutors(ctx context.Context) (response *agent.Response, err error)
	ListFiles(ctx context.Context, call *agent.Call_ListFiles) (response *agent.Response, err error)
	ReadFile(ctx context.Context, call *agent.Call_ReadFile) (response *agent.Response, err error)
	GetFlags(ctx context.Context) (response *agent.Response, err error)
	GetFrameworks(ctx context.Context) (response *agent.Response, err error)
	GetHealth(ctx context.Context) (response *agent.Response, err error)
	GetLoggingLevel(ctx context.Context) (response *agent.Response, err error)
	SetLoggingLevel(ctx context.Context, call *agent.Call_SetLoggingLevel) (err error)
	GetMetrics(ctx context.Context) (response *agent.Response, err error)
	GetState(ctx context.Context) (response *agent.Response, err error)
	GetTasks(ctx context.Context) (response *agent.Response, err error)
	GetVersion(ctx context.Context) (response *agent.Response, err error)
}

// HTTPError is a custom error type for HTTP errors outside of the 200 range.
type HTTPError struct {
	statusCode int
	msg        string
}

// Error implements the error interface for HTTPError, printing status_code and
// msg information.
func (e HTTPError) Error() string {
	return fmt.Sprintf("request failed: status_code: %d msg: %s", e.statusCode, e.msg)
}

// IPv4toInt64 parses a string in the form of an IPv4 address and returns an
// int64
func IPv4toUint32(s string) (result uint32, err error) {
	var split []string = strings.Split(s, ".")
	var result64 int64
	for i, val := range split {
		var n int64
		n, err = strconv.ParseInt(val, 10, 32)
		if err != nil {
			return
		}
		// e.g. 127.0.0.1   X is 8 bits
		// 127 << (8*3)     127.X.X.X
		// 0   << (8*2)     0.X.X
		// 0   << (8*1)     0.X
		// 1   << (8*0)     1
		result64 |= n << (8 * uint(3-i))
	}
	result = uint32(result64)
	return
}

// Uint32toIPv4 converts a uint32 to a string in IPv4 format
func Uint32toIPv4(i uint32) (result string, err error) {
	var octets []uint = []uint{uint(24), uint(16), uint(8), uint(0)}
	for index, val := range octets {
		tmp := i >> val & 255
		if index == 0 {
			result += strconv.Itoa(int(tmp))
		} else {
			result += fmt.Sprintf(".%s", strconv.Itoa(int(tmp)))
		}
	}
	return
}

// client handles most of the HTTP interactions with the Mesos Operator API. It
// used by the Master and Agent and not often used independently. For an easy
// start, see Master and Agent.
type client struct {
	httpclient *http.Client
	version    *string
	userAgent  *string
	baseURL    *url.URL
	maxRetries *int
}

// clientBuilder is a builder that constructs a pointer to a client. In most
// cases, you'll want a MasterBuilder or AgentBuilder instead.
type clientBuilder struct {
	*client
	serverURL *string
}

// clientBuilder hold a pointer to a client and has setters for optional
// arguments. Its Build method returns the pointer to the constructed client
func newClientBuilder(serverURL string) *clientBuilder {
	return &clientBuilder{client: &client{}, serverURL: &serverURL}
}

// setServerURL ... (see MasterBuilder and AgentBuilder)
func (b *clientBuilder) setServerURL(serverURL *url.URL) *clientBuilder {
	b.client.baseURL = serverURL
	return b
}

// setHTTPClient ... (see MasterBuilder and AgentBuilder)
func (b *clientBuilder) setHTTPclient(httpclient *http.Client) *clientBuilder {
	b.client.httpclient = httpclient
	return b
}

// setMaxRetries ... (see MasterBuilder and AgentBuilder)
func (b *clientBuilder) setMaxRetries(maxRetries int) *clientBuilder {
	b.client.maxRetries = &maxRetries
	return b
}

// build returns a pointer to a constructed client
func (b *clientBuilder) build() (client *client, err error) {
	// Append api path prefix if not present
	if !strings.HasSuffix(*b.serverURL, "api/v1") {
		if !strings.HasSuffix(*b.serverURL, "/") {
			*b.serverURL += "/api/v1"
		} else {
			*b.serverURL += "api/v1"
		}
	}
	var u *url.URL
	u, err = url.Parse(*b.serverURL)
	if err != nil {
		return
	}
	b.setServerURL(u)

	// Set a sane default http.Client if not set
	if b.client.httpclient == nil {
		b.setHTTPclient(http.DefaultClient)
	}
	// Set maxRetries if not set
	if b.client.maxRetries == nil {
		b.setMaxRetries(10)
	}

	// Set UserAgent
	var userAgent string = fmt.Sprintf("mesops/%s", pkg.Version)
	b.userAgent = &userAgent
	client = b.client
	return
}

// doWithRetryAndLoad executes the HTTP POST request and retries up to the configured value of
// MaxRetries. ctx is a context.Context. body is an io.Reader, likely implemented
// as a []byte. interface is the struct to which we are Unmarshaling. In most
// cases, you will
func (c *client) doWithRetryAndLoad(ctx context.Context, body io.Reader, i interface{}) (res *http.Response, err error) {
	var r []int = make([]int, *c.maxRetries+1) // Setup range for retries
	var start time.Time                        // for generating the round trip time
	var elapsed time.Duration
	var backoff *binaryExponentialBackoff = &binaryExponentialBackoff{}
	var req *http.Request
	var errChan chan error = make(chan error)
	go func() {
		for count := range r {

			// If it is not the first request, then wait
			if count != 0 {
				backoff.wait(count)
			}
			// If the round trip time is not set, then set the start time
			if backoff.rtt == nil {
				start = time.Now()
			}
			req, err = c.newRequest(body)
			if err != nil {
				errChan <- err
				return
			}
			res, err = c.do(ctx, req, i)
			// If the round trip time is not set, then calculate the elapsed time and
			// set it to the round trip time. We will use this in later iterations to
			// allow the backoff to wait for the the correct interval.
			if backoff.rtt == nil {
				elapsed = time.Since(start)
				backoff.rtt = &elapsed
			}
			// If there was no error, then return. If there was an HTTPError then do not
			// retry. Only retry on other errors.
			if err == nil {
				errChan <- err
				return
			} else {
				var ok bool
				var httpError HTTPError
				if httpError, ok = err.(HTTPError); ok {
					errChan <- httpError
					return
				}
			}
		}
		errChan <- fmt.Errorf("exceeded %d retries", *c.maxRetries)
		return
	}()
	select {
	case <-ctx.Done():
		err = ctx.Err()
		return
	case err = <-errChan:
		return
	}
}

// do executes the HTTP request and Unmarshals data if i is not nil.
func (c *client) do(ctx context.Context, req *http.Request, i interface{}) (res *http.Response, err error) {
	req = req.WithContext(ctx)
	res, err = c.httpclient.Do(req)

	if err != nil {
		return
	}

	if res.StatusCode > 299 || res.StatusCode < 200 {
		var msg []byte
		msg, _ = ioutil.ReadAll(res.Body)
		err = HTTPError{statusCode: res.StatusCode, msg: string(msg)}
		return
	}

	if i != nil {
		var j []byte
		j, err = ioutil.ReadAll(res.Body)
		if err != nil {
			return
		}
		err = json.Unmarshal(j, i)
	}
	return
}

// newRequest returns a new *http.Request. Each request is a POST that requires
// a payload that must contain the type of request.
func (c *client) newRequest(body io.Reader) (req *http.Request, err error) {
	req, err = http.NewRequest(http.MethodPost, c.baseURL.String(), body)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", *c.userAgent)
	return
}

func (c *client) doProtoWrapper(ctx context.Context, body io.Reader, pb proto.Message) (res *http.Response, err error) {
	var r []int = make([]int, *c.maxRetries+1) // Setup range for retries
	var start time.Time                        // for generating the round trip time
	var elapsed time.Duration
	var backoff *binaryExponentialBackoff = &binaryExponentialBackoff{}
	var errChan chan error = make(chan error)
	go func() {
		var finalErr error
		for count := range r {

			// If it is not the first request, then wait
			if count != 0 {
				backoff.wait(count)
			}
			// If the round trip time is not set, then set the start time
			if backoff.rtt == nil {
				start = time.Now()
			}
			res, err = c.doProto(ctx, body, pb)
			// If the round trip time is not set, then calculate the elapsed time and
			// set it to the round trip time. We will use this in later iterations to
			// allow the backoff to wait for the the correct interval.
			if backoff.rtt == nil {
				elapsed = time.Since(start)
				backoff.rtt = &elapsed
			}
			// If there was no error, then return. If there was an HTTPError then do not
			// retry. Only retry on other errors.
			if err == nil {
				errChan <- err
				return
			} else {
				var ok bool
				var httpError HTTPError
				if httpError, ok = err.(HTTPError); ok {
					errChan <- httpError
					return
				} else {
					finalErr = err
				}
			}
		}
		errChan <- fmt.Errorf("exceeded %d retries: %s", *c.maxRetries, finalErr)
		return
	}()
	select {
	case <-ctx.Done():
		err = ctx.Err()
		return
	case err = <-errChan:
		return
	}
}

func (c *client) doProto(ctx context.Context, body io.Reader, pb proto.Message) (httpRes *http.Response, err error) {
	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, c.baseURL.String(), body)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/x-protobuf")
	req.Header.Set("Accept", "application/x-protobuf")
	req.Header.Set("User-Agent", *c.userAgent)

	req = req.WithContext(ctx)

	httpRes, err = c.httpclient.Do(req)
	if err != nil {
		return
	}

	if httpRes.StatusCode > 299 || httpRes.StatusCode < 200 {
		var msg []byte
		msg, _ = ioutil.ReadAll(httpRes.Body)
		err = HTTPError{statusCode: httpRes.StatusCode, msg: string(msg)}
		return
	}

	if pb != nil {
		var j []byte
		j, err = ioutil.ReadAll(httpRes.Body)
		if err != nil {
			return
		}
		err = proto.Unmarshal(j, pb)
		if err != nil {
			return
		}
	}
	return
}

// binaryExponentialBackoff is a stateful implementation of binary exponential
// backoff
type binaryExponentialBackoff struct {
	// Estimated Round Trip Time to host
	rtt *time.Duration
}

// wait for the amount of time specified by the binary exponential backoff
// algorithm described in section 8.2.4 of RFC 2616
func (b *binaryExponentialBackoff) wait(count int) {
	var retryDuration time.Duration = time.Duration(math.Pow(2, float64(count)))
	time.Sleep(*b.rtt * retryDuration)
}

func simpleRequestPayload(typeStr string) *bytes.Buffer {
	return bytes.NewBuffer([]byte(fmt.Sprintf(`{"type": "%s"}`, typeStr)))
}

func requestPayload(j string) *bytes.Buffer {
	return bytes.NewBuffer([]byte(j))
}

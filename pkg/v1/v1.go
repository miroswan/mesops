// MIT License
//
// Copyright (c) [2017,2018] [Demitri Swan]
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
	"strings"
	"time"

	"github.com/miroswan/mesops/pkg"
	"github.com/miroswan/mesops/pkg/v1/agent"
	"github.com/miroswan/mesops/pkg/v1/master"
)

// API is the common interface between the MasterAPI and AgentAPI. It is not
// commonly used on its own.
type API interface {
	GetFlags(ctx context.Context) (gf *GetFlagsResponse, err error)
	GetHealth(ctx context.Context) (gh *GetHealthResponse, err error)
	GetVersion(ctx context.Context) (gv *GetVersionResponse, err error)
	GetMetrics(ctx context.Context) (gm *GetMetricsResponse, err error)
	GetLoggingLevel(ctx context.Context) (gll *GetLoggingLevelResponse, err error)
	SetLoggingLevel(ctx context.Context, ll int) (err error)
	ListFiles(ctx context.Context, path string) (lf *ListFilesResponse, err error)
	ReadFile(ctx context.Context, readFile *ReadFile) (rf *ReadFileResponse, err error)
}

// MasterAPI is implemented by Master and specifies the interface for a Master
// client.
type MasterAPI interface {
	API
	CreateVolumes(ctx context.Context, createVolumes *CreateVolumes) (err error)
	DestroyVolumes(ctx context.Context, destroyVolumes *DestroyVolumes) (err error)
	GetAgents(ctx context.Context) (ga *GetAgentsResponse, err error)
	GetExecutors(ctx context.Context) (ge *master.GetExecutorsResponse, err error)
	GetMaintenanceSchedule(ctx context.Context) (gs *GetMaintenanceScheduleResponse, err error)
	GetMaintenanceStatus(ctx context.Context) (gs *GetMaintenanceStatusResponse, err error)
	GetMaster(ctx context.Context) (gm *GetMasterResponse, err error)
	GetQuota(ctx context.Context) (gq *GetQuotaResponse, err error)
	GetRoles(ctx context.Context) (gr *GetRolesResponse, err error)
	GetState(ctx context.Context) (gs *master.GetStateResponse, err error)
	GetFrameworks(ctx context.Context) (gf *master.GetFrameworksResponse, err error)
	GetWeights(ctx context.Context) (gw *GetWeightsResponse, err error)
	MarkAgentGone(ctx context.Context, agentUUID string) (err error)
	RemoveQuota(ctx context.Context, role string) (err error)
	ReserveResources(ctx context.Context, reserveResources *ReserveResources) (err error)
	SetQuota(ctx context.Context, setQuota *SetQuota) (err error)
	StartMaintenance(ctx context.Context, startMaintenance *StartMaintenance)
	StopMaintenance(ctx context.Context, stopMaintenance *StopMaintenance) (err error)
	UnreserveResources(ctx context.Context, unreserveResources *UnreserveResources) (err error)
	UpdateMaintenanceSchedule(ctx context.Context, updateMaintenanceSchedule *UpdateMaintenanceSchedule) (err error)
	UpdateWeights(ctx context.Context, updateWeights *UpdateWeights) (err error)
}

// AgentAPI is implemented by Agent and specifies the interface for an Agent
// client
type AgentAPI interface {
	API
	GetState(ctx context.Context) (gs *agent.GetStateResponse, err error)
	GetContainers(ctx context.Context, showNested bool, showStandalone bool) (gc *GetContainersResponse, err error)
	GetExecutors(ctx context.Context) (ge *agent.GetExecutorsResponse, err error)
	GetFrameworks(ctx context.Context) (gf *agent.GetFrameworksResponse, err error)
	KillNestedContainer(ctx context.Context, k *KillNestedContainer) (err error)
	LaunchNestedContainer(ctx context.Context, l *LaunchNestedContainerPayload) (err error)
	WaitNestedContainer(ctx context.Context, w *WaitNestedContainer) (err error, wnc *WaitNestedContainerResponse)
}

// HTTPError is a custom error type for HTTP errors outside of the 200 range.
type HTTPError struct {
	statusCode int
	msg        string
}

type ChanError struct {
	err error
}

func (c ChanError) Error() string {
	return c.err.Error()
}

// Error implements the error interface for HTTPError, printing status_code and
// msg information.
func (e HTTPError) Error() string {
	return fmt.Sprintf("request failed: status_code: %d msg: %s", e.statusCode, e.msg)
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
		b.client.httpclient.Timeout = time.Second * 15
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
func (c *client) doWithRetryAndLoad(ctx context.Context, body io.Reader, i interface{}) (err error) {
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
			err = c.do(ctx, req, i)
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
func (c *client) do(ctx context.Context, req *http.Request, i interface{}) (err error) {
	var res *http.Response
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

	var j []byte
	j, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	if i != nil {
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

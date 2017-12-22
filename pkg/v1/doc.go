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

/*
v1 contains the code necessary to interact with version 1 of the Mesos Operator
API.

The API contains two components: the Master and the Agent. The MasterAPI
contains methods for interacting with the Mesos cluster through the Master. The
AgentAPI contain methods for interacting with an individual Mesos Agent. The
MasterAPI and AgentAPI are interfaces that describe the corresponding Master
and Agent implementations. Both the Master and Agent have builders to safely
construct each corresponding struct.

For example:

  var masterClient *Master
  var agentClient *Agent
  var err error

  // You can customize the masterClient
  masterClient, err = NewMasterBuilder("http://127.0.0.1:5050").
    SetMaxRetries(5).
    SetHTTPClient(http.DefaultClient).
    Build()

  if err != nil {
    log.Fatal(err)
  }

  // You can use the default settings
  agentClient, err = NewAgentBuilder("http://127.0.0.1:5051").Build()

  if err != nil {
    log.Fatal(err)
  }

With clients configured, you can now interact with the API.

For example:

  // Print whether the configured master is healthy
  gh, err := masterClient.GetHealth(context.Background())
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("The master is healthy: %b", *gh.GetHealth.Healthy)

  // Set Quota
  quota := `
  {
    "set_quota": {
      "quota_request": {
        "force": true,
        "guarantee": [
          {
            "name": "cpus",
            "role": "*",
            "scalar": {
              "value": 1.0
            },
            "type": "SCALAR"
          },
          {
            "name": "mem",
            "role": "*",
            "scalar": {
              "value": 512.0
            },
            "type": "SCALAR"
          }
        ],
        "role": "role1"
      }
    }
  }
  `
  sq := &SetQuota{}
  err := json.Unmarshal([]byte(quota), sq)
  if err != nil {
    log.Fatal(err)
  }
  // create a context with a timeout. The client timeout respects each attempt
  // whereas the context timeout will cancel regardless of the number of retries
  // configured. Think of it as a global timeout for the entire call.
  ctx, _ := context.WithTimeout(context.Background(), 2 * time.Second)

  err = masterClient.SetQuota(ctx, sq)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("Successfully set quota for role1")
}

The methods for each client return pointers to structs that represent the decoded
JSON form of the response payload. See Master and Agent for more details on the
methods and types available to you.

For the most part, you should not have to worry about HTTP when using this
client. However, if a request fails with a response code outside of the 200
range, the calling method will return an HTTPError. This error holds the
response code and message, including both in the error message. Furthermore,
each non-array value within a payload and response type is a pointer to a value.
This allows the caller to check for nil values before accessing data within the
unmarshalled payload and response structs.

There are some cases where the Master and Agent share response structs of the
same name, but differ in structure. In this case, the response structs can be
found in github.com/miroswan/mesops/pkg/v1/master and
github.com/miroswan/mesops/pkg/v1/agent.

This package currently does not support calls for events or streaming responses.
*/
package v1

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

  var masterClient *v1.Master
  var mesos_v1_agentClient *v1.Agent
  var err error

  // You can customize the masterClient
  masterClient, err = v1.NewMasterBuilder("http://127.0.0.1:5050").
    SetMaxRetries(5).
    SetHTTPClient(http.DefaultClient).
    Build()

  if err != nil {
    log.Fatal(err)
  }

  // You can use the default settings
  agentClient, err = v1.NewAgentBuilder("http://127.0.0.1:5051").Build()

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
  fmt.Println("The master is healthy: %b", gh.GetGetHealth().GetHealthy())

  // Build Call_SetQuota
  force := false
  role := "test-role"
  resourceName := "test-mem"
  valueType := mesos_v1.Value_Type(1.0)
  resourceValue := 1.0
  call := &mesos_v1_master.Call_SetQuota{
    QuotaRequest: &mesos_v1_quota.QuotaRequest{
    Force: &force,
      Role:  &role,
      Guarantee: []*mesos_v1.Resource{
        &mesos_v1.Resource{
          Name: &resourceName,
          Type: &valueType,
          Scalar: &mesos_v1.Value_Scalar{
            Value: &resourceValue,
          },
        },
      },
    },
  }

  // create a context with a timeout. The client timeout respects each attempt
  // whereas the context timeout will cancel regardless of the number of retries
  // configured. Think of it as a global timeout for the entire call.
  ctx, _ := context.WithTimeout(context.Background(), 2 * time.Second)
	err := masterClient.SetQuota(s.Ctx(), call)
	if err != nil {
		t.Error(err)
	}

  err = masterClient.SetQuota(ctx, sq)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("Successfully set quota for test-role")
}

For the most part, you should not have to worry about HTTP when using this
client. However, if a request fails with a response code outside of the 200
range, the calling method will return an HTTPError. This error holds the
response code and message, including both in the error message.
*/
package v1

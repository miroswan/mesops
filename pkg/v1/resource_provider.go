package v1

import (
	"context"
	"net/http"

	"github.com/gogo/protobuf/proto"
	"github.com/mesos/go-proto/mesos/v1/agent"
)

// AddResourceProviderConfig launches a Local Resource Provider on the agent
// with the specified ResourceProviderInfo.
func (a *Agent) AddResourceProviderConfig(
	ctx context.Context, call *mesos_v1_agent.Call_AddResourceProviderConfig,
) (err error) {
	var httpResponse *http.Response
	var callType mesos_v1_agent.Call_Type = mesos_v1_agent.Call_ADD_RESOURCE_PROVIDER_CONFIG
	var message proto.Message = &mesos_v1_agent.Call{Type: &callType, AddResourceProviderConfig: call}
	httpResponse, err = a.client.makeCall(ctx, message, nil)
	defer httpResponse.Body.Close()
	return
}

// UpdateResourceProviderConfig updates a Local Resource Provider on the agent
// with the specified ResourceProviderInfo.
func (a *Agent) UpdateResourceProviderConfig(
	ctx context.Context, call *mesos_v1_agent.Call_UpdateResourceProviderConfig,
) (err error) {
	var httpResponse *http.Response
	var callType mesos_v1_agent.Call_Type = mesos_v1_agent.Call_UPDATE_RESOURCE_PROVIDER_CONFIG
	var message proto.Message = &mesos_v1_agent.Call{Type: &callType, UpdateResourceProviderConfig: call}
	httpResponse, err = a.client.makeCall(ctx, message, nil)
	defer httpResponse.Body.Close()
	return
}

// RemoveResourceProviderConfig  terminates a given Local Resource Provider on
// the agent and prevents it from being launched again until the config is added
// back. The master and the agent will think the resource provider has
// disconnected, similar to agent disconnection.
//
// If there exists a task that is using the resources provided by the resource
// provider, its execution will not be affected. However, offer operations for
// the local resource provider will not be successful. In fact, if a local
// resource provider is disconnected, the master will rescind the offers related
// to that local resource provider, effectively disallowing frameworks to
// perform operations on the disconnected local resource provider.
//
// The local resource provider can be re-added after its removal using
// ADD_RESOURCE_PROVIDER_CONFIG. Note that removing a local resource provider
// is different than marking a local resource provider as gone, in which case
// the local resource provider will not be allowed to be re-added. Marking a
// local resource provider as gone is not yet supported.
func (a *Agent) RemoveResourceProviderConfig(
	ctx context.Context, call *mesos_v1_agent.Call_RemoveResourceProviderConfig,
) (err error) {
	var httpResponse *http.Response
	var callType mesos_v1_agent.Call_Type = mesos_v1_agent.Call_REMOVE_RESOURCE_PROVIDER_CONFIG
	var message proto.Message = &mesos_v1_agent.Call{Type: &callType, RemoveResourceProviderConfig: call}
	httpResponse, err = a.client.makeCall(ctx, message, nil)
	defer httpResponse.Body.Close()
	return
}

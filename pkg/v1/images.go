package v1

import (
	"context"
	"net/http"

	"github.com/gogo/protobuf/proto"
	"github.com/mesos/go-proto/mesos/v1/agent"
)

// PruneImages itriggers garbage collection for container images. This call can
// only be made when all running containers are launched with Mesos version
// 1.5 or newer. An optional list of excluded images from GC can be speficied
// via prune_images.excluded_images field.
func (a *Agent) PruneImages(
	ctx context.Context, call *mesos_v1_agent.Call_PruneImages,
) (err error) {
	var httpResponse *http.Response
	var callType mesos_v1_agent.Call_Type = mesos_v1_agent.Call_PRUNE_IMAGES
	var message proto.Message = &mesos_v1_agent.Call{Type: &callType, PruneImages: call}
	httpResponse, err = a.client.makeCall(ctx, message, nil)
	defer httpResponse.Body.Close()
	return
}

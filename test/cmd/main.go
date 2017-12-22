package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/miroswan/mesops/pkg/v1"
	"github.com/miroswan/mesops/pkg/v1/master"
)

func main() {
	client, err := v1.NewMasterBuilder("http://192.168.33.10:5050").Build()
	if err != nil {
		log.Fatal(err)
	}
	es := make(v1.EventStream, 0)
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	go func() {
		err := client.Subscribe(ctx, es)
		if err != nil {
			log.Fatal(err)
		}
	}()

	for {
		select {
		case <-ctx.Done():
			break
		case e := <-es:
			switch e.GetType() {
			case master.Event_SUBSCRIBED:
				fmt.Println(e.GetSubscribed().GetGetState())
			case master.Event_TASK_ADDED:
				fmt.Println(e.GetTaskAdded().GetTask())
			case master.Event_TASK_UPDATED:
				fmt.Println(e.GetTaskUpdated().GetState())
			case master.Event_AGENT_ADDED:
				fmt.Println(e.GetAgentAdded().GetAgent())
			case master.Event_AGENT_REMOVED:
				fmt.Println(e.GetAgentRemoved().GetAgentId())
			case master.Event_FRAMEWORK_ADDED:
				fmt.Println(e.GetFrameworkAdded().GetFramework())
			case master.Event_FRAMEWORK_UPDATED:
				fmt.Println(e.GetFrameworkUpdated().GetFramework())
			case master.Event_FRAMEWORK_REMOVED:
				fmt.Println(e.GetFrameworkRemoved().GetFrameworkInfo())
			case master.Event_UNKNOWN:
				fmt.Println("Event unknown")
			}
		}
	}
}

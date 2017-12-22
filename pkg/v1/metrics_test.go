package v1

import (
	"context"
	"net/http"
	"testing"
)

func testGetMetrics(t *testing.T, c API, mux *http.ServeMux) {
	output := `
  {
    "type": "GET_METRICS",
    "get_metrics": {
      "metrics": [
        {
          "name": "allocator/event_queue_dispatches",
          "value": 1.0
        },
        {
          "name": "master/slaves_active",
          "value": 0.0
        },
        {
          "name": "allocator/mesos/resources/cpus/total",
          "value": 0.0
        },
        {
          "name": "master/messages_revive_offers",
          "value": 0.0
        },
        {
          "name": "allocator/mesos/allocation_runs",
          "value": 0.0
        },
        {
          "name": "master/mem_used",
          "value": 0.0
        },
        {
          "name": "master/valid_executor_to_framework_messages",
          "value": 0.0
        },
        {
          "name": "allocator/mesos/resources/mem/total",
          "value": 0.0
        },
        {
          "name": "log/recovered",
          "value": 1.0
        },
        {
          "name": "registrar/registry_size_bytes",
          "value": 123.0
        },
        {
          "name": "master/slaves_inactive",
          "value": 0.0
        },
        {
          "name": "master/messages_unregister_slave",
          "value": 0.0
        },
        {
          "name": "master/gpus_total",
          "value": 0.0
        },
        {
          "name": "master/disk_revocable_total",
          "value": 0.0
        },
        {
          "name": "master/gpus_percent",
          "value": 0.0
        },
        {
          "name": "master/mem_revocable_used",
          "value": 0.0
        },
        {
          "name": "master/slave_shutdowns_completed",
          "value": 0.0
        },
        {
          "name": "master/invalid_status_updates",
          "value": 0.0
        },
        {
          "name": "master/slave_removals",
          "value": 0.0
        },
        {
          "name": "master/messages_status_update",
          "value": 0.0
        },
        {
          "name": "master/messages_framework_to_executor",
          "value": 0.0
        },
        {
          "name": "master/cpus_revocable_percent",
          "value": 0.0
        },
        {
          "name": "master/recovery_slave_removals",
          "value": 0.0
        },
        {
          "name": "master/event_queue_dispatches",
          "value": 0.0
        },
        {
          "name": "master/messages_update_slave",
          "value": 0.0
        },
        {
          "name": "allocator/mesos/resources/mem/offered_or_allocated",
          "value": 0.0
        },
        {
          "name": "master/messages_register_framework",
          "value": 0.0
        },
        {
          "name": "master/cpus_percent",
          "value": 0.0
        },
        {
          "name": "master/slave_reregistrations",
          "value": 0.0
        },
        {
          "name": "master/cpus_revocable_total",
          "value": 0.0
        },
        {
          "name": "master/gpus_revocable_total",
          "value": 0.0
        },
        {
          "name": "master/valid_status_updates",
          "value": 0.0
        },
        {
          "name": "system/load_15min",
          "value": 1.25
        },
        {
          "name": "master/event_queue_http_requests",
          "value": 0.0
        },
        {
          "name": "master/messages_decline_offers",
          "value": 0.0
        },
        {
          "name": "master/tasks_staging",
          "value": 0.0
        },
        {
          "name": "master/messages_register_slave",
          "value": 0.0
        },
        {
          "name": "allocator/mesos/resources/disk/offered_or_allocated",
          "value": 0.0
        },
        {
          "name": "system/mem_free_bytes",
          "value": 2320146432.0
        },
        {
          "name": "system/cpus_total",
          "value": 4.0
        },
        {
          "name": "master/mem_percent",
          "value": 0.0
        },
        {
          "name": "master/event_queue_messages",
          "value": 0.0
        },
        {
          "name": "master/messages_reregister_slave",
          "value": 0.0
        },
        {
          "name": "master/gpus_used",
          "value": 0.0
        },
        {
          "name": "registrar/state_fetch_ms",
          "value": 16.787968
        },
        {
          "name": "master/messages_launch_tasks",
          "value": 0.0
        },
        {
          "name": "master/gpus_revocable_percent",
          "value": 0.0
        },
        {
          "name": "master/disk_percent",
          "value": 0.0
        },
        {
          "name": "system/load_1min",
          "value": 1.74
        },
        {
          "name": "registrar/queued_operations",
          "value": 0.0
        },
        {
          "name": "master/slaves_disconnected",
          "value": 0.0
        },
        {
          "name": "master/invalid_status_update_acknowledgements",
          "value": 0.0
        },
        {
          "name": "system/load_5min",
          "value": 1.65
        },
        {
          "name": "master/tasks_failed",
          "value": 0.0
        },
        {
          "name": "master/slave_registrations",
          "value": 0.0
        },
        {
          "name": "master/frameworks_connected",
          "value": 0.0
        },
        {
          "name": "allocator/mesos/event_queue_dispatches",
          "value": 0.0
        },
        {
          "name": "master/messages_executor_to_framework",
          "value": 0.0
        },
        {
          "name": "system/mem_total_bytes",
          "value": 8057147392.0
        },
        {
          "name": "master/cpus_revocable_used",
          "value": 0.0
        },
        {
          "name": "master/tasks_killing",
          "value": 0.0
        },
        {
          "name": "allocator/mesos/resources/cpus/offered_or_allocated",
          "value": 0.0
        },
        {
          "name": "master/messages_exited_executor",
          "value": 0.0
        },
        {
          "name": "master/valid_status_update_acknowledgements",
          "value": 0.0
        },
        {
          "name": "master/disk_used",
          "value": 0.0
        },
        {
          "name": "master/gpus_revocable_used",
          "value": 0.0
        },
        {
          "name": "master/disk_revocable_percent",
          "value": 0.0
        },
        {
          "name": "master/mem_revocable_percent",
          "value": 0.0
        },
        {
          "name": "master/invalid_executor_to_framework_messages",
          "value": 0.0
        },
        {
          "name": "master/slave_shutdowns_scheduled",
          "value": 0.0
        },
        {
          "name": "master/slave_removals/reason_registered",
          "value": 0.0
        },
        {
          "name": "master/messages_suppress_offers",
          "value": 0.0
        },
        {
          "name": "master/uptime_secs",
          "value": 0.038900992
        },
        {
          "name": "allocator/mesos/resources/disk/total",
          "value": 0.0
        },
        {
          "name": "master/slave_removals/reason_unregistered",
          "value": 0.0
        },
        {
          "name": "master/disk_total",
          "value": 0.0
        },
        {
          "name": "master/messages_resource_request",
          "value": 0.0
        },
        {
          "name": "master/cpus_total",
          "value": 0.0
        },
        {
          "name": "master/valid_framework_to_executor_messages",
          "value": 0.0
        },
        {
          "name": "master/cpus_used",
          "value": 0.0
        },
        {
          "name": "master/slave_removals/reason_unhealthy",
          "value": 0.0
        },
        {
          "name": "master/messages_kill_task",
          "value": 0.0
        },
        {
          "name": "master/slave_shutdowns_canceled",
          "value": 0.0
        },
        {
          "name": "master/messages_deactivate_framework",
          "value": 0.0
        },
        {
          "name": "master/messages_unregister_framework",
          "value": 0.0
        },
        {
          "name": "master/mem_revocable_total",
          "value": 0.0
        },
        {
          "name": "master/messages_reregister_framework",
          "value": 0.0
        },
        {
          "name": "master/dropped_messages",
          "value": 0.0
        },
        {
          "name": "master/invalid_framework_to_executor_messages",
          "value": 0.0
        },
        {
          "name": "master/tasks_error",
          "value": 0.0
        },
        {
          "name": "master/tasks_lost",
          "value": 0.0
        },
        {
          "name": "master/messages_reconcile_tasks",
          "value": 0.0
        },
        {
          "name": "master/tasks_killed",
          "value": 0.0
        },
        {
          "name": "master/tasks_finished",
          "value": 0.0
        },
        {
          "name": "master/frameworks_inactive",
          "value": 0.0
        },
        {
          "name": "master/tasks_running",
          "value": 0.0
        },
        {
          "name": "master/tasks_starting",
          "value": 0.0
        },
        {
          "name": "registrar/state_store_ms",
          "value": 5.55392
        },
        {
          "name": "master/mem_total",
          "value": 0.0
        },
        {
          "name": "master/outstanding_offers",
          "value": 0.0
        },
        {
          "name": "master/frameworks_active",
          "value": 0.0
        },
        {
          "name": "master/messages_authenticate",
          "value": 0.0
        },
        {
          "name": "master/disk_revocable_used",
          "value": 0.0
        },
        {
          "name": "master/frameworks_disconnected",
          "value": 0.0
        },
        {
          "name": "master/slaves_connected",
          "value": 0.0
        },
        {
          "name": "master/messages_status_update_acknowledgement",
          "value": 0.0
        },
        {
          "name": "master/elected",
          "value": 1.0
        }
      ]
    }
  }
  `
	SetOutput(mux, output)

	data, err := c.GetMetrics(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	eventQueueDispatches := *data.GetMetrics.Metrics[0].Value
	if eventQueueDispatches != 1.0 {
		t.Errorf("expected 1.0: got %f", eventQueueDispatches)
	}
}

func TestMasterGetMetrics(t *testing.T) {
	master, mux, teardown := MasterSetup()
	defer teardown()
	testGetMetrics(t, master, mux)
}

func TestAgentGetMetrics(t *testing.T) {
	agent, mux, teardown := AgentSetup()
	defer teardown()
	testGetMetrics(t, agent, mux)
}

package v1

import (
	"context"
	"testing"
)

func TestMasterGetFlags(t *testing.T) {
	master, mux, teardown := MasterSetup()
	defer teardown()

	// Setup Handler
	output := `
{
  "type": "GET_FLAGS",
  "get_flags": {
    "flags": [
      {
        "name": "acls",
        "value": ""
      },
      {
        "name": "agent_ping_timeout",
        "value": "15secs"
      },
      {
        "name": "agent_reregister_timeout",
        "value": "10mins"
      },
      {
        "name": "allocation_interval",
        "value": "1secs"
      },
      {
        "name": "allocator",
        "value": "HierarchicalDRF"
      },
      {
        "name": "authenticate_agents",
        "value": "true"
      },
      {
        "name": "authenticate_frameworks",
        "value": "true"
      },
      {
        "name": "authenticate_http_frameworks",
        "value": "true"
      },
      {
        "name": "authenticate_http_readonly",
        "value": "true"
      },
      {
        "name": "authenticate_http_readwrite",
        "value": "true"
      },
      {
        "name": "authenticators",
        "value": "crammd5"
      },
      {
        "name": "authorizers",
        "value": "local"
      },
      {
        "name": "credentials",
        "value": "/tmp/directory/credentials"
      },
      {
        "name": "framework_sorter",
        "value": "drf"
      },
      {
        "name": "help",
        "value": "false"
      },
      {
        "name": "hostname_lookup",
        "value": "true"
      },
      {
        "name": "http_authenticators",
        "value": "basic"
      },
      {
        "name": "http_framework_authenticators",
        "value": "basic"
      },
      {
        "name": "initialize_driver_logging",
        "value": "true"
      },
      {
        "name": "log_auto_initialize",
        "value": "true"
      },
      {
        "name": "logbufsecs",
        "value": "0"
      },
      {
        "name": "logging_level",
        "value": "INFO"
      },
      {
        "name": "max_agent_ping_timeouts",
        "value": "5"
      },
      {
        "name": "max_completed_frameworks",
        "value": "50"
      },
      {
        "name": "max_completed_tasks_per_framework",
        "value": "1000"
      },
      {
        "name": "quiet",
        "value": "false"
      },
      {
        "name": "recovery_agent_removal_limit",
        "value": "100%"
      },
      {
        "name": "registry",
        "value": "replicated_log"
      },
      {
        "name": "registry_fetch_timeout",
        "value": "1mins"
      },
      {
        "name": "registry_store_timeout",
        "value": "100secs"
      },
      {
        "name": "registry_strict",
        "value": "true"
      },
      {
        "name": "root_submissions",
        "value": "true"
      },
      {
        "name": "user_sorter",
        "value": "drf"
      },
      {
        "name": "version",
        "value": "false"
      },
      {
        "name": "webui_dir",
        "value": "/usr/local/share/mesos/webui"
      },
      {
        "name": "work_dir",
        "value": "/tmp/directory/master"
      },
      {
        "name": "zk_session_timeout",
        "value": "10secs"
      }
    ]
  }
}
  `
	SetOutput(mux, output)

	// Call
	getFlags, err := master.GetFlags(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	// Assert
	for _, flag := range getFlags.GetFlags.Flags {
		switch *flag.Name {
		case "agent_ping_timeout":
			if *flag.Value != "15secs" {
				t.Error("expected 15secs: got %s", *flag.Value)
			}
		case "authenticate_agents":
			if *flag.Value != "true" {
				t.Error("expected \"true\": got %s", *flag.Value)
			}
		}
	}
}

func TestAgentGetFlags(t *testing.T) {
	agent, mux, teardown := AgentSetup()
	defer teardown()

	output := `
	{
	  "type": "GET_FLAGS",
	  "get_flags": {
	    "flags": [
	      {
	        "name": "acls",
	        "value": ""
	      },
	      {
	        "name": "appc_simple_discovery_uri_prefix",
	        "value": "http://"
	      },
	      {
	        "name": "appc_store_dir",
	        "value": "/tmp/mesos/store/appc"
	      },
	      {
	        "name": "authenticate_http_readonly",
	        "value": "true"
	      },
	      {
	        "name": "authenticate_http_readwrite",
	        "value": "true"
	      },
	      {
	        "name": "authenticatee",
	        "value": "crammd5"
	      },
	      {
	        "name": "authentication_backoff_factor",
	        "value": "1secs"
	      },
	      {
	        "name": "authorizer",
	        "value": "local"
	      },
	      {
	        "name": "cgroups_cpu_enable_pids_and_tids_count",
	        "value": "false"
	      },
	      {
	        "name": "cgroups_enable_cfs",
	        "value": "false"
	      },
	      {
	        "name": "cgroups_hierarchy",
	        "value": "/sys/fs/cgroup"
	      },
	      {
	        "name": "cgroups_limit_swap",
	        "value": "false"
	      },
	      {
	        "name": "cgroups_root",
	        "value": "mesos"
	      },
	      {
	        "name": "container_disk_watch_interval",
	        "value": "15secs"
	      },
	      {
	        "name": "containerizers",
	        "value": "mesos"
	      },
	      {
	        "name": "credential",
	        "value": "/tmp/directory/credential"
	      },
	      {
	        "name": "default_role",
	        "value": "*"
	      },
	      {
	        "name": "disk_watch_interval",
	        "value": "1mins"
	      },
	      {
	        "name": "docker",
	        "value": "docker"
	      },
	      {
	        "name": "docker_kill_orphans",
	        "value": "true"
	      },
	      {
	        "name": "docker_registry",
	        "value": "https://registry-1.docker.io"
	      },
	      {
	        "name": "docker_remove_delay",
	        "value": "6hrs"
	      },
	      {
	        "name": "docker_socket",
	        "value": "/var/run/docker.sock"
	      },
	      {
	        "name": "docker_stop_timeout",
	        "value": "0ns"
	      },
	      {
	        "name": "docker_store_dir",
	        "value": "/tmp/mesos/store/docker"
	      },
	      {
	        "name": "docker_volume_checkpoint_dir",
	        "value": "/var/run/mesos/isolators/docker/volume"
	      },
	      {
	        "name": "enforce_container_disk_quota",
	        "value": "false"
	      },
	      {
	        "name": "executor_registration_timeout",
	        "value": "1mins"
	      },
	      {
	        "name": "executor_shutdown_grace_period",
	        "value": "5secs"
	      },
	      {
	        "name": "fetcher_cache_dir",
	        "value": "/tmp/directory/fetch"
	      },
	      {
	        "name": "fetcher_cache_size",
	        "value": "2GB"
	      },
	      {
	        "name": "frameworks_home",
	        "value": ""
	      },
	      {
	        "name": "gc_delay",
	        "value": "1weeks"
	      },
	      {
	        "name": "gc_disk_headroom",
	        "value": "0.1"
	      },
	      {
	        "name": "hadoop_home",
	        "value": ""
	      },
	      {
	        "name": "help",
	        "value": "false"
	      },
	      {
	        "name": "hostname_lookup",
	        "value": "true"
	      },
	      {
	        "name": "http_authenticators",
	        "value": "basic"
	      },
	      {
	        "name": "http_command_executor",
	        "value": "false"
	      },
	      {
	        "name": "http_credentials",
	        "value": "/tmp/directory/http_credentials"
	      },
	      {
	        "name": "image_provisioner_backend",
	        "value": "copy"
	      },
	      {
	        "name": "initialize_driver_logging",
	        "value": "true"
	      },
	      {
	        "name": "isolation",
	        "value": "posix/cpu,posix/mem"
	      },
	      {
	        "name": "launcher_dir",
	        "value": "/my-directory"
	      },
	      {
	        "name": "logbufsecs",
	        "value": "0"
	      },
	      {
	        "name": "logging_level",
	        "value": "INFO"
	      },
	      {
	        "name": "oversubscribed_resources_interval",
	        "value": "15secs"
	      },
	      {
	        "name": "perf_duration",
	        "value": "10secs"
	      },
	      {
	        "name": "perf_interval",
	        "value": "1mins"
	      },
	      {
	        "name": "qos_correction_interval_min",
	        "value": "0ns"
	      },
	      {
	        "name": "quiet",
	        "value": "false"
	      },
	      {
	        "name": "recover",
	        "value": "reconnect"
	      },
	      {
	        "name": "recovery_timeout",
	        "value": "15mins"
	      },
	      {
	        "name": "registration_backoff_factor",
	        "value": "10ms"
	      },
	      {
	        "name": "resources",
	        "value": "cpus:2;gpus:0;mem:1024;disk:1024;ports:[31000-32000]"
	      },
	      {
	        "name": "revocable_cpu_low_priority",
	        "value": "true"
	      },
	      {
	        "name": "sandbox_directory",
	        "value": "/mnt/mesos/sandbox"
	      },
	      {
	        "name": "strict",
	        "value": "true"
	      },
	      {
	        "name": "switch_user",
	        "value": "true"
	      },
	      {
	        "name": "systemd_enable_support",
	        "value": "true"
	      },
	      {
	        "name": "systemd_runtime_directory",
	        "value": "/run/systemd/system"
	      },
	      {
	        "name": "version",
	        "value": "false"
	      },
	      {
	        "name": "work_dir",
	        "value": "/tmp/directory"
	      }
	    ]
	  }
	}
	`

	SetOutput(mux, output)

	// Call
	getFlags, err := agent.GetFlags(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	// Assert
	for _, flag := range getFlags.GetFlags.Flags {
		switch *flag.Name {
		case "docker_registry":
			if *flag.Value != "https://registry-1.docker.io" {
				t.Error("expected https://registry-1.docker.io: got %s", *flag.Value)
			}
		}
	}
}

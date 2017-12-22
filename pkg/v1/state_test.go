package v1

import (
	"context"
	"testing"

	"github.com/miroswan/mesops/pkg/v1/agent"
	"github.com/miroswan/mesops/pkg/v1/master"
)

func TestMasterGetState(t *testing.T) {
	caller, mux, teardown := MasterSetup()
	defer teardown()

	output := `
  {
    "type": "GET_STATE",
    "get_state": {
      "get_agents": {
        "agents": [
          {
            "active": true,
            "agent_info": {
              "hostname": "myhost",
              "id": {
                "value": "628984d0-4213-4140-bcb0-99d7ef46b1df-S0"
              },
              "port": 34626,
              "resources": [
                {
                  "name": "cpus",
                  "role": "*",
                  "scalar": {
                    "value": 2.0
                  },
                  "type": "SCALAR"
                },
                {
                  "name": "mem",
                  "role": "*",
                  "scalar": {
                    "value": 1024.0
                  },
                  "type": "SCALAR"
                },
                {
                  "name": "disk",
                  "role": "*",
                  "scalar": {
                    "value": 1024.0
                  },
                  "type": "SCALAR"
                },
                {
                  "name": "ports",
                  "ranges": {
                    "range": [
                      {
                        "begin": 31000,
                        "end": 32000
                      }
                    ]
                  },
                  "role": "*",
                  "type": "RANGES"
                }
              ]
            },
            "pid": "slave(3)@127.0.1.1:34626",
            "registered_time": {
              "nanoseconds": 1470820172046531840
            },
            "total_resources": [
              {
                "name": "cpus",
                "role": "*",
                "scalar": {
                  "value": 2.0
                },
                "type": "SCALAR"
              },
              {
                "name": "mem",
                "role": "*",
                "scalar": {
                  "value": 1024.0
                },
                "type": "SCALAR"
              },
              {
                "name": "disk",
                "role": "*",
                "scalar": {
                  "value": 1024.0
                },
                "type": "SCALAR"
              },
              {
                "name": "ports",
                "ranges": {
                  "range": [
                    {
                      "begin": 31000,
                      "end": 32000
                    }
                  ]
                },
                "role": "*",
                "type": "RANGES"
              }
            ],
            "version": "1.1.0"
          }
        ]
      },
      "get_executors": {
        "executors": [
          {
            "agent_id": {
              "value": "628984d0-4213-4140-bcb0-99d7ef46b1df-S0"
            },
            "executor_info": {
              "command": {
                "shell": true,
                "value": ""
              },
              "executor_id": {
                "value": "default"
              },
              "framework_id": {
                "value": "628984d0-4213-4140-bcb0-99d7ef46b1df-0000"
              }
            }
          }
        ]
      },
      "get_frameworks": {
        "frameworks": [
          {
            "active": true,
            "connected": true,
            "framework_info": {
              "checkpoint": false,
              "failover_timeout": 0.0,
              "hostname": "abcdev",
              "id": {
                "value": "628984d0-4213-4140-bcb0-99d7ef46b1df-0000"
              },
              "name": "default",
              "principal": "my-principal",
              "role": "*",
              "user": "root"
            },
            "registered_time": {
              "nanoseconds": 1470820172039300864
            },
            "reregistered_time": {
              "nanoseconds": 1470820172039300864
            }
          }
        ]
      },
      "get_tasks": {
        "completed_tasks": [
          {
            "agent_id": {
              "value": "628984d0-4213-4140-bcb0-99d7ef46b1df-S0"
            },
            "executor_id": {
              "value": "default"
            },
            "framework_id": {
              "value": "628984d0-4213-4140-bcb0-99d7ef46b1df-0000"
            },
            "name": "test-task",
            "resources": [
              {
                "name": "cpus",
                "role": "*",
                "scalar": {
                  "value": 2.0
                },
                "type": "SCALAR"
              },
              {
                "name": "mem",
                "role": "*",
                "scalar": {
                  "value": 1024.0
                },
                "type": "SCALAR"
              },
              {
                "name": "disk",
                "role": "*",
                "scalar": {
                  "value": 1024.0
                },
                "type": "SCALAR"
              },
              {
                "name": "ports",
                "ranges": {
                  "range": [
                    {
                      "begin": 31000,
                      "end": 32000
                    }
                  ]
                },
                "role": "*",
                "type": "RANGES"
              }
            ],
            "state": "TASK_FINISHED",
            "status_update_state": "TASK_FINISHED",
            "status_update_uuid": "IWjmPnfgQCWxGVlNNwctcg==",
            "statuses": [
              {
                "agent_id": {
                  "value": "628984d0-4213-4140-bcb0-99d7ef46b1df-S0"
                },
                "container_status": {
                  "network_infos": [
                    {
                      "ip_addresses": [
                        {
                          "ip_address": "127.0.1.1"
                        }
                      ]
                    }
                  ]
                },
                "executor_id": {
                  "value": "default"
                },
                "source": "SOURCE_EXECUTOR",
                "state": "TASK_RUNNING",
                "task_id": {
                  "value": "eb5cb680-a998-4605-8811-e79db8734c02"
                },
                "timestamp": 1470820172.07315,
                "uuid": "hTaLQ0b5Q1OZuab7QclTKQ=="
              },
              {
                "agent_id": {
                  "value": "628984d0-4213-4140-bcb0-99d7ef46b1df-S0"
                },
                "container_status": {
                  "network_infos": [
                    {
                      "ip_addresses": [
                        {
                          "ip_address": "127.0.1.1"
                        }
                      ]
                    }
                  ]
                },
                "executor_id": {
                  "value": "default"
                },
                "source": "SOURCE_EXECUTOR",
                "state": "TASK_FINISHED",
                "task_id": {
                  "value": "eb5cb680-a998-4605-8811-e79db8734c02"
                },
                "timestamp": 1470820172.09382,
                "uuid": "IWjmPnfgQCWxGVlNNwctcg=="
              }
            ],
            "task_id": {
              "value": "eb5cb680-a998-4605-8811-e79db8734c02"
            }
          }
        ]
      }
    }
  }
  `

	SetOutput(mux, output)

	data, err := caller.GetState(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	// Lots of data to validate here. Let's start out by asserting that data is
	// a GetState
	d := func(i interface{}) interface{} {
		return i
	}(data)

	if _, ok := d.(*master.GetStateResponse); !ok {
		t.Error("data is not a pointer to GetState")
	}
}

func TestAgentGetState(t *testing.T) {
	caller, mux, teardown := AgentSetup()
	defer teardown()

	output := `
	{
	  "type": "GET_STATE",
	  "get_state": {
	    "get_executors": {
	      "executors": [
	        {
	          "executor_info": {
	            "command": {
	              "arguments": [
	                "mesos-executor",
	                "--launcher_dir=/my-directory"
	              ],
	              "shell": false,
	              "value": "my-directory"
	            },
	            "executor_id": {
	              "value": "1"
	            },
	            "framework_id": {
	              "value": "8903b84e-112f-4b5f-aad3-7366f6ae7ecc-0000"
	            },
	            "name": "Command Executor (Task: 1) (Command: sh -c 'sleep 1000')",
	            "resources": [
	              {
	                "name": "cpus",
	                "role": "*",
	                "scalar": {
	                  "value": 0.1
	                },
	                "type": "SCALAR"
	              },
	              {
	                "name": "mem",
	                "role": "*",
	                "scalar": {
	                  "value": 32.0
	                },
	                "type": "SCALAR"
	              }
	            ],
	            "source": "1"
	          }
	        }
	      ]
	    },
	    "get_frameworks": {
	      "frameworks": [
	        {
	          "framework_info": {
	            "checkpoint": false,
	            "failover_timeout": 0.0,
	            "hostname": "myhost",
	            "id": {
	              "value": "8903b84e-112f-4b5f-aad3-7366f6ae7ecc-0000"
	            },
	            "name": "default",
	            "principal": "my-principal",
	            "role": "*",
	            "user": "root"
	          }
	        }
	      ]
	    },
	    "get_tasks": {
	      "launched_tasks": [
	        {
	          "agent_id": {
	            "value": "8903b84e-112f-4b5f-aad3-7366f6ae7ecc-S0"
	          },
	          "framework_id": {
	            "value": "8903b84e-112f-4b5f-aad3-7366f6ae7ecc-0000"
	          },
	          "name": "",
	          "resources": [
	            {
	              "name": "cpus",
	              "role": "*",
	              "scalar": {
	                "value": 2.0
	              },
	              "type": "SCALAR"
	            },
	            {
	              "name": "mem",
	              "role": "*",
	              "scalar": {
	                "value": 1024.0
	              },
	              "type": "SCALAR"
	            },
	            {
	              "name": "disk",
	              "role": "*",
	              "scalar": {
	                "value": 1024.0
	              },
	              "type": "SCALAR"
	            },
	            {
	              "name": "ports",
	              "ranges": {
	                "range": [
	                  {
	                    "begin": 31000,
	                    "end": 32000
	                  }
	                ]
	              },
	              "role": "*",
	              "type": "RANGES"
	            }
	          ],
	          "state": "TASK_RUNNING",
	          "status_update_state": "TASK_RUNNING",
	          "status_update_uuid": "2qlPayEJRJGPeaWlahI+WA==",
	          "statuses": [
	            {
	              "agent_id": {
	                "value": "8903b84e-112f-4b5f-aad3-7366f6ae7ecc-S0"
	              },
	              "container_status": {
	                "executor_pid": 19846,
	                "network_infos": [
	                  {
	                    "ip_addresses": [
	                      {
	                        "ip_address": "127.0.1.1"
	                      }
	                    ]
	                  }
	                ]
	              },
	              "executor_id": {
	                "value": "1"
	              },
	              "source": "SOURCE_EXECUTOR",
	              "state": "TASK_RUNNING",
	              "task_id": {
	                "value": "1"
	              },
	              "timestamp": 1470898839.48066,
	              "uuid": "2qlPayEJRJGPeaWlahI+WA=="
	            }
	          ],
	          "task_id": {
	            "value": "1"
	          }
	        }
	      ]
	    }
	  }
	}
  `

	SetOutput(mux, output)

	data, err := caller.GetState(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	// Lots of data to validate here. Let's start out by asserting that data is
	// a GetState
	d := func(i interface{}) interface{} {
		return i
	}(data)

	if _, ok := d.(*agent.GetStateResponse); !ok {
		t.Error("data is not a pointer to GetState")
	}
}

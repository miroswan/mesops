package smoke

import (
	"context"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/miroswan/mesops/pkg/v1"
)

func boolify(s string) bool {
	switch s {
	case "true", "True", "TRUE":
		return true
	case "false", "False", "FALSE":
		return false
	default:
		return false
	}
}

func getMaster() *v1.Master {
	builder := v1.NewMasterBuilder("http://192.168.33.10:5050").
		SetHTTPClient(&http.Client{Timeout: 2 * time.Second})
	client, _ := builder.Build()
	return client
}

func TestGethealth(t *testing.T) {
	client := getMaster()
	getHealth, err := client.GetHealth(context.Background())
	if err != nil {
		t.Error(err)
	}
	if *getHealth.GetHealth.Healthy != true {
		t.Errorf("expected true: got %t", *getHealth.GetHealth.Healthy)
	}
}

func TestGetFlags(t *testing.T) {
	client := getMaster()
	getFlags, err := client.GetFlags(context.Background())
	if err != nil {
		t.Error(err)
	}
	for _, flag := range getFlags.GetFlags.Flags {
		if *flag.Name == "authenticate_frameworks" {
			if boolify(*flag.Value) != false {
				t.Errorf("expected false: got %t", *flag.Value)
			}
			break
		}
	}
}

func TestGetVersion(t *testing.T) {
	client := getMaster()
	getVersion, err := client.GetVersion(context.Background())
	if err != nil {
		t.Error(err)
	}
	if *getVersion.GetVersion.VersionInfo.Version != "1.4.1" {
		t.Errorf("expected 1.4.1: got %s", *getVersion.GetVersion.VersionInfo.Version)
	}
}

func TestGetMetrics(t *testing.T) {
	client := getMaster()
	getMetrics, err := client.GetMetrics(context.Background())
	if err != nil {
		t.Error(err)
	}
	for _, metric := range getMetrics.GetMetrics.Metrics {
		if *metric.Name == "master/slaves_active" {
			if *metric.Value != 1.0 {
				t.Errorf("expected 1.0, got %f", *metric.Value)
			}
		}
	}
}

func TestGetLoggingLevel(t *testing.T) {
	client := getMaster()
	getLoggingLevel, err := client.GetLoggingLevel(context.Background())
	if err != nil {
		t.Error(err)
	}
	if *getLoggingLevel.GetLoggingLevel.Level != 0 {
		t.Errorf("expected 0: got %d", *getLoggingLevel.GetLoggingLevel.Level)
	}
}

func TestSetLoggingLevel(t *testing.T) {
	client := getMaster()
	err := client.SetLoggingLevel(context.Background(), 1)
	if err != nil {
		t.Error(err)
	}
	getLoggingLevel, err := client.GetLoggingLevel(context.Background())
	if err != nil {
		t.Error(err)
	}
	if *getLoggingLevel.GetLoggingLevel.Level != 1 {
		t.Errorf("expected 1: got %d", *getLoggingLevel.GetLoggingLevel.Level)
	}
	err = client.SetLoggingLevel(context.Background(), 0)
	if err != nil {
		t.Error(err)
	}
}

func TestGetState(t *testing.T) {
	hostname := "192.168.33.10"
	client := getMaster()
	getState, err := client.GetState(context.Background())
	if err != nil {
		t.Error(err)
	}
	for _, value := range getState.GetState.GetAgents.Agents {
		if *value.AgentInfo.Hostname != hostname {
			t.Errorf("expected %s: got %s", hostname, *value.AgentInfo.Hostname)
		}
		break
	}
}

func TestGetAgents(t *testing.T) {
	client := getMaster()
	hostname := "192.168.33.10"
	getAgents, err := client.GetAgents(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for _, value := range getAgents.GetAgents.Agents {
		if *value.AgentInfo.Hostname != hostname {
			t.Errorf("expected %s: got %s", hostname, *value.AgentInfo.Hostname)
		}
		break
	}
}

func TestGetFrameworks(t *testing.T) {
	client := getMaster()
	name := "marathon"
	getFrameworks, err := client.GetFrameworks(context.Background())
	if err != nil {
		t.Error(err)
	}
	for _, value := range getFrameworks.GetFrameworks.Frameworks {
		if *value.FrameworkInfo.Name != name {
			t.Errorf("expected %s: got %s", name, *value.FrameworkInfo.Name)
		}
	}
}

func TestGetExecutors(t *testing.T) {
	client := getMaster()
	_, err := client.GetExecutors(context.Background())
	if err != nil {
		t.Error(err)
	}
}

func TestGetTasks(t *testing.T) {
	client := getMaster()
	_, err := client.GetTasks(context.Background())
	if err != nil {
		t.Error(err)
	}
}

func TestGetRoles(t *testing.T) {
	client := getMaster()
	_, err := client.GetRoles(context.Background())
	if err != nil {
		t.Error(err)
	}
}

func TestGetWeights(t *testing.T) {
	client := getMaster()
	_, err := client.GetWeights(context.Background())
	if err != nil {
		t.Error(err)
	}
}

func TestGetMaster(t *testing.T) {
	client := getMaster()
	address := "192.168.33.10"
	getMaster, err := client.GetMaster(context.Background())
	if err != nil {
		t.Error(err)
	}
	if *getMaster.GetMaster.MasterInfo.Address.IP != address {
		t.Errorf("expected %s: got %s", address, *getMaster.GetMaster.MasterInfo.Address)
	}
}

func TestGetMaintenanceStatus(t *testing.T) {
	client := getMaster()
	getMaintenanceStatus, err := client.GetMaintenanceStatus(context.Background())
	if err != nil {
		t.Error(err)
	}
	d := len(getMaintenanceStatus.GetMaintenanceStatus.Status.DrainingMachines)
	if d != 0 {
		t.Errorf("expected 0: got %d", d)
	}
}

func TestGetMaintenanceSchedule(t *testing.T) {
	client := getMaster()
	getMaintenanceSchedule, err := client.GetMaintenanceSchedule(context.Background())
	if err != nil {
		t.Error(err)
	}
	w := getMaintenanceSchedule.GetMaintenanceSchedule.Schedule.Windows
	if len(w) != 0 {
		t.Errorf("expected 0: got %d", len(w))
	}
}

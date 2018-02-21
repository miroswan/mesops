package smoke

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/mesos/go-proto/mesos/v1"
	"github.com/mesos/go-proto/mesos/v1/master"

	"github.com/miroswan/mesops/pkg/v1"
)

func getMaster() *v1.Master {
	builder := v1.NewMasterBuilder("http://192.168.33.10:5050").
		SetHTTPClient(&http.Client{Timeout: 2 * time.Second})
	client, _ := builder.Build()
	return client
}

func TestGethealth(t *testing.T) {
	client := getMaster()
	data, err := client.GetHealth(context.Background())
	if err != nil {
		t.Error(err)
	}
	result := data.GetHealth.GetHealthy()
	if result != true {
		t.Errorf("expected true: got %t", result)
	}
}

func TestGetVersion(t *testing.T) {
	client := getMaster()
	res, err := client.GetVersion(context.Background())
	if err != nil {
		t.Error(err)
	}
	result := res.GetGetVersion().GetVersionInfo().GetVersion()
	if result != "1.5.0" {
		t.Errorf("expected 1.5.0: got %s", result)
	}
}

func TestSetLoggingLevel(t *testing.T) {
	client := getMaster()
	level := uint32(2)
	nanoseconds := int64(600000)
	call := &mesos_v1_master.Call_SetLoggingLevel{
		Level: &level,
		Duration: &mesos_v1.DurationInfo{
			Nanoseconds: &nanoseconds,
		},
	}
	err := client.SetLoggingLevel(context.Background(), call)
	if err != nil {
		t.Fatal(err)
	}
	data, err := client.GetLoggingLevel(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	result := data.GetLoggingLevel.GetLevel()
	if level != 2 {
		t.Errorf("expected 2, got %d", result)
	}
}

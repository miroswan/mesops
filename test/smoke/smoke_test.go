package smoke

import (
	"context"
	"net/http"
	"testing"
	"time"

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
	if result != "1.4.1" {
		t.Errorf("expected 1.4.1: got %s", result)
	}
}

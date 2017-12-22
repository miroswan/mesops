package v1

import (
	"net/http"
	"testing"
)

func TestNewAgentBuilder(t *testing.T) {
	a := NewAgentBuilder("test-url")
	i := func(i interface{}) interface{} {
		return i
	}(a)
	if _, ok := i.(*AgentBuilder); !ok {
		t.Error("expected returned type to be a pointer to an AgentBuilder")
	}
}

func TestAgentSetHTTPClient(t *testing.T) {
	a := NewAgentBuilder("test-url").SetHTTPClient(http.DefaultClient)
	i := func(i interface{}) interface{} {
		return i
	}(a)
	if _, ok := i.(*AgentBuilder); !ok {
		t.Error("expected returned type to be a pointer to an AgentBuilder")
	}
}

func TestAgentMaxRetries(t *testing.T) {
	a := NewAgentBuilder("test-url").SetMaxRetries(5)
	i := func(i interface{}) interface{} {
		return i
	}(a)
	if _, ok := i.(*AgentBuilder); !ok {
		t.Error("expected returned type to be a pointer to an AgentBuilder")
	}
}

func TestAgentBuild(t *testing.T) {
	a, err := NewAgentBuilder("test-url").Build()
	if err != nil {
		t.Error(err)
	}
	i := func(i interface{}) interface{} {
		return i
	}(a)
	if _, ok := i.(*Agent); !ok {
		t.Error("expected returned type to be a pointer to an Agent")
	}
}

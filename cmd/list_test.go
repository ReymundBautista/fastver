package cmd

import (
	"testing"

	"github.com/fastly/go-fastly/fastly"
)

type mockFastlyClientList struct{}

func (f *mockFastlyClientList) ListServices(i *fastly.ListServicesInput) ([]*fastly.Service, error) {
	return []*fastly.Service{
		{
			Name: "Service A",
		},
		{
			Name: "Service B",
		},
		{
			Name: "Service C",
		},
	}, nil
}

func TestListServices(t *testing.T) {
	client := &mockFastlyClientList{}
	services, _ := listServices(client)

	if services[0].Name != "Service A" {
		t.Errorf("expected '%v' but got '%v'", "Service A", services[0].Name)
	}

	if services[1].Name != "Service B" {
		t.Errorf("expected '%v' but got '%v'", "Service B", services[1].Name)
	}

	if services[2].Name != "Service C" {
		t.Errorf("expected '%v' but got '%v'", "Service C", services[2].Name)
	}
}

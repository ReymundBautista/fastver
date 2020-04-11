package cmd

import (
	"testing"

	"github.com/fastly/go-fastly/fastly"
)

type mockFastlyClientGet struct{}

func (client *mockFastlyClientGet) GetService(i *fastly.GetServiceInput) (*fastly.Service, error) {
	return &fastly.Service{
		ID: "12345667780",
		// Need to add the value type or we will get this error:
		// "missing type in composite literal"
		Versions: []*fastly.Version{
			{
				Number: 1,
				Active: false,
			},
			{
				Number: 2,
				Active: false,
			},
			{
				Number: 3,
				Active: false,
			},
			{
				Number: 4,
				Active: false,
			},
			{
				Number: 5,
				Active: true,
			},
		},
	}, nil
}

func TestService(t *testing.T) {
	tests := []struct {
		name  string
		id    string
		limit int
	}{
		{
			name:  "Expect 3 versions",
			id:    "12345667780",
			limit: 3,
		},
		{
			name:  "Expect 4 versions",
			id:    "12345667780",
			limit: 4,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := &mockFastlyClientGet{}
			service, _ := getService(client)
			versions := filterVersions(service.Versions, test.limit)

			if service.ID != test.id {
				t.Errorf("Failed to get the expected service %v, got %v", test.id, service.ID)
			}

			if len(versions) != test.limit {
				t.Errorf("Failed to get the expected version length %d, got %d", test.limit, len(versions))
			}
		})
	}
}

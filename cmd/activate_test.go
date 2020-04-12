package cmd

import (
	"testing"
)

func TestVersionChecker(t *testing.T) {
	client := &mockFastlyClientGet{}
	service, _ := getService(client)

	t.Run("Valid version returns true response ", func(t *testing.T) {
		response := verifyVersion(3, service)
		if response != true {
			t.Errorf("Failed to verify that version %d exists, got %t response", 3, response)
		}
	})

	t.Run("Invalid version returns false response", func(t *testing.T) {
		response := verifyVersion(6, service)
		if response != false {
			t.Errorf("Failed to verify that version %d does not exist, got %t response", 6, response)
		}
	})
}

package cmd

import (
	"testing"

	"github.com/fastly/go-fastly/fastly"
)

type mockFastlyClient struct{}

func (f *mockFastlyClient) GetCurrentUser() (*fastly.User, error) {
	return &fastly.User{
		Login: "johnny@fakey.seed",
	}, nil
}

func TestAuth(t *testing.T) {
	client := &mockFastlyClient{}
	user, _ := getCurrentUser(client)
	login := user.Login
	expected := "johnny@fakey.seed"

	if login != expected {
		t.Errorf("expected '%v' but got '%v'", expected, login)
	}
}

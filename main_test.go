package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_WhenHandlesRegistered_ThenDownloadRouteIsRegistered(t *testing.T) {
	// Given
	givenRequiredEnvsRegistered()
	port := "2000"
	expectedMethod := "POST"
	expectedPath := "/download"

	// When
	engine := handleRequests(port)

	// Then
	routes := engine.Routes()
	for _, route := range routes {
		if route.Method == expectedMethod &&
			route.Path == expectedPath {
			return
		}
	}
	assert.Fail(t, "Required registered method not found.")
}

func givenRequiredEnvsRegistered() {
	envs := []struct {
		key   string
		value string
	}{
		{"BITTORRENT_SERVICE_USERNAME", "username"},
		{"BITTORRENT_SERVICE_PASSWORD", "password"},
		{"BITTORRENT_SERVICE_URL", "http://url.url"},
		{"PLEX_TOKEN", "token"},
		{"PLEX_SERVICE_URL", "http://url.url"},
	}

	for _, env := range envs {
		os.Setenv(env.key, env.value)
	}

}

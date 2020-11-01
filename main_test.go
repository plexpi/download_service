package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_WhenHandlesRegistered_ThenDownloadRouteIsRegistered(t *testing.T) {
	// Given
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

package plex

import (
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getLibraries(t *testing.T) {

	scanner := givenMediaScanner()

	libraries, err := scanner.getLibraries()
	assert.NoError(t, err)
	assert.NotEmpty(t, libraries)
}

func Test_ScanLibraries(t *testing.T) {
	scanner := givenMediaScanner()

	err := scanner.ScanLibraries()
	assert.NoError(t, err)
}

func givenMediaScanner() MediaScanner {
	return NewMediaScanner(http.DefaultClient, os.Getenv("PLEX_TOKEN"))
}

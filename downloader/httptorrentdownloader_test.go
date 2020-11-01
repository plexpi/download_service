package downloader

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_WhenDownloadCalledWithEmptyBody_ThenExpectBadRequestError(t *testing.T) {
	// Given
	mockBittorrentAPI := givenMockBittorrentAPI()
	mockMediaScanner := givenMockMediaScanner()
	downloader := givenHTTPTorrentDownloader(mockBittorrentAPI, mockMediaScanner)
	router := givenRouterServingDownload(downloader)

	// When
	recorder := whenDownloadCalledWithBody(t, router, "")

	// Then
	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func Test_GivenAddTorrentFails_WhenDownloadCalledWithValidBody_ThenExpectInternalServerError(t *testing.T) {
	// Given
	expectedError := errors.New("sad error")
	mockBittorrentAPI := givenMockBittorrentAPI().GivenAddTorrentFails(expectedError)
	mockMediaScanner := givenMockMediaScanner()
	downloader := givenHTTPTorrentDownloader(mockBittorrentAPI, mockMediaScanner)
	router := givenRouterServingDownload(downloader)

	// When
	recorder := whenDownloadCalledWithBody(t, router, givenValidBody())

	// Then
	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	assert.Equal(t, expectedError.Error(), recorder.Body.String())
}

func Test_GivenAddTorrentSucceedsAndMediaScannerFails_WhenDownloadCalled_ThenExpectInternalServerError(t *testing.T) {
	// Given
	expectedError := errors.New("sad error")
	mockBittorrentAPI := givenMockBittorrentAPI().GivenAddTorrentSucceeds()
	mockMediaScanner := givenMockMediaScanner().GivenScanLibrariesFails(expectedError)
	downloader := givenHTTPTorrentDownloader(mockBittorrentAPI, mockMediaScanner)
	router := givenRouterServingDownload(downloader)

	// When
	recorder := whenDownloadCalledWithBody(t, router, givenValidBody())

	// Then
	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	assert.Equal(t, fmt.Sprintf("Could not refresh library: %s", expectedError), recorder.Body.String())
}

func Test_GivenAddTorrentSucceedsAndMediaScannerSucceeds_WhenDownloadCalled_ThenExpectStatusOK(t *testing.T) {
	// Given
	mockBittorrentAPI := givenMockBittorrentAPI().GivenAddTorrentSucceeds()
	mockMediaScanner := givenMockMediaScanner().GivenScanLibrariesSucceeds()
	downloader := givenHTTPTorrentDownloader(mockBittorrentAPI, mockMediaScanner)
	router := givenRouterServingDownload(downloader)

	// When
	recorder := whenDownloadCalledWithBody(t, router, givenValidBody())

	// Then
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Empty(t, recorder.Body.String())
}

func givenHTTPTorrentDownloader(torrentAPI BittorrentAPI, mediaScanner MediaScanner) HTTPTorrentDownloader {
	return HTTPTorrentDownloader{torrentAPI, mediaScanner, 0}
}

func givenMockBittorrentAPI() *MockBittorrentAPI {
	return new(MockBittorrentAPI)
}

func givenMockMediaScanner() *MockMediaScanner {
	return new(MockMediaScanner)
}

func givenRouterServingDownload(downloader HTTPTorrentDownloader) *gin.Engine {
	router := gin.Default()
	router.POST("/download", downloader.Download)
	return router
}

func givenValidBody() string {
	return `{
		"category": "movies",
		"url": "http://torrent.torrent"
	}`
}

func whenDownloadCalledWithBody(t *testing.T, router *gin.Engine, body string) *httptest.ResponseRecorder {
	request, err := http.NewRequest("POST", "/download", strings.NewReader(body))
	require.NoError(t, err)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	return recorder
}

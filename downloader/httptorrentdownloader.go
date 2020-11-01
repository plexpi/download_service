package downloader

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// BittorrentAPI ...
type BittorrentAPI interface {
	AddTorrent(url, category string) error
}

// MediaScanner ...
type MediaScanner interface {
	ScanLibraries() error
}

// HTTPTorrentDownloader ...
type HTTPTorrentDownloader struct {
	torrentAPI   BittorrentAPI
	mediaScanner MediaScanner
}

// NewHTTPTorrentDownloader ...
func NewHTTPTorrentDownloader(
	torrentAPI BittorrentAPI,
	mediaScanner MediaScanner) HTTPTorrentDownloader {
	return HTTPTorrentDownloader{
		torrentAPI:   torrentAPI,
		mediaScanner: mediaScanner,
	}
}

// Download ...
func (downloader HTTPTorrentDownloader) Download(c *gin.Context) {
	fmt.Println("Download called")

	var request downloadRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("DownloadRequest: %+v \n", request)
	if err := downloader.torrentAPI.AddTorrent(request.URL, request.Category); err != nil {
		c.String(http.StatusInternalServerError, "%s", err.Error())
		return
	}

	time.Sleep(5 * time.Second)
	if err := downloader.mediaScanner.ScanLibraries(); err != nil {
		c.String(http.StatusInternalServerError, "Could not refresh library: %s", err.Error())
		return
	}

	c.Status(200)
}

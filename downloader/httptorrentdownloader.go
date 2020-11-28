package downloader

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/plexpi/download_service/bittorrent"
)

// MediaScanner ...
type MediaScanner interface {
	ScanLibraries() error
}

// HTTPTorrentDownloader ...
type HTTPTorrentDownloader struct {
	torrentAPI   bittorrent.API
	mediaScanner MediaScanner
	scanerOffset time.Duration
}

// NewHTTPTorrentDownloader ...
func NewHTTPTorrentDownloader(
	torrentAPI bittorrent.API,
	mediaScanner MediaScanner) HTTPTorrentDownloader {
	return HTTPTorrentDownloader{
		torrentAPI:   torrentAPI,
		mediaScanner: mediaScanner,
		scanerOffset: 5,
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
	addTorrentParams := downloader.addTorrentParams(request)
	if err := downloader.torrentAPI.AddTorrent(addTorrentParams); err != nil {
		c.String(http.StatusInternalServerError, "%s", err.Error())
		return
	}

	time.Sleep(downloader.scanerOffset * time.Second)
	if err := downloader.mediaScanner.ScanLibraries(); err != nil {
		c.String(http.StatusInternalServerError, "Could not refresh library: %s", err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func (downloader HTTPTorrentDownloader) addTorrentParams(request downloadRequest) bittorrent.AddTorrentParams {
	return bittorrent.AddTorrentParams{
		Category:           request.Category,
		Location:           request.Location,
		URLs:               request.URL,
		SequentialDownload: request.SequentialDownload,
		FirstLastPiecePrio: request.FirstLastPiecePrio,
		AutoTMM:            request.AutoTMM,
	}
}

package downloader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
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
func (downloader HTTPTorrentDownloader) Download(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Download called")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var request *downloadRequest
	json.Unmarshal(reqBody, &request)

	if request == nil {
		json.NewEncoder(w).Encode("Failed to parse request.")
	}

	fmt.Printf("DownloadRequest: %+v \n", request)
	if err := downloader.torrentAPI.AddTorrent(request.URL, request.Category); err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	time.Sleep(2 * time.Second)
	if err := downloader.mediaScanner.ScanLibraries(); err != nil {
		json.NewEncoder(w).Encode(fmt.Errorf("Could not refresh library: %s", err).Error())
		return
	}
	json.NewEncoder(w).Encode("OK")
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/plexpi/download_service/bittorrent"
	"github.com/plexpi/download_service/downloader"
	"github.com/plexpi/download_service/plex"
)

const (
	defaultPort = "45780"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	bittorrentAPI := newBittorentAPI()
	mediaScanner := plex.NewMediaScanner(http.DefaultClient, os.Getenv("PLEX_TOKEN"))
	torretDownloader := downloader.NewHTTPTorrentDownloader(bittorrentAPI, mediaScanner)
	myRouter.HandleFunc("/download", torretDownloader.Download).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	fmt.Printf("Listening on port: %s \n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), myRouter))
}

func main() {
	fmt.Println("Download service started.")
	handleRequests()
}

func newBittorentAPI() bittorrent.API {
	username := os.Getenv("BITTORRENT_SERVICE_USERNAME")
	password := os.Getenv("BITTORRENT_SERVICE_USERNAME")
	url := os.Getenv("BITTORRENT_SERVICE_URL")
	return bittorrent.NewAPI(
		url,
		username,
		password,
		http.Client{})
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/adborbas/plexpicore/bittorrent"
	"github.com/adborbas/plexpicore/downloader"
	"github.com/adborbas/plexpicore/plex"
	"github.com/gorilla/mux"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	bittorrentAPI := newBittorentAPI()
	mediaScanner := plex.NewMediaScanner(http.DefaultClient, os.Getenv("PLEX_TOKEN"))
	torretDownloader := downloader.NewHTTPTorrentDownloader(bittorrentAPI, mediaScanner)
	myRouter.HandleFunc("/download", torretDownloader.Download).Methods("POST")

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Rest API v1.0")
	handleRequests()
}

func newBittorentAPI() bittorrent.API {
	return bittorrent.NewAPI("admin", "adminadmin", http.Client{})
}

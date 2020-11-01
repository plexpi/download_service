package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/plexpi/download_service/bittorrent"
	"github.com/plexpi/download_service/downloader"
	"github.com/plexpi/download_service/plex"
)

const (
	defaultPort = "45780"
)

func handleRequests(port string) {
	engine := newEngine()
	registerDownloader(engine)

	fmt.Printf("Listening on port: %s \n", port)
	engine.Run(fmt.Sprintf(":%s", port))
}

func main() {
	fmt.Println("Download service started.")

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	handleRequests(port)
}

func registerDownloader(engine *gin.Engine) {
	bittorrentAPI := newBittorentAPI()
	mediaScanner := newPlexMediaScanner()
	torretDownloader := downloader.NewHTTPTorrentDownloader(bittorrentAPI, mediaScanner)

	engine.POST("/download", torretDownloader.Download)
}

func newBittorentAPI() bittorrent.API {
	username := readRequiredEnv("BITTORRENT_SERVICE_USERNAME")
	password := readRequiredEnv("BITTORRENT_SERVICE_PASSWORD")
	url := readRequiredEnv("BITTORRENT_SERVICE_URL")
	return bittorrent.NewAPI(
		url,
		username,
		password,
		http.Client{})
}

func newPlexMediaScanner() downloader.MediaScanner {
	plexToken := readRequiredEnv("PLEX_TOKEN")
	url := readRequiredEnv("PLEX_SERVICE_URL")
	scanner := plex.NewMediaScanner(url, http.DefaultClient, plexToken)
	return scanner
}

func newEngine() *gin.Engine {
	engine := gin.Default()
	engine.Use(gin.Logger())

	engine.GET("/log200", func(c *gin.Context) {
		c.Status(200)
	})
	return engine
}

func readRequiredEnv(key string) string {
	env := os.Getenv(key)
	if env == "" {
		panic(fmt.Sprintf("Could not read env value for: %s", key))
	}

	return env
}

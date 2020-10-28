package downloader

type downloadRequest struct {
	Category string `json:"category"`
	URL      string `json:"url"`
}

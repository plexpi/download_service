package downloader

type downloadRequest struct {
	Category           string `json:"category"`
	Location           string `json:"location"`
	URL                string `json:"url" binding:"required"`
	SequentialDownload bool   `json:"is_sequential"`
	FirstLastPiecePrio bool   `json:"is_first_last_prio"`
	AutoTMM            bool   `json:"is_auto_managed"`
}

package bittorrent

type startTorrentRequest struct {
	Category string `json:"category"`
	URL      string `json:"url"`
}

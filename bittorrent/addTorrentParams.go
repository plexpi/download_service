package bittorrent

import (
	"net/url"
	"strconv"
)

// AddTorrentParams ...
type AddTorrentParams struct {
	Category           string
	Location           string
	URLs               string
	SequentialDownload bool
	FirstLastPiecePrio bool
	AutoTMM            bool
}

// URLValues ...
func (params AddTorrentParams) urlValues() url.Values {
	data := url.Values{}
	data.Set("urls", params.URLs)
	data.Set("category", params.Category)
	data.Set("savepath", params.Location)
	data.Set("sequentialDownload", strconv.FormatBool(params.SequentialDownload))
	data.Set("firstLastPiecePrio", strconv.FormatBool(params.FirstLastPiecePrio))
	data.Set("autoTMM", strconv.FormatBool(params.AutoTMM))
	return data
}

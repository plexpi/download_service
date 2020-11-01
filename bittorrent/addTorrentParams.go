package bittorrent

import (
	"net/url"
	"strconv"
)

type addTorrentParams struct {
	category           string
	urls               string
	sequentialDownload bool
	firstLastPiecePrio bool
	autoTMM            bool
}

func (params addTorrentParams) urlValues() url.Values {
	data := url.Values{}
	data.Set("urls", params.urls)
	data.Set("category", params.category)
	data.Set("sequentialDownload", strconv.FormatBool(params.sequentialDownload))
	data.Set("firstLastPiecePrio", strconv.FormatBool(params.firstLastPiecePrio))
	data.Set("autoTMM", strconv.FormatBool(params.autoTMM))
	return data
}

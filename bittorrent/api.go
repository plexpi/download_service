package bittorrent

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type addTorrentParams struct {
	urls               string
	sequentialDownload bool
	firstLastPiecePrio bool
	savepath           string
}

func (params addTorrentParams) urlValues() url.Values {
	data := url.Values{}
	data.Set("urls", params.urls)
	data.Set("sequentialDownload", strconv.FormatBool(params.sequentialDownload))
	data.Set("firstLastPiecePrio", strconv.FormatBool(params.firstLastPiecePrio))
	data.Set("savepath", params.savepath)
	return data
}

// API ...
type API struct {
	baseURL string
	client  http.Client
}

// NewAPI ...
func NewAPI(baseURL, username, password string, client http.Client) API {
	client.Transport = &qbittorrentAuthTransport{
		baseURL:           baseURL,
		username:          username,
		password:          password,
		originalTransport: http.DefaultTransport,
		client:            client,
	}

	return API{
		baseURL: baseURL,
		client:  client,
	}
}

// AddTorrent ...
func (api API) AddTorrent(torrentURL, category string) error {
	fmt.Println("AddTorrent called")

	fmt.Println("Creating http request.")

	rawParams, err := api.buildAddTorrentParams(torrentURL, category)
	if err != nil {
		return err
	}

	encodedURLValues := rawParams.urlValues().Encode()
	httpRequest, err := http.NewRequest(http.MethodPost, api.baseURL+"/api/v2/torrents/add", strings.NewReader(encodedURLValues))
	if err != nil {
		fmt.Printf("Failed to create request: %s", err)
		return err
	}
	httpRequest.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	httpRequest.Header.Add("Content-Length", strconv.Itoa(len(encodedURLValues)))

	fmt.Printf("Sending request: %+v \n", httpRequest)

	resp, err := api.client.Do(httpRequest)
	if err != nil {
		fmt.Printf("Network error: %s \n", err.Error())
		return err
	}
	fmt.Printf("Response received: %+v \n", resp)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Could not read body: %s", string(body))
		return err
	}
	fmt.Printf("Received: %s \n", string(body))

	return nil
}

func (api API) buildAddTorrentParams(torrentURL, category string) (addTorrentParams, error) {
	savePath := ""
	switch category {
	case "series":
		savePath = "/media/plex/series/"
	case "movies":
		savePath = "/media/plex/movies/"
	default:
		return addTorrentParams{}, fmt.Errorf("Invalid category: %s", category)
	}

	return addTorrentParams{
		urls:               torrentURL,
		sequentialDownload: true,
		firstLastPiecePrio: true,
		savepath:           savePath,
	}, nil
}

type qbittorrentAuthTransport struct {
	baseURL           string
	username          string
	password          string
	originalTransport http.RoundTripper
	client            http.Client
}

func (transport *qbittorrentAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	cookie, err := transport.login()
	if err != nil {
		fmt.Printf("Failed to login: %s \n", err)
		return nil, err
	}

	fmt.Printf("New cookie: %s \n", cookie.String())
	req.AddCookie(cookie)
	fmt.Printf("Resending request with logged in: %+v \n", req)
	return transport.originalTransport.RoundTrip(req)

	// Clone the request
	// resp, err := transport.originalTransport.RoundTrip(req)
	// if err != nil {
	// 	return resp, err
	// }

	// if resp.StatusCode == http.StatusForbidden {
	// 	cookie, err := transport.login()
	// 	if err != nil {
	// 		fmt.Printf("Failed to login: %s \n", err)
	// 		return nil, err
	// 	}

	// 	fmt.Printf("New cookie: %s \n", cookie.String())
	// 	req.AddCookie(cookie)
	// 	fmt.Printf("Resending request with logged in: %+v \n", req)
	// 	return transport.originalTransport.RoundTrip(req)
	// }

	// return resp, err
}

func (transport *qbittorrentAuthTransport) login() (*http.Cookie, error) {
	fmt.Println("Logging in ...")
	url := transport.baseURL + fmt.Sprintf("/api/v2/auth/login?username=%s&password=%s", transport.username, transport.password)
	resp, err := transport.client.Get(url)
	if err != nil {
		fmt.Printf("Failed to login: %+v\n", resp)
		return nil, err
	}

	fmt.Println("Successfully loged in!")
	return resp.Cookies()[0], nil
}

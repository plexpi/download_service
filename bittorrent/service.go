package bittorrent

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// API ...
type API interface {
	AddTorrent(params AddTorrentParams) error
}

// Service ...
type Service struct {
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

	return Service{
		baseURL: baseURL,
		client:  client,
	}
}

// AddTorrent ...
func (service Service) AddTorrent(params AddTorrentParams) error {
	fmt.Println("AddTorrent called")

	fmt.Println("Creating http request.")

	encodedURLValues := params.urlValues().Encode()
	httpRequest, err := http.NewRequest(http.MethodPost, service.baseURL+"/api/v2/torrents/add", strings.NewReader(encodedURLValues))
	if err != nil {
		fmt.Printf("Failed to create request: %s", err)
		return err
	}
	httpRequest.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	httpRequest.Header.Add("Content-Length", strconv.Itoa(len(encodedURLValues)))

	fmt.Printf("Sending request: %+v \n", httpRequest)

	resp, err := service.client.Do(httpRequest)
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

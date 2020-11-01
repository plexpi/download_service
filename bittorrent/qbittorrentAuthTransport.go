package bittorrent

import (
	"fmt"
	"net/http"
)

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

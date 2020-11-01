package plex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

// MediaScanner ...
type MediaScanner struct {
	baseURL string
	client  *http.Client
	token   string
}

// NewMediaScanner ...
func NewMediaScanner(baseURL string, client *http.Client, token string) MediaScanner {
	return MediaScanner{
		baseURL: baseURL,
		client:  client,
		token:   token,
	}
}

// ScanLibraries ...
func (scanner MediaScanner) ScanLibraries() error {
	libraries, err := scanner.getLibraries()
	if err != nil {
		return err
	}

	fmt.Printf("Starting scan for %d libraries \n", len(libraries))
	errChan := make(chan error)
	var wg sync.WaitGroup
	for _, library := range libraries {
		wg.Add(1)
		fmt.Printf("Starting scan for: %s \n", library)
		go scanner.scanLibrary(library, &wg, errChan)
	}

	fmt.Println("Waiting for scans to complete.")
	wg.Wait()
	close(errChan)
	fmt.Println("All scans have completed.")

	errorString := ""
	for err := range errChan {
		errorString += fmt.Sprintf("%s\n", err.Error())
		return fmt.Errorf("failed to scan library(ies): %s", errorString)
	}

	return nil
}

func (scanner MediaScanner) scanLibrary(id string, wg *sync.WaitGroup, errChan chan error) {
	defer func() {
		fmt.Printf("Scan finised for: %s\n", id)
		wg.Done()
	}()

	request, err := http.NewRequest(
		http.MethodGet,
		scanner.url(fmt.Sprintf("/library/sections/%s/refresh", id)),
		nil)

	if err != nil {
		errChan <- err
		return
	}

	fmt.Printf("Sending request: %+v\n", request)
	resp, err := scanner.client.Do(request)
	if err != nil {
		fmt.Printf("Network error: %s \n", err.Error())
		errChan <- err
		return
	}
	fmt.Printf("Response received: %+v \n", resp)
}

func (scanner MediaScanner) getLibraries() ([]string, error) {
	request, err := http.NewRequest(http.MethodGet, scanner.url("/library/sections"), nil)
	if err != nil {
		return []string{}, err
	}
	request.Header.Add("Accept", "application/json")

	fmt.Printf("Sending request: %+v\n", request)
	resp, err := scanner.client.Do(request)
	if err != nil {
		fmt.Printf("Network error: %s \n", err.Error())
		return []string{}, err
	}
	fmt.Printf("Response received: %+v \n", resp)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Could not read body: %s", string(body))
		return []string{}, err
	}
	fmt.Printf("Received: %s \n", string(body))

	libraries := librariesResponse{}
	if err := json.Unmarshal(body, &libraries); err != nil {
		return []string{}, err
	}

	var ids []string
	for _, directory := range libraries.MediaContainer.Directory {
		ids = append(ids, directory.Key)
	}

	return ids, nil
}

func (scanner MediaScanner) url(path string) string {
	return scanner.baseURL + fmt.Sprintf("%s?X-Plex-Token=%s", path, scanner.token)
}

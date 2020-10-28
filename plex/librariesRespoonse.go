package plex

type librariesResponse struct {
	MediaContainer struct {
		Directory []struct {
			Key string `json:"key"`
		} `json:"Directory"`
	} `json:"MediaContainer"`
}

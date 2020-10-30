# Download Service

## Used technologies

- go
- go modules
- [mux](github.com/gorilla/mux)

## Requirements

- [Install](https://dev.to/rohansawant/installing-docker-and-docker-compose-on-the-raspberry-pi-in-5-simple-steps-3mgl) docker-compose on RaspberryPi.

- Add the following environment variables

    ```bash
    export BITTORRENT_SERVICE_USERNAME="<username>"
    export BITTORRENT_SERVICE_PASSWORD="<password>"
    export BITTORRENT_SERVICE_URL="<url:port>"
    ```
  
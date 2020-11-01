# Download Service

Part of [plexpi](https://github.com/plexpi/plexpi).

## Used technologies

- go
- go modules
- [gin](https://github.com/gin-gonic/gin)

## Requirements

- [Install](https://dev.to/rohansawant/installing-docker-and-docker-compose-on-the-raspberry-pi-in-5-simple-steps-3mgl) docker-compose on RaspberryPi.
- Start up a `qbittorrent` docker container, like [linuxserver/bittorrent](https://hub.docker.com/r/linuxserver/qbittorrent).
- Start up a `plex` [docker container], like [linuxserver/plex](https://hub.docker.com/r/linuxserver/plex).
  
## Start

1. Add the following environment variables

    ```bash
    export BITTORRENT_SERVICE_USERNAME="<username>"
    export BITTORRENT_SERVICE_PASSWORD="<password>"
    export BITTORRENT_SERVICE_URL="<url:port>"
    export PLEX_TOKEN="<plextoken>"
    export PLEX_SERVICE_URL="<plexurl>"
    ```

1. Run via docker.

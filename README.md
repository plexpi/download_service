# Download Service

## Used technologies

- go
- go modules
- [mux](github.com/gorilla/mux)

## Requirements

- [Install](https://dev.to/rohansawant/installing-docker-and-docker-compose-on-the-raspberry-pi-in-5-simple-steps-3mgl) docker-compose on RaspberryPi.

- Add the following environment variables

    ```bash
    export GH_ACCESS_TOKEN_USRERNAME="<github_username>"
    export GH_ACCESS_TOKEN_PASSWORD="<github_access_token>"
    export BITTORRENT_SERVICE_USERNAME="<username>"
    export BITTORRENT_SERVICE_PASSWORD="<password>"
    export PLEX_TOKEN="<token>"
    ```

- [Install](https://www.linuxbabe.com/ubuntu/install-qbittorrent-ubuntu-18-04-desktop-server) qBittorrent
  - [WebUI API](https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#add-new-torrent)
  
- [Install](https://pimylifeup.com/raspberry-pi-plex-server/) plex in RaspberryPi.
  - Web - `yourraspberrypiurl:32400/web/`
  - Add `plex` user to qbittorrent-nox user group. `sudo adduser plex  qbittorrent-nox`
  - [Update](https://support.plex.tv/articles/235974187-enable-repository-updating-for-supported-linux-server-distributions/)
  - [API](https://support.plex.tv/articles/201638786-plex-media-server-url-commands/)

## Bash Aliases

`sudo nano /home/user/.bash_aliases`

```bash
# transmission
alias transmission-start='sudo service transmission-daemon start'
alias transmission-stop='sudo service transmission-daemon stop'
alias transmission-reload='sudo service transmission-daemon reload'
```

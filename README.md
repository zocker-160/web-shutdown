# web-shutdown
Very simple tool which allows you to trigger shutdown / reboot using a HTTP GET request.

## Build

You need a Linux system and golang installed

- git clone this repository
- `go build web-shutdown.go`

## Run

`./web-shutdown [port]`

If no port is specified, 8080 will be used by default.

## Commands

- `/`: shows very basic website with links to the available triggers
- `/cmd/shutdown`: trigger a shutdown
- `/cmd/reboot`: trigger a reboot

## Installation as systemd service

- copy compiled binary to `/usr/local/bin` or `/usr/bin`
- copy `webshutdown.service` to `/etc/systemd/system`
- `sudo systemctl enable webshutdown`
- `sudo systemctl start webshutdown`

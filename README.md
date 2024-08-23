# radeon_exporter
[![Build](https://github.com/kmulvey/radeon_exporter/actions/workflows/build.yml/badge.svg)](https://github.com/kmulvey/radeon_exporter/actions/workflows/build.yml) [![Release](https://github.com/kmulvey/radeon_exporter/actions/workflows/release.yml/badge.svg)](https://github.com/kmulvey/radeon_exporter/actions/workflows/release.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/kmulvey/radeon_exporter)](https://goreportcard.com/report/github.com/kmulvey/radeon_exporter) [![Go Reference](https://pkg.go.dev/badge/github.com/kmulvey/radeon_exporter.svg)](https://pkg.go.dev/github.com/kmulvey/radeon_exporter)

Prometheus exporter for Radeon (AMD) graphics cards. 

## Installation and Usage
### Docker
- `docker pull kmulvey/radeon_exporter:latest`
- `docker run --publish 9200:9200 radeon-exporter`

### Systemd
Several linux package formats are available in the releases. Manual linux install can be done as follows:  
- `sudo cp radeon_exporter /usr/bin/` (this path can be changed if you like, just be sure to change the path in the service file as well)
- `sudo cp radeon_exporter.service /etc/systemd/system/`
- `sudo systemctl daemon-reload`
- `sudo systemctl enable radeon_exporter`
- `sudo systemctl restart radeon_exporter`
- Import grafana-config.json to your grafana instance
- enjoy!

![Screenshot](https://github.com/kmulvey/radeon_exporter/blob/main/screenshot.jpg?raw=true)

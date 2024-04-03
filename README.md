# radeon_exporter
[![radeon_exporter](https://github.com/kmulvey/radeon_exporter/actions/workflows/release_build.yml/badge.svg)](https://github.com/kmulvey/radeon_exporter/actions/workflows/release_build.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/kmulvey/radeon_exporter)](https://goreportcard.com/report/github.com/kmulvey/radeon_exporter) [![Go Reference](https://pkg.go.dev/badge/github.com/kmulvey/radeon_exporter.svg)](https://pkg.go.dev/github.com/kmulvey/radeon_exporter)

Prometheus exporter for Radeon (AMD) graphics cards. 

## Installation and Usage
- `sudo cp radeon_exporter /opt/` (this path can be changed if you like, just be sure to change the path in the service file as well)
- `sudo cp radeon_exporter.service /etc/systemd/system/`
- `sudo systemctl daemon-reload`
- `sudo systemctl enable radeon_exporter`
- `sudo systemctl restart radeon_exporter`
- Import grafana-config.json to your grafana instance
- enjoy!

![Original](https://github.com/kmulvey/radeon_exporter/blob/main/screenshot.jpg?raw=true)

# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

ASWS (A Static Web Server) is a lightweight, configurable static web server built with Go. It uses the Gin web framework and includes built-in Prometheus metrics support. The project is designed to be deployed as a Docker container, with multi-architecture support via GoReleaser.

## Architecture

### Single Binary Application
- Main entry point: `cmd/asws.go`
- Self-contained Go binary with no external dependencies
- All configuration via environment variables or command-line flags
- No test files present in the codebase

### Core Components

**Web Server (`cmd/asws.go`)**
- Uses Gin web framework in release mode (or debug mode when DEBUG=true)
- Structured logging via uber/zap with request logging middleware
- Two-server architecture:
  - Main server: Serves static content (configurable IP/port)
  - Metrics server: Prometheus metrics endpoint (runs in goroutine on separate port)

**Static File Serving**
- Primary static path: Configurable via STATIC_PATH/STATIC_DIR (default: "/" serves from "./www")
- Optional filesystem mode: Additional file browsing via FS_PATH/FS_DIR (default: "/files" from "./files")
- Custom 404 handling: Can serve custom 404.html or redirect to a path

**Metrics & Observability**
- Prometheus metrics exposed on separate port (default: 2112)
- Service info metrics include Go version, service version, and service name
- Request logging integrated with zap logger

## Common Commands

### Development

**Build the application:**
```bash
go build -o asws ./cmd/asws.go
```

**Run locally:**
```bash
go run ./cmd/asws.go -debug=true -port=8080
```

**Run with custom configuration:**
```bash
DEBUG=true PORT=8080 STATIC_DIR=./www go run ./cmd/asws.go
```

### Docker

**Build Docker image:**
```bash
docker build -t asws:local .
```

**Run Docker container:**
```bash
docker run -e DEBUG=true -p 2701:80 -v "$(pwd)"/www:/www txn2/asws:latest
```

### Release Management

**Test release locally (without publishing):**
```bash
goreleaser --skip-publish --rm-dist --skip-validate
```

**Create and publish release:**
```bash
GITHUB_TOKEN=$GITHUB_TOKEN goreleaser --rm-dist
```

## Configuration

All configuration uses environment variables or command-line flags (flags override env vars). Key variables:

- **IP/PORT**: Bind address and main server port (default: 127.0.0.1:8080)
- **STATIC_DIR/STATIC_PATH**: Where and how to serve static files (default: ./www at /)
- **FS_ENABLED/FS_DIR/FS_PATH**: Optional file browsing endpoint (default: disabled)
- **NOT_FOUND_REDIRECT/NOT_FOUND_REDIRECT_PATH**: Redirect behavior for 404s (default: false)
- **NOT_FOUND_FILE**: Custom 404 page (default: ./www/404.html)
- **DEBUG**: Enable debug mode and verbose logging (default: false)
- **METRICS/METRICS_PORT**: Prometheus metrics server (default: enabled on 2112)
- **APP_NAME/APP_VERSION**: Application metadata for logging

See `cmd/asws.go:20-35` for complete configuration reference with defaults.

## Release Process

GoReleaser handles multi-platform builds and Docker image creation:

- **Platforms**: Linux and Darwin (macOS) for 386, amd64, arm, arm64
- **Docker variants**:
  - Scratch-based images (minimal size, multi-tag: latest, versioned, major version)
  - Alpine-based images (with shell/debugging tools)
- **Artifacts**: Binaries, checksums, deb/rpm/apk packages
- **Version injection**: Version string injected at build time via ldflags (`-X main.Version={{.Version}}`)

The build is configured in `goreleaser.yml`.

## Directory Structure

```
cmd/           - Main application entry point
www/           - Default static content directory
files/         - Default file browsing directory (when FS_ENABLED=true)
dockerfiles/   - Docker build configurations (scratch and alpine variants)
```

## Dependencies

Core dependencies:
- `github.com/gin-gonic/gin` - HTTP web framework
- `github.com/gin-contrib/zap` - Zap logging middleware for Gin
- `go.uber.org/zap` - Structured logging
- `github.com/prometheus/client_golang` - Prometheus metrics

## Notes

- No automated tests exist in this codebase
- The application is designed for container deployment with security considerations (runs as nobody user in Docker)
- CGO is disabled for static binary compilation
- Version information is injected at build time, not hardcoded
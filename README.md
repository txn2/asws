# ASWS Static Web Server

## Environment Variable Defaults

- PORT="80"
- GIN_MODE="release" # options: release | debug
- STATIC_DIR="./www"
- STATIC_PATH="./www"
- FS_ENABLED="no"
- FS_DIR="./files"
- FS_PATH="/files"

# Development
Uses goreleaser:

Install goreleaser with brew (mac): brew install goreleaser/tap/goreleaser

Build without releasing: goreleaser --skip-publish --rm-dist --skip-validate

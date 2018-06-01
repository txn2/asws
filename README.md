# [ASWS] Static Web Server

## Serve a Static Site

The following example exposes port 2701 on your local machine and forwards all traffic to port 80 on the [asws docker container]:

```bash
docker run -e DEBUG=true -p 2701:80 -v "$(pwd)"/www:/www txn2/asws:1.2.1
```

## Environment Variable Defaults

- PORT="80"
- STATIC_DIR="./www"
- STATIC_PATH="./www"
- FS_ENABLED="no"
- FS_DIR="./files"
- FS_PATH="/files"
- DEBUG="false"

# Development
Uses goreleaser:

Install goreleaser with brew (mac): brew install goreleaser/tap/goreleaser

Build without releasing: goreleaser --skip-publish --rm-dist --skip-validate

[asws docker container]: https://hub.docker.com/r/txn2/asws/
[ASWS]: https://github.com/txn2/asws
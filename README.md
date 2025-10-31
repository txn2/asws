# [ASWS] Static Web Server

## Serve a Static Site

The following example exposes port 2701 on your local machine and forwards all traffic to port 80 on the [asws docker container]:

```bash
docker run -e DEBUG=true -p 2701:80 -v "$(pwd)"/www:/www txn2/asws:v1.6.1
```

## SPA (Single Page Application) Mode

For modern React, Vue, or Angular applications with client-side routing, enable SPA fallback mode:

```bash
docker run -e SPA_FALLBACK=true -p 8080:80 -v "$(pwd)"/dist:/www txn2/asws:latest
```

This mode serves `index.html` with a 200 status code for any non-file request, allowing client-side routers to handle routing.

**Note:** SPA_FALLBACK takes precedence over NOT_FOUND_REDIRECT and NOT_FOUND_FILE.

## Environment Variable Defaults

- PORT="80"
- STATIC_DIR="./www"
- STATIC_PATH="./www"
- SPA_FALLBACK="false"
- FS_ENABLED="no"
- FS_DIR="./files"
- FS_PATH="/files"
- DEBUG="false"
- METRICS="true"
- METRICS_PORT="9696"

### Build Release

Build test release:
```bash
goreleaser --skip-publish --rm-dist --skip-validate
```

Build and release:
```bash
GITHUB_TOKEN=$GITHUB_TOKEN goreleaser --rm-dist
```


[asws docker container]: https://hub.docker.com/r/txn2/asws/
[ASWS]: https://github.com/txn2/asws

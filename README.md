# [ASWS] Static Web Server

## Serve a Static Site

The following example exposes port 2701 on your local machine and forwards all traffic to port 80 on the [asws docker container]:

```bash
docker run -e DEBUG=true -p 2701:80 -v "$(pwd)"/www:/www txn2/asws:1.2.2
```

## Environment Variable Defaults

- PORT="80"
- STATIC_DIR="./www"
- STATIC_PATH="./www"
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
# Dexctl

## Description

A command line tool which communicate with [The Dex API](https://dexidp.io/docs/api/). With help of this tool you could `create`, `update` and `delete` such dex entity like static password of static oauth2 clients.

### Usage

Values for global flags `host`, `port`, `ca-cert`, `client-cert` and `client-key` could be also read from env variables with prefix `DEX_API`

```bash
DEX_API_HOST="api.mydex.com" dexctl create password --email my@email.com --password verystrongpassword
```

## Development

### Prepare

```bash
go mod download
```

### Build

```bash
go build -o bin/$(go env GOOS)/dexctl \
    -ldflags="-s -w \
    -X 'github.com/agilestacks/dexctl/cmd.ref=$(git rev-parse --abbrev-ref HEAD)' \
    -X 'github.com/agilestacks/dexctl/cmd.commit=$(git rev-parse --short HEAD)' \
    -X 'github.com/agilestacks/dexctl/cmd.buildAt=$(date +"%Y.%m.%d %H:%M %Z")'"
```

### Run

```bash
./bin/$(go env GOOS)/dexctl
```

### Docker image

Docker image is based on distoless image [static-debian11](https://github.com/GoogleContainerTools/distroless/blob/main/base/README.md)

```bash
IMAGE_NAME="gcr.io/superhub/dexctl";
IMAGE_TAG="$(git rev-parse --short HEAD)";
docker buildx build --tag "${IMAGE_NAME}:${IMAGE_TAG}" --tag "${IMAGE_NAME}:latest" . ;
```

### Contribute

Before commit and submit pull request run next commands

```bash
go fmt github.com/agilestacks/dexctl/...
go vet github.com/agilestacks/dexctl/...
```

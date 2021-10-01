ARG VARIANT="1.14-alpine"

# =============================================================================
#  Main Build
# =============================================================================
FROM golang:${VARIANT}

RUN apk add --no-cache \
    git \
    alpine-sdk \
    build-base

COPY . /workspaces

WORKDIR /workspaces

RUN go mod download

ENTRYPOINT go version && go test .

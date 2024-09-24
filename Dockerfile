# syntax=docker/dockerfile:1
ARG GO_VERSION=1.23.1
ARG ALPINE_VERSION=3.20

FROM golang:${GO_VERSION}-alpine AS base

RUN set -ex \
    && apk update \
    && apk add --no-cache ca-certificates \
    bash \
    git \
    openssh \
    gcc \
    musl-dev \
    linux-headers

WORKDIR /src
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x


FROM base AS builder
ARG TARGETOS=linux
ARG TARGETARCH=amd64
ENV GOOS=$TARGETOS
ENV GOARCH=$TARGETARCH
ENV CGO_ENABLED=0
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,target=. \
    go build -a -installsuffix cgo -o /bin/app /src/cmd/api


FROM base AS test
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,target=. \
    go test -v -coverprofile=/tmp/coverage.txt ./... > /tmp/result.txt; \
    [[ $? -eq 0 ]] || { cat /tmp/result.txt; exit 1; }

FROM scratch AS export-test
COPY --from=test /tmp/coverage.txt /
COPY --from=test /tmp/result.txt /

FROM scratch AS binaries
COPY --from=builder  /bin/app /bin/app

FROM alpine:${ALPINE_VERSION}
WORKDIR /app
COPY --from=builder /bin/app /bin/app
EXPOSE 5000
ENTRYPOINT ["/bin/app"]
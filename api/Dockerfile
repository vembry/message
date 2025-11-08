FROM golang:1.25.4-alpine AS base

ARG TARGETOS
ARG TARGETARCH

ENV CGO_ENABLED=0

WORKDIR /home

# RUN --mount=type=cache,target=/go/pkg/mod \
#     # --mount=type=bind,source=./go.mod,target=./api/go.mod \
#     # --mount=type=bind,source=./go.sum,target=./api/go.sum \
#     go mod download

FROM base AS build

WORKDIR /home

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=bind,source=./,target=./ \
    GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /tmp/api

FROM gcr.io/distroless/static-debian12


COPY --from=build /tmp/api /

CMD ["/api"]
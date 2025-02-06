# syntax=docker/dockerfile:1

#
# fetch-stage
#

ARG GO_VERSION=1.23.4
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION} AS fetch-stage
LABEL org.opencontainers.image.source=https://github.com/mattcollie/mattcollecutt.com
WORKDIR /src

COPY go.mod /src
COPY go.sum /src
RUN go mod download


#
# generate-stage
#

FROM ghcr.io/a-h/templ:latest AS generate-stage
COPY --chown=65532:65532 . /src
WORKDIR /src
RUN ["templ", "generate"]


#
# server-stage
#

ARG GO_VERSION=1.23.4
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION} AS server-stage
LABEL org.opencontainers.image.source=https://github.com/mattcollie/mattcollecutt.com
WORKDIR /src

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

ARG TARGETARCH
COPY --from=generate-stage /src /src

RUN go generate

RUN --mount=type=cache,target=/go/pkg/mod/ \
    CGO_ENABLED=0 GOARCH=$TARGETARCH go build -o /bin/server .


#
#  tailwind-stage
#

FROM node:22-alpine AS tailwind-stage
LABEL org.opencontainers.image.source=https://github.com/mattcollie/mattcollecutt.com

WORKDIR /src

# Copy the app package and package-lock.json file
COPY . /src

RUN npm install -D tailwindcss

RUN npx tailwindcss -i ./input.css -o ./static/output.css --minify


#
#  final
#

FROM gcr.io/distroless/static:nonroot AS final
LABEL org.opencontainers.image.source=https://github.com/mattcollie/mattcollecutt.com

COPY --from=tailwind-stage /src/static /bin/static
COPY --from=server-stage /bin/server /bin/
COPY --from=ghcr.io/tarampampam/curl:8.6.0 /bin/curl /bin/curl

EXPOSE 8080

HEALTHCHECK --interval=10s --timeout=4s --retries=3 CMD curl -f http://0.0.0.0:8080/api/health || exit 1
ENTRYPOINT [ "/bin/server", "-docker", "-static=/bin/static"]
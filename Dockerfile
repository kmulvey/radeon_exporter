FROM golang:1.22  AS build-stage

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -v -ldflags="-s -w" -o /radeon-exporter ./...


FROM gcr.io/distroless/base-debian12 AS build-release-stage

WORKDIR /

COPY --from=build-stage /radeon-exporter /radeon-exporter

EXPOSE 9200

USER nonroot:nonroot

ENTRYPOINT ["/radeon-exporter"]

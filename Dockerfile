FROM golang:latest AS build
WORKDIR /app

ARG GOPROXY
ARG GOSUMDB

COPY go.mod go.sum ./
RUN GOPROXY=${GOPROXY} GOSUMDB=${GOSUMDB} go mod download

COPY ./internal ./internal
RUN CGO_ENABLED=0 go test -v ./...

COPY . .
RUN CGO_ENABLED=0 go test -v ./...

RUN CGO_ENABLED=0 go build -o ./out/app ./cmd/app

FROM alpine:latest
COPY --from=build /app/out/app /usr/local/bin/app

CMD ["/usr/local/bin/app"]

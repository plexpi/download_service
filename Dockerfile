# Build go app
FROM golang:1.15-alpine as builder
WORKDIR /
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o /service .

ENTRYPOINT ["/service"]
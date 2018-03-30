FROM golang:1.10
WORKDIR /go/src/github.com/OpenBazaar/status-server
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build --ldflags '-w -extldflags "-static"' -o /status-server .

FROM scratch
COPY --from=0 /status-server /status-server
EXPOSE 80
ENTRYPOINT ["/status-server"]
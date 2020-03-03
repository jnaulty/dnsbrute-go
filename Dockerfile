FROM golang:1.13.2-buster as golang

WORKDIR /go/src/github.com/jnaulty/dnsbrute-go

COPY . .

RUN env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build; mv dnsbrute-go /dnsbrute-go

FROM scratch

COPY --from=golang /dnsbrute-go /dnsbrute-go

COPY subdomain-lists /subdomain-lists

ENTRYPOINT ["/dnsbrute-go"]

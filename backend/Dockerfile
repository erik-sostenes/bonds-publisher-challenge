FROM golang:alpine AS build

RUN apk add --update git
WORKDIR /go/src/github.com/erik-sostenes/bonds-publisher-api
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/bonds-publisher-api ./cmd/http/server.go

FROM scratch
COPY --from=build /go/bin/bonds-publisher-api /go/bin/bonds-publisher-api
COPY --from=build /go/src/github.com/erik-sostenes/bonds-publisher-api/sql/* /docker-entrypoint-initdb.d/
ENTRYPOINT ["/go/bin/bonds-publisher-api"]

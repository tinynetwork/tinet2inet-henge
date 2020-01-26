FROM golang:1.13.6-alpine3.10

RUN apk add --update --no-cache ca-certificates git

RUN mkdir /work
WORKDIR /work
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 go build ./cmd/tn2ih/tn2ih.go

RUN mv tn2ih /usr/local/bin

EXPOSE 8080

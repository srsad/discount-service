FROM golang:1.19.3

WORKDIR /usr/src/app

RUN go install github.com/cosmtrek/air@latest
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY . .
RUN go mod tidy

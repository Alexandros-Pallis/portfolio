FROM golang:1.21-alpine as build

ENV PROJECT_DIR=/usr/src/app \
    GO111MODULE=on

WORKDIR /usr/src/app

ARG artifacts_folder

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY $artifacts_folder/go.mod $artifacts_folder/go.sum ./
RUN go mod download && go mod verify

RUN go install -mod=mod github.com/githubnemo/CompileDaemon

COPY $artifacts_folder ./

ENTRYPOINT CompileDaemon --polling=true --build="go build -o /build/app" --command="/build/app"

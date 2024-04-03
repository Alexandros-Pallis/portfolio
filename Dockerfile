ARG ALPINE_VERSION=latest
FROM golang:1.21-alpine as build

WORKDIR /usr/src/app

ARG artifacts_folder

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY $artifacts_folder/go.mod $artifacts_folder/go.sum ./
RUN go mod download && go mod verify

COPY $artifacts_folder ./
RUN go build -v -o /usr/local/bin/app .

FROM alpine:${ALPINE_VERSION}
RUN echo "ALPINE_VERSION=${ALPINE_VERSION}"
COPY --from=build /usr/local/bin/app /usr/local/bin/app

EXPOSE 80

CMD ["app"]

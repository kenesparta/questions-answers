FROM golang:1.17-alpine AS go_builder

WORKDIR /qaService

COPY . .

RUN apk add gcc musl-dev && \
    go mod tidy && \
    go get -d -u && \
    go test -v ./... -cover && \
    go build -o qas


# Building to finish docker container
FROM alpine:3.14

COPY --from=go_builder /qaService/qas .

ENTRYPOINT [ "./qas" ]

FROM golang:1.12 as build

WORKDIR /go/src/github.com/jarek-przygodzki/journald2elastic
COPY . .

RUN make build


# Now copy it into our base image.
FROM gcr.io/distroless/base

COPY --from=build /go/src/github.com/jarek-przygodzki/journald2elastic/bin/journald2elastic /journald2elastic

ENTRYPOINT ["/journald2elastic"]
FROM golang:1.18-bullseye as builder

RUN apt-get update -y && \
    apt-get install make && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /go/src/github.com/guni1192/spelunker

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . /go/src/github.com/guni1192/spelunker

RUN --mount=type=cache,target=/root/.cache/go-build \
    make build

FROM gcr.io/distroless/base-debian11

COPY --from=builder /go/src/github.com/guni1192/spelunker/bin/* /

CMD ["/spelunker"]

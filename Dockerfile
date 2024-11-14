FROM golang:1.23-bullseye as builder

RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN GOOS=linux CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main cmd/cli/main.go

FROM scratch
COPY --from=builder /build/main /bin/main
WORKDIR /work
ENTRYPOINT ["/bin/main"]

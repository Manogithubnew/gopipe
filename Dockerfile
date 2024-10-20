FROM golang:1.22.4 AS builder


WORKDIR /app
COPY go.mod ./
RUN go mod download


COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /test-app


FROM scratch
COPY --from=builder /test-app /test-app


EXPOSE 8083


CMD ["/test-app"]

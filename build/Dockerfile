FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN GOOS=linux go build -ldflags="-w -s" -o stress-test .

FROM scratch
COPY --from=builder /app/stress-test .
ENTRYPOINT ["/stress-test"]
CMD ["--help"]
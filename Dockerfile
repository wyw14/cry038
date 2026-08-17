FROM golang:1.24-bookworm AS builder
WORKDIR /classroom
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -o /bin/classroom-ready ./cmd/server
FROM debian:bookworm-slim
RUN useradd --uid 10002 --create-home classroom
USER classroom
COPY --from=builder /bin/classroom-ready /usr/local/bin/classroom-ready
EXPOSE 8080
ENTRYPOINT ["classroom-ready"]

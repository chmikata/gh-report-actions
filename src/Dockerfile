FROM golang:1.22.2-bookworm as builder

ENV CGO_ENABLED=0
ENV GO111MODULE=on
WORKDIR /app
COPY ./ ./
RUN go mod download && go build -o /app/out/gh-report-cli /app/main.go


FROM debian:bookworm-slim

WORKDIR /app
COPY --from=builder --chown=nobody:nogroup /app/out/gh-report-cli /app/

ENTRYPOINT ["/app/gh-reoprt-cli"]
CMD ["--help"]

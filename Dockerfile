FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git ca-certificates
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-w -s" -o findmecar

FROM scratch
COPY --from=builder /app/findmecar app

ENTRYPOINT ["/app"]
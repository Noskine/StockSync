FROM golang:1.23.2-alpine AS builder

WORKDIR /app

COPY . .
RUN go mod download

COPY . .

RUN cd ./cmd && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/app

FROM scratch

COPY --from=builder /go/bin/app /app

EXPOSE 3031

ENTRYPOINT ["/app"]
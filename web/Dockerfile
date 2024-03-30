FROM golang:1.22.1-alpine3.19 as builder

ARG TARGETOS=linux
ARG TARGETARCH=amd64

ADD ./go.mod .
ADD ./go.sum .
ADD ./main.go .
ADD ./cmd ./cmd
ADD ./internal ./internal

RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /app

FROM alpine:3.19

COPY --from=builder /app /app
COPY ./.env.example ./.env

CMD ["/app"]
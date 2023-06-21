FROM golang:1.20.5-alpine3.18 as builder
RUN apk update && apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/todos

FROM gcr.io/distroless/static-debian11
COPY --from=builder /go/bin/todos /

EXPOSE 8080
ENTRYPOINT [ "/todos" ]
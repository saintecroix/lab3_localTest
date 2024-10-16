FROM golang:latest AS APP
RUN mkdir /app
COPY /. /app
RUN apt update && apt install -y ca-certificates
WORKDIR /app
RUN go mod download
RUN go build ./cmd/web

FROM ubuntu:22.04
COPY --from=APP /app/. /app/.
COPY --from=APP /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /app
EXPOSE 4000
CMD ["/app/web"]
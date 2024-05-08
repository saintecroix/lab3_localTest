FROM golang:1.21 AS APP
RUN mkdir /app
COPY /. /app
WORKDIR /app
RUN go mod download
RUN go build ./cmd/web

FROM ubuntu:22.04
COPY --from=APP /app/. /app/.
WORKDIR /app
EXPOSE 4000
CMD ["/app/web"]
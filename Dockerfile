FROM golang:latest AS app
RUN mkdir /app
COPY /. /app
RUN apt update
WORKDIR /app
RUN go mod download
RUN go build ./cmd/web

FROM ubuntu:22.04
COPY --from=app /app/. /app/.
RUN apt update && apt install -y tzdata
WORKDIR /app
EXPOSE 4000
CMD ["/app/web"]
FROM golang:1.11 as builder
RUN mkdir /droplet-src
WORKDIR /droplet-src
COPY ./ .
RUN CGO_ENABLED=0 make setup all

FROM alpine:latest
RUN mkdir /app
WORKDIR /app
COPY --from=builder /droplet-src/bin/droplet ./
EXPOSE 8080
CMD ["./droplet"]
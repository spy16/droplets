FROM golang:1.11 as builder
RUN mkdir /droplets-src
WORKDIR /droplets-src
COPY ./ .
RUN CGO_ENABLED=0 make setup all

FROM alpine:latest
RUN mkdir /app
WORKDIR /app
COPY --from=builder /droplets-src/bin/droplets ./
COPY --from=builder /droplets-src/web ./web
EXPOSE 8080
CMD ["./droplets"]

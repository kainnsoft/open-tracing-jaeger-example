FROM golang:1.19.5-buster as builder

WORKDIR /service-b

COPY . ./

RUN go mod download
RUN mkdir -p build && \ 
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /service-b/build/service-b

FROM alpine:3.15.4 as task
COPY --from=builder /service-b/build /service-b
WORKDIR /service-b
EXPOSE 8082
CMD ["/service-b/service-b"]
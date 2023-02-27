FROM golang:1.19.5-buster as builder

WORKDIR /service-a

COPY . ./

RUN go mod download
RUN mkdir -p build && \ 
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /service-a/build/service-a

FROM alpine:3.15.4 as task
COPY --from=builder /service-a/build /service-a
WORKDIR /service-a
EXPOSE 8081
CMD ["/service-a/service-a"]
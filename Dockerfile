FROM golang:1.14.6-alpine3.12 as builder

COPY go.mod go.sum /go/src/github.com/CCoderX/parking_API/
WORKDIR /go/src/github.com/CCoderX/parking_API
RUN go mod download
COPY . /go/src/github.com/CCoderX/parking_API
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/parking_API github.com/CCoderX/parking_API


FROM alpine

RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/CCoderX/parking_API /usr/bin/parking_API

EXPOSE 8080 8080

ENTRYPOINT ["/usr/bin/parking_API"]

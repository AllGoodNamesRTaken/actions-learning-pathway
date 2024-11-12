FROM golang:1.22 AS builder

WORKDIR /go/src/

COPY go* .
COPY *.go .
RUN go mod download
RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o my-app .

# deployment image
FROM alpine:latest  
RUN apk --no-cache add ca-certificates

LABEL author="José Ramón"

WORKDIR /root/
COPY --from=builder /go/src/my-app .

CMD [ "./my-app" ]

EXPOSE 8080
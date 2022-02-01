FROM golang:1.17 as builder

LABEL maintainer="leffen <leffen@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myservice .

FROM alpine:3.15

RUN apk --no-cache add ca-certificates curl

WORKDIR /app

COPY --from=builder /app/myservice .
COPY public /app/public/
COPY template /app/template/

EXPOSE 8080

CMD ["/app/myservice"]
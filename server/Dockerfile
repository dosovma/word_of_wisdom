FROM golang:1.23
LABEL authors="madosov"

WORKDIR /server

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . ./

RUN CGO_ENABLED=0 go build -o tcpserver cmd/*

ENTRYPOINT ["./tcpserver"]

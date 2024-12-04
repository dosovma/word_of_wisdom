FROM golang:1.23
LABEL authors="madosov"

WORKDIR /client

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . ./

RUN CGO_ENABLED=0 go build -o tcpclient cmd/*

ENTRYPOINT ["./tcpclient"]

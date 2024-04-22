FROM golang:1.21.9

RUN go version

ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o app ./main.go

CMD ["./app"]

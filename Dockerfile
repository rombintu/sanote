FROM golang:1.17-alpine3.15
COPY . /opt/sanote/
WORKDIR /opt/sanote/
RUN go build -o dist/api cmd/main.go
CMD ["./dist/api"]
#Dockerfile 
FROM golang:latest
RUN go get github.com/astaxie/beego
RUN go get github.com/beego/bee
COPY . /go/src
ENV GOPATH=/go
WORKDIR /go
RUN CGO_ENABLED=1 GOOS=linux go build -o app main.go
EXPOSE 8080
CMD ["./app"]
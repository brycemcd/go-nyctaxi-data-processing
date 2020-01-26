# Run with: time docker run -v $(pwd):/opt/project -it taxi_data_go
FROM golang:1.13-buster

RUN mkdir /opt/project

WORKDIR /opt/project

COPY main.go .

RUN go get -d ./...

CMD /usr/local/go/bin/go build main.go && ./main

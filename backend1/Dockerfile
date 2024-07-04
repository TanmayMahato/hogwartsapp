
#Base Golang Image
FROM golang:alpine


WORKDIR /backend1


COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /bcd1

EXPOSE 8001

CMD [ "/bcd1" ]



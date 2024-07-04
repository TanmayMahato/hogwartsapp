
FROM golang:alpine


WORKDIR /backend3


COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /bcd3

EXPOSE 8003

CMD [ "/bcd3" ]
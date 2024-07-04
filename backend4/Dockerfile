FROM golang:alpine


WORKDIR /backend4


COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /bcd4

EXPOSE 8004

CMD [ "/bcd4" ]
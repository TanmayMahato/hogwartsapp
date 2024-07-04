
#Base Golang Image
FROM golang:1.22.5-alpine


WORKDIR /backend1


COPY . ./
RUN go mod download

COPY . .

RUN go build -o main server.go

EXPOSE 8001

CMD [ "./main" ]



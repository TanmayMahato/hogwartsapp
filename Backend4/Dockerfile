
FROM golang:latest


RUN mkdir app

COPY . /backend1/


WORKDIR /backend1/

RUN go build -o main .

EXPOSE 8004


CMD [ "/backend1/main" ]

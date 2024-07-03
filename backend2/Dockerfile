

FROM golang:latest


RUN mkdir app

COPY . /backend2/


WORKDIR /backend2/

RUN go build -o main .

EXPOSE 8002


CMD [ "/backend1/main" ]
